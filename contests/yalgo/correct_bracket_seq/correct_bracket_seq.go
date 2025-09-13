package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	if len(s)%2 == 1 {
		fmt.Println("no")
		return
	}

	stack := make([]byte, 0, len(s))
	match := map[byte]byte{')': '(', ']': '[', '}': '{'}

	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != match[c] {
				fmt.Println("no")
				return
			}
			stack = stack[:len(stack)-1]
		default:
			fmt.Println("no")
			return
		}
	}
	if len(stack) == 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

}
