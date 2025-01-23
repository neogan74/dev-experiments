package main

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

var posts []Post
posts = NewPostsFromFS("some-folder")
