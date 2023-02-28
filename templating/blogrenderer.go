package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

type PostRender struct {
	templ *template.Template
}

func NewPostRender() (*PostRender, error) {
	t, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRender{
		templ: t,
	}, nil
}

func (p *PostRender) Render(w io.Writer, post Post) error {
	if err := p.templ.Execute(w, post); err != nil {
		return err
	}
	return nil
}

func RenderIndex(w io.Writer, posts []Post) error {
	return nil
}
