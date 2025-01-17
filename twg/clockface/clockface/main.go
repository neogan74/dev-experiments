package main

import (
	"os"
	"time"

	"github.com/neogan74/twg/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
