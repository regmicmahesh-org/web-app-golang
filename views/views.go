package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

//View :This is the base of view
type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layoutName string, files ...string) *View {

	layoutFiles, err := filepath.Glob("views/layouts/*.html")
	if err != nil {
		panic(err)
	}
	files = append(files, layoutFiles...)

	templ, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: templ,
		Layout:   layoutName,
	}

}

//Render Function
//This function renders on screen.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return (v.Template.ExecuteTemplate(w, v.Layout, data))
}
