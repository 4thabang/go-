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
	/*
	 *  t.Run("2 files in filesystem", func(t *testing.T) {
	 *    fs := fstest.MapFS{
	 *      "hello_world.md":  {Data: []byte("hi")},
	 *      "hello_world2.md": {Data: []byte("hola")},
	 *    }
	 *
	 *    posts, err := blogposts.NewPostFromFS(fs)
	 *
	 *    if err != nil {
	 *      t.Fatal(err)
	 *    }
	 *
	 *    if len(posts) != len(fs) {
	 *      t.Errorf("got: %d posts, want: %d posts", len(posts), len(fs))
	 *    }
	 *  })
	 */

	t.Run("using filesystem fields", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello_world.md":  {Data: []byte("Title: Post 1")},
			"hello_world2.md": {Data: []byte("Title: Post 2")},
		}

		posts, _ := blogposts.NewPostFromFS(fs)
		got := posts[0]
		want := blogposts.Post{Title: "Post 1"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %+v, want: %+v", got, want)
		}
	})
}
