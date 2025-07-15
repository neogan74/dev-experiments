package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(1000)

	value := 0
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			value++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Printf("Result: %v", value)
}
