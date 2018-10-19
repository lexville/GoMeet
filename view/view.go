package view

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// View contains the template as well as the
// layout
type View struct {
	Template *template.Template
	Layout   string
}

// AddTemplateFiles takes in all the files
// needed to render a view and then returns
// the view struct
func AddTemplateFiles(layout string, files ...string) *View {
	files = append(
		files,
		getLayoutFiles()...,
	)
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal("Unable to parse the view files: ", err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

// Render is responsible for rendering the view with the data needed
// by the view
func (v *View) Render(w http.ResponseWriter, data interface{}) {
	if err := v.Template.ExecuteTemplate(w, v.Layout, data); err != nil {
		log.Fatal("Unable to render the view: ", err)
	}
}

var (
	layoutDirectory = "templates/layouts/"
	layoutExtension = ".gohtml"
)

func getLayoutFiles() []string {
	files, err := filepath.Glob(layoutDirectory + "*" + layoutExtension)
	if err != nil {
		log.Fatal("Unable to get the layout files: ", err)
	}
	return files
}
