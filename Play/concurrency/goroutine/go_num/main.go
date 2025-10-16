package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 15; i++ {
		fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(2 * time.Second)

}
