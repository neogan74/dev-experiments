package reading

import (
	"bufio"
	"io/fs"
	"testing/fstest"
)

type Post struct {
	Title       string
	Description string
}

func NewPostsFromFS(fileSystem fstest.MapFS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}
func newPost(postFile fs.File) (Post, error) {
	scanner := bufio.NewScanner(postFile)
	scanner.Scan()
	titleLine := scanner.Text()

	scanner.Scan()
	descriptionLine := scanner.Text()

	post := Post{Title: titleLine[7:],
		Description: descriptionLine[14:]}
	return post, nil
}
