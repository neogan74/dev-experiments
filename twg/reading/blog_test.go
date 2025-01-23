package main

import (
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello.md":  {Data: []byte("hi")},
		"hello2.md": {Data: []byte("hola")},
	}

	posts := NewPostsFromFS(fs)
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
