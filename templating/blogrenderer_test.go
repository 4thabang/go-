package blogrenderer_test

import (
	"bytes"
	blogrenderer "go-tests/templating"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

var (
	aPost = blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}
)

func TestRender(t *testing.T) {
	t.Run("converts single post to HTML", func(t *testing.T) {
		var buf bytes.Buffer
		templ, err := blogrenderer.NewPostRender()
		if err != nil {
			t.Fatal(err)
		}

		if err := templ.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	templ, err := blogrenderer.NewPostRender()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		templ.Render(io.Discard, aPost)
	}
}
