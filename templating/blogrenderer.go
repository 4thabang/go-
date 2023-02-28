package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
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
	templ    *template.Template
	mdParser *parser.Parser
}

func NewPostRender() (*PostRender, error) {
	t, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	ext := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(ext)

	return &PostRender{
		templ:    t,
		mdParser: parser,
	}, nil
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
}

func (r *PostRender) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", p)
}

func (r *PostRender) RenderIndex(w io.Writer, p []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", p)
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post, r *PostRender) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}
