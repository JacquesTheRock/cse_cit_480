package handlers

import (
	"bloomgenetics.tech/bloom/util"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
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
		util.PrintDebug(err)
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

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Clean("/" + r.URL.Path) //Root the path and clean
	util.PrintInfo(filepath.SplitList(path))
	path = filepath.Join(strings.Split(path, "/")[1:]...)
	path = util.Config.HTMLRoot + "/" + path
	body, err := ioutil.ReadFile(path)
	ext := filepath.Ext(path)
	if err != nil {
		util.PrintError(path)
		util.PrintDebug(err)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	if ext == ".css" {
		w.Header().Set("Content-Type", "text/css")
	}
	w.Write(body)
}
