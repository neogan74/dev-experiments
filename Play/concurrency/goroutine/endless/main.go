package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("Just problem case always retunr 0")
	problem()
	fmt.Println("Solution of the problem with atomic sync.")
	solution()
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
}
