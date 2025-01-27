package reading

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: Post 1")},
		"hello-world2.md": {Data: []byte("Title: Post 2")},
	}
	got := posts[0]
	want := Post{Title: "Post 1"}

	if !reflect.DeepEqual(got , want)
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
