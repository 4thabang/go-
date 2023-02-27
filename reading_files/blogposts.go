package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

const (
	titleField = "Title: "
	descField  = "Description: "
	tagsField  = "Tags: "
	bodyField  = "Body: "
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
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
	scanner := bufio.NewScanner(postFile)

	readLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	titleLine := readLine(titleField)
	descriptionLine := readLine(descField)
	tags := readLine(tagsField)
	body := readBody(scanner)

	return Post{
		Title:       titleLine,
		Description: descriptionLine,
		Tags:        strings.Split(tags, ", "),
		Body:        body,
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore a line
	var buf bytes.Buffer
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
