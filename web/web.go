package web

import (
	"web"
	"io"
	"text/template"
)

//go:embed *
var files embed.FS

var env = parse("env.html")

type EnvParams struct {
	Title   string
	Version string
}

func Env(w io.Writer, p DashboardParams) error {
	return env.Execute(w, p)
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "layout.html", file)
	)
}