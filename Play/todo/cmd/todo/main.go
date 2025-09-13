package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/neogan74/dev-experiments/Play/todo/pkg/todo"
)

const todoFile = ".todo.json"

func main() {
	l := &todo.List{}
	if err := l.Get(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)

		if err := l.Save(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

}
