package blogrenderer_test

import (
	"bytes"
	blogrenderer "go-tests/templating"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

var (
	aPost = blogrenderer.PostViewModel{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}
)

func TestRender(t *testing.T) {
	t.Run("converts single post to HTML", func(t *testing.T) {
		var buf bytes.Buffer
		postRender, err := blogrenderer.NewPostRender()
		if err != nil {
			t.Fatal(err)
		}

		if err := postRender.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})

	t.Run("renders an index of posts", func(t *testing.T) {
		var buf bytes.Buffer
		posts := []blogrenderer.PostViewModel{
			{Title: "Hello World"},
			{Title: "Hello World 2"},
		}
		postRender, err := blogrenderer.NewPostRender()
		if err != nil {
			t.Fatal(err)
		}

		if err := postRender.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<ol><li><a href="/post/hello-world">Hello World</a></li><li><a href="/post/hello-world-2">Hello World 2</a></li></ol>`

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}

func BenchmarkRender(b *testing.B) {
	postRender, err := blogrenderer.NewPostRender()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRender.Render(io.Discard, aPost)
	}
}
