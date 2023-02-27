package blogrenderer

import (
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(buf io.Writer, p Post) error {
	return nil
}
