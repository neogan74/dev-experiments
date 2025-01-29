package main

import (
	"log"
	"os"

	"github.com/neogan74/twg/reading"
)

func main() {
	posts, err := reading.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
