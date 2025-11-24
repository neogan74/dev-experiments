package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// ============================================================================
// FAN-OUT PATTERN
// One input channel distributes work to multiple workers
// ============================================================================

// Worker processes data from input channel
type Worker struct {
	ID     int
	Input  <-chan int
	Output chan<- Result
	ctx    context.Context
	wg     *sync.WaitGroup
}

type Result struct {
	WorkerID int
	Value    int
	Result   int
}

// Start begins processing for this worker
func (w *Worker) Start() {
	defer w.wg.Done()
	for {
		select {
		case <-w.ctx.Done():
			return
		case val, ok := <-w.Input:
			if !ok {
				return
			}
			// Simulate work
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
			result := val * val
			w.Output <- Result{
				WorkerID: w.ID,
				Value:    val,
				Result:   result,
			}
		}
	}
}

// FanOut distributes work from input to multiple workers
func FanOut(ctx context.Context, input <-chan int, numWorkers int) <-chan Result {
	output := make(chan Result, numWorkers)
	var wg sync.WaitGroup

	// Spawn workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		worker := &Worker{
			ID:     i + 1,
			Input:  input,
			Output: output,
			ctx:    ctx,
			wg:     &wg,
		}
		go worker.Start()
	}

	// Close output when all workers are done
	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

// ============================================================================
// FAN-IN PATTERN
// Multiple input channels merge into one output channel
// ============================================================================

// Source generates data
type Source struct {
	ID    int
	ctx   context.Context
	delay time.Duration
}

// Generate produces values on a channel
func (s *Source) Generate() <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		counter := 0
		for {
			select {
			case <-s.ctx.Done():
				return
			case <-time.After(s.delay):
				output <- counter
				counter++
				if counter >= 5 { // Each source generates 5 items
					return
				}
			}
		}
	}()
	return output
}

// FanIn merges multiple input channels into one output channel
func FanIn(ctx context.Context, channels ...<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup

	// Start a goroutine for each input channel
	for i, ch := range channels {
		wg.Add(1)
		go func(id int, c <-chan int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-c:
					if !ok {
						return
					}
					output <- val
				}
			}
		}(i, ch)
	}

	// Close output when all inputs are exhausted
	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

// FanInWithSelect uses select for merging (alternative implementation)
func FanInWithSelect(ctx context.Context, inputs ...<-chan int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)

		// Create cases for select
		cases := make([]chan int, len(inputs))
		for i, ch := range inputs {
			// Convert read-only to bidirectional for reflection
			cases[i] = make(chan int)
			go func(src <-chan int, dst chan int) {
				for v := range src {
					dst <- v
				}
				close(dst)
			}(ch, cases[i])
		}

		remaining := len(cases)
		for remaining > 0 {
			for i := 0; i < len(cases); i++ {
				if cases[i] == nil {
					continue
				}
				select {
				case <-ctx.Done():
					return
				case val, ok := <-cases[i]:
					if !ok {
						cases[i] = nil
						remaining--
						continue
					}
					output <- val
				default:
					// Non-blocking, move to next channel
				}
			}
			time.Sleep(time.Millisecond) // Prevent tight loop
		}
	}()

	return output
}

// ============================================================================
// MULTIPLEXOR PATTERN
// Route data from multiple sources to multiple destinations based on criteria
// ============================================================================

type Message struct {
	SourceID int
	Data     int
	TargetID int // Which output should receive this
}

// Multiplexor routes messages from multiple inputs to multiple outputs
type Multiplexor struct {
	numOutputs int
	outputs    []chan Message
	ctx        context.Context
	wg         sync.WaitGroup
}

// NewMultiplexor creates a new multiplexor with specified outputs
func NewMultiplexor(ctx context.Context, numOutputs int, bufferSize int) *Multiplexor {
	mx := &Multiplexor{
		numOutputs: numOutputs,
		outputs:    make([]chan Message, numOutputs),
		ctx:        ctx,
	}

	for i := 0; i < numOutputs; i++ {
		mx.outputs[i] = make(chan Message, bufferSize)
	}

	return mx
}

// Route starts routing messages from inputs to outputs
func (mx *Multiplexor) Route(inputs ...<-chan Message) {
	for _, input := range inputs {
		mx.wg.Add(1)
		go func(in <-chan Message) {
			defer mx.wg.Done()
			for {
				select {
				case <-mx.ctx.Done():
					return
				case msg, ok := <-in:
					if !ok {
						return
					}
					// Route to appropriate output
					targetID := msg.TargetID % mx.numOutputs
					select {
					case mx.outputs[targetID] <- msg:
					case <-mx.ctx.Done():
						return
					}
				}
			}
		}(input)
	}

	// Close all outputs when routing is complete
	go func() {
		mx.wg.Wait()
		for _, out := range mx.outputs {
			close(out)
		}
	}()
}

// GetOutput returns the channel for a specific output
func (mx *Multiplexor) GetOutput(id int) <-chan Message {
	if id < 0 || id >= mx.numOutputs {
		return nil
	}
	return mx.outputs[id]
}

// GetAllOutputs returns all output channels
func (mx *Multiplexor) GetAllOutputs() []<-chan Message {
	outputs := make([]<-chan Message, len(mx.outputs))
	for i, ch := range mx.outputs {
		outputs[i] = ch
	}
	return outputs
}

// ============================================================================
// ADVANCED: SCALABLE PATTERNS (2-200 sources/outputs)
// ============================================================================

// ScalableSource generates messages with configurable rate
type ScalableSource struct {
	ID         int
	TotalItems int
	RatePerSec int
	ctx        context.Context
}

func (s *ScalableSource) Generate() <-chan Message {
	output := make(chan Message, 10)
	go func() {
		defer close(output)
		interval := time.Second / time.Duration(s.RatePerSec)
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		count := 0
		for count < s.TotalItems {
			select {
			case <-s.ctx.Done():
				return
			case <-ticker.C:
				msg := Message{
					SourceID: s.ID,
					Data:     count,
					TargetID: rand.Intn(10), // Random routing
				}
				output <- msg
				count++
			}
		}
	}()
	return output
}

// ScalablePipeline creates a complete fan-in -> process -> fan-out -> multiplex pipeline
type ScalablePipeline struct {
	NumSources int
	NumWorkers int
	NumOutputs int
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewScalablePipeline(numSources, numWorkers, numOutputs int) *ScalablePipeline {
	ctx, cancel := context.WithCancel(context.Background())
	return &ScalablePipeline{
		NumSources: numSources,
		NumWorkers: numWorkers,
		NumOutputs: numOutputs,
		ctx:        ctx,
		cancel:     cancel,
	}
}

// Run executes the pipeline
func (sp *ScalablePipeline) Run() {
	fmt.Printf("\n=== Starting Scalable Pipeline ===\n")
	fmt.Printf("Sources: %d, Workers: %d, Outputs: %d\n\n",
		sp.NumSources, sp.NumWorkers, sp.NumOutputs)

	// Step 1: Create sources
	sourceChannels := make([]<-chan Message, sp.NumSources)
	for i := 0; i < sp.NumSources; i++ {
		source := &ScalableSource{
			ID:         i + 1,
			TotalItems: 5,
			RatePerSec: 10,
			ctx:        sp.ctx,
		}
		sourceChannels[i] = source.Generate()
	}

	// Step 2: Fan-in all sources
	merged := sp.fanInMessages(sourceChannels...)

	// Step 3: Fan-out to workers for processing
	processed := sp.processMessages(merged, sp.NumWorkers)

	// Step 4: Multiplex to outputs
	mx := NewMultiplexor(sp.ctx, sp.NumOutputs, 50)
	mx.Route(processed)

	// Step 5: Consume from all outputs
	var outputWg sync.WaitGroup
	for i, output := range mx.GetAllOutputs() {
		outputWg.Add(1)
		go func(id int, ch <-chan Message) {
			defer outputWg.Done()
			count := 0
			for msg := range ch {
				count++
				if count <= 3 || count%10 == 0 { // Sample output
					fmt.Printf("Output[%d] received: Source=%d, Data=%d\n",
						id, msg.SourceID, msg.Data)
				}
			}
			fmt.Printf("Output[%d] processed %d messages\n", id, count)
		}(i, output)
	}

	outputWg.Wait()
	fmt.Println("\n=== Pipeline Complete ===")
}

func (sp *ScalablePipeline) fanInMessages(channels ...<-chan Message) <-chan Message {
	output := make(chan Message, 100)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan Message) {
			defer wg.Done()
			for msg := range c {
				select {
				case output <- msg:
				case <-sp.ctx.Done():
					return
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func (sp *ScalablePipeline) processMessages(input <-chan Message, numWorkers int) <-chan Message {
	output := make(chan Message, 100)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case <-sp.ctx.Done():
					return
				case msg, ok := <-input:
					if !ok {
						return
					}
					// Simulate processing
					time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
					msg.Data = msg.Data * 2 // Simple transformation
					output <- msg
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func (sp *ScalablePipeline) Shutdown() {
	sp.cancel()
}

// ============================================================================
// EXAMPLES AND DEMONSTRATIONS
// ============================================================================

func demonstrateFanOut() {
	fmt.Println("\n=== FAN-OUT PATTERN ===")
	fmt.Println("One producer -> Multiple workers")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Create input channel
	input := make(chan int, 10)

	// Start producer
	go func() {
		for i := 1; i <= 20; i++ {
			select {
			case input <- i:
				fmt.Printf("Produced: %d\n", i)
			case <-ctx.Done():
				close(input)
				return
			}
		}
		close(input)
	}()

	// Fan out to 5 workers
	results := FanOut(ctx, input, 5)

	// Collect results
	count := 0
	for result := range results {
		fmt.Printf("Worker %d processed %d -> %d\n",
			result.WorkerID, result.Value, result.Result)
		count++
	}

	fmt.Printf("Total results processed: %d\n", count)
}

func demonstrateFanIn() {
	fmt.Println("\n=== FAN-IN PATTERN ===")
	fmt.Println("Multiple producers -> One consumer")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Create 4 sources
	numSources := 4
	sources := make([]<-chan int, numSources)
	for i := 0; i < numSources; i++ {
		source := &Source{
			ID:    i + 1,
			ctx:   ctx,
			delay: time.Millisecond * time.Duration(50+rand.Intn(100)),
		}
		sources[i] = source.Generate()
	}

	// Fan in all sources
	merged := FanIn(ctx, sources...)

	// Consume merged stream
	count := 0
	for val := range merged {
		fmt.Printf("Received: %d\n", val)
		count++
	}

	fmt.Printf("Total values received: %d\n", count)
}

func demonstrateMultiplexor() {
	fmt.Println("\n=== MULTIPLEXOR PATTERN ===")
	fmt.Println("Multiple sources -> Route to multiple outputs")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Create 3 input channels
	inputs := make([]<-chan Message, 3)
	for i := 0; i < 3; i++ {
		ch := make(chan Message, 5)
		inputs[i] = ch

		// Start producer for this input
		go func(id int, c chan Message) {
			defer close(c)
			for j := 0; j < 10; j++ {
				msg := Message{
					SourceID: id,
					Data:     j,
					TargetID: rand.Intn(4), // Route to one of 4 outputs
				}
				select {
				case c <- msg:
				case <-ctx.Done():
					return
				}
				time.Sleep(50 * time.Millisecond)
			}
		}(i+1, ch)
	}

	// Create multiplexor with 4 outputs
	mx := NewMultiplexor(ctx, 4, 10)
	mx.Route(inputs...)

	// Consume from all outputs
	var wg sync.WaitGroup
	for i, output := range mx.GetAllOutputs() {
		wg.Add(1)
		go func(id int, ch <-chan Message) {
			defer wg.Done()
			count := 0
			for msg := range ch {
				fmt.Printf("Output[%d] <- Source %d, Data: %d\n",
					id, msg.SourceID, msg.Data)
				count++
			}
			fmt.Printf("Output[%d] total: %d messages\n", id, count)
		}(i, output)
	}

	wg.Wait()
}

func RunDemos() {
	rand.Seed(time.Now().UnixNano())

	// Run demonstrations
	demonstrateFanOut()
	time.Sleep(500 * time.Millisecond)

	demonstrateFanIn()
	time.Sleep(500 * time.Millisecond)

	demonstrateMultiplexor()
	time.Sleep(500 * time.Millisecond)

	// Run scalable pipeline with configurable sources/workers/outputs
	fmt.Println("\n" + strings.Repeat("=", 60))

	// Small scale test
	pipeline := NewScalablePipeline(5, 3, 2)
	pipeline.Run()
	pipeline.Shutdown()

	time.Sleep(500 * time.Millisecond)

	// Medium scale test
	fmt.Println("\n" + strings.Repeat("=", 60))
	pipeline = NewScalablePipeline(20, 10, 5)
	pipeline.Run()
	pipeline.Shutdown()

	// For large scale (50-200 sources), uncomment below:
	// time.Sleep(500 * time.Millisecond)
	// fmt.Println("\n" + strings.Repeat("=", 60))
	// pipeline = NewScalablePipeline(100, 20, 10)
	// pipeline.Run()
	// pipeline.Shutdown()
}
