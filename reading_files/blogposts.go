package blogposts

import (
	"io"
	"io/fs"
)

type Post struct {
	Title string
	/*
	 *	Description, Body string
	 *  Tags                     []string
	 */
}

func NewPostFromFS(fileSystem fs.FS) ([]Post, error) {
	var posts []Post
	dir, err := fs.ReadDir(fileSystem, ".")
	if len(dir) == 0 {
		return nil, err
	}
	for _, f := range dir {
		p, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, f string) (Post, error) {
	postFile, err := fileSystem.Open(f)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: string(postData)[7:]}
	return post, nil
}
