package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("Just problem case always retunr 0")
	go problem()
	wg.Add(1)
	fmt.Println("Solution of the problem with atomic sync.")
	go solution()
	wg.Add(1)
	fmt.Println("Solution of the problem with mutex")
	go SolutionMutex()
	wg.Add(1)

	wg.Wait()
	fmt.Println("Finish")
}

// problem that there are no sync between main func and gorotune
// so print will be executed before increasing happen.
func problem() {
	runtime.GOMAXPROCS(1)

	var i int
	go func() {
		for {
			i++
		}
	}()

	fmt.Println(i)
	wg.Done()
}

func solution() {
	runtime.GOMAXPROCS(1)

	var i int64
	go func() {
		for {
			atomic.AddInt64(&i, 1)
		}
	}()

	time.Sleep(2 * time.Millisecond)
	fmt.Println(atomic.LoadInt64(&i))
	wg.Done()
}

func SolutionMutex() {
	runtime.GOMAXPROCS(1)

	var i int
	var mu sync.Mutex

	go func() {
		for {
			mu.Lock()
			i++
			mu.Unlock()
		}
	}()

	time.Sleep(time.Millisecond)

	mu.Lock()
	fmt.Println(i)
	mu.Unlock()
	wg.Done()
}
