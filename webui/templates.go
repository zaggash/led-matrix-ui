package webui

import (
	"embed"
	"html/template"
	"io/fs"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-contrib/multitemplate"
)

// Tweak function to fix multitemplate which do not support embedFS
// https://github.com/gin-contrib/multitemplate/issues/30
func LoadTemplates(e embed.FS) multitemplate.Renderer {
	templatesDir := "templates/"
	r := multitemplate.NewRenderer()

	funcMap := template.FuncMap{
		"formatAsDate": func(t time.Time) string {
			return t.Format("Jan 2, 2006")
		},
	}

	layouts, err := fs.Glob(e, templatesDir+"layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	partials, err := fs.Glob(e, templatesDir+"partials/*.html")
	if err != nil {
		panic(err.Error())
	}

	views, err := fs.Glob(e, templatesDir+"views/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts, partials, and views directories
	for _, view := range views {
		assets := []string{}
		assets = append(assets, layouts...)
		assets = append(assets, partials...)
		files := append(assets, view)

		// should be same name as the root file so that we don't get "incomplete" template error
		tname := filepath.Base(files[0])
		t := template.Must(template.New(tname).Funcs(funcMap).ParseFS(
			e,
			files...,
		))

		fileName := filepath.Base(view)
		templateName := strings.TrimSuffix(fileName, ".html")
		r.Add(templateName, t)
	}
	return r
}
