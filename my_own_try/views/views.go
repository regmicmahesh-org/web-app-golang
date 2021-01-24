package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	LAYOUT_DIR = "views/layouts/"
	TEMP_EXT   = ".html"
)

type View struct {
	Template *template.Template
	Layout   string
}

func getTemplates() []string {
	layouts, err := filepath.Glob(LAYOUT_DIR + "*" + TEMP_EXT)
	if err != nil {
		panic(err)
	}
	return layouts
}

func New(layout string, files ...string) *View {

	templates := getTemplates()
	files = append(files, templates...)

	parsedTemplate, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: parsedTemplate,
		Layout:   layout,
	}

}

func (v *View) Render(wr http.ResponseWriter, data interface{}) error {
	return (v.Template.ExecuteTemplate(wr, v.Layout, data))
}
