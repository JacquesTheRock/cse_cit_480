package handlers

import (
	"bloomgenetics.tech/bloom/util"
	"html/template"
	"net/http"
	"path/filepath"
)

var templates *template.Template

type page struct {
	Name   string
	Author string
}

func initTemplates() {
	templates = template.Must(template.ParseFiles(
		util.Config.TemplateRoot+"/generic/head.tmpl",
		util.Config.TemplateRoot+"/generic/navbar.tmpl",
		util.Config.TemplateRoot+"/generic/footer.tmpl"))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if templates == nil {
		initTemplates()
	}
	path := filepath.Clean("/" + r.URL.Path) //Root the path and clean
	path = util.Config.TemplateRoot + "/content" + path
	t, err := template.ParseFiles(path)
	if err != nil {
		util.PrintError(path)
		util.PrintError(err)
		t, err = template.ParseFiles(util.Config.TemplateRoot + "/404.html")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	templates.ExecuteTemplate(w, "head.tmpl", nil)
	templates.ExecuteTemplate(w, "navbar.tmpl", nil)
	t.Execute(w, nil)
	templates.ExecuteTemplate(w, "footer.tmpl", nil)
}
