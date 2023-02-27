package blogposts_test

import (
	"fmt"
	blogposts "go-tests/reading_files"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct{}

func (StubFailingFS) Open(name string) (fs.File, error) {
	return nil, fmt.Errorf("always failing")
}

func TestNewBlogPost(t *testing.T) {
	t.Run("filesystem title, description", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
			secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
Deploi
Fordabl`
		)

		fs := fstest.MapFS{
			"hello_world.md":  {Data: []byte(firstBody)},
			"hello_world2.md": {Data: []byte(secondBody)},
		}

		posts, _ := blogposts.NewPostFromFS(fs)
		got := posts[0]
		want := blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		}
		assertPost(t, got, want)
	})
}

func assertPost(t testing.TB, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}
