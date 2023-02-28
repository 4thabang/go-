package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"
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

func (p *PostRender) RenderIndex(w io.Writer, posts []Post) error {
	indexTempl := `<ol>{{range .}}<li><a href="/post/{{sanitiseTitle .Title}}">{{.Title}}</a></li>{{end}}</ol>`
	templ, err := template.New("index").Funcs(template.FuncMap{
		"sanitiseTitle": func(title string) string {
			return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
		},
	}).Parse(indexTempl)
	if err != nil {
		return err
	}

	if templ.Execute(w, posts); err != nil {
		return err
	}
	return nil
}
