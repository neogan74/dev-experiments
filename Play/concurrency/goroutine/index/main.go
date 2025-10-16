package main

import (
	"fmt"
	"time"
)

// in this experiment we have result wit number in unexpeceted oreder
// For example:
// 4
// 3
// 0
// 2
// 1

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(2 * time.Second)
}
