package view

import (
	"GoMeet/models"
	"GoMeet/session"
	"fmt"
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

type Data struct {
	Alert           *Alert
	SessionUserName string
	SessionUserID   string
	IsAuth          bool
	User            *models.User
	Yield           interface{}
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
func (v *View) Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text-html")
	var renderData Data
	if alert := getAlert(r); alert != nil {
		renderData.Alert = alert
		clearAlert(w)
	}
	username, _ := session.GetUserSession(w, r)
	fmt.Println("****************")
	fmt.Println(username)
	if username != "" {
		renderData.IsAuth = true
	}
	renderData.Yield = data
	if err := v.Template.ExecuteTemplate(w, v.Layout, renderData); err != nil {
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
