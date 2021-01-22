package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	LayoutDir   = "views/layouts/"
	TemplateExt = ".html"
)

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

//NewView creates one view having footer.
func NewView(layout string, files ...string) *View {

	files = append(files, layoutFiles()...)
	template, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: template,
		Layout:   layout,
	}
}

//View is a structure having one value Template as *template.Template
type View struct {
	Template *template.Template
	Layout   string
}

//Render is a render function.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
