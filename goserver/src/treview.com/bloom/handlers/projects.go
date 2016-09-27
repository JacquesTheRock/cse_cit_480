package handlers

import (
	"strconv"
	"encoding/json"
	"net/http"
	"treview.com/bloom/entity"
	"treview.com/bloom/util"
	authlib "treview.com/bloom/auth"
)

func Projects(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjects(w, r)
	case "POST":
		postProjects(w, r)
	}
}
func getProjects(w http.ResponseWriter, r *http.Request) {
	//TODO: List all projects
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	p := [10]entity.Project{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(p)
}
func postProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	token := r.Header.Get("Authorization")
	p := entity.Project{}
	if authlib.VerifyPermissions(token) {
		r.ParseForm()
		name := r.FormValue("name")
		desc := r.FormValue("description")
		visible,err := strconv.ParseBool(r.FormValue("public"))
		if err != nil{
			util.PrintError(err)
			w.WriteHeader(400)
			return
		}
		p,err = entity.NewProject(name, desc, visible)
	} else {
		util.PrintInfo("User Access denied")
	}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(p)
	
}

func ProjectsPid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPid(w, r)
	case "PUT":
		putProjectsPid(w, r)
	case "DELETE":
		deleteProjectsPid(w, r)
	}
}
func getProjectsPid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	p := entity.Project{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(p)
}
func putProjectsPid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func deleteProjectsPid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ProjectsPidTraits(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidTraits(w, r)
	case "POST":
		postProjectsPidTraits(w, r)
	}
}
func getProjectsPidTraits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	t := [10]entity.Trait{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(t)
}
func postProjectsPidTraits(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ProjectsPidTraitsTid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidTraitsTid(w, r)
	case "POST":
		putProjectsPidTraitsTid(w, r)
	case "DELETE":
		deleteProjectsPidTraitsTid(w, r)
	}
}
func getProjectsPidTraitsTid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	t := entity.Trait{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(t)
}
func putProjectsPidTraitsTid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func deleteProjectsPidTraitsTid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func ProjectsPidCrosses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidCrosses(w, r)
	case "POST":
		postProjectsPidCrosses(w, r)
	}
}
func getProjectsPidCrosses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	c := [10]entity.Cross{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(c)
}
func postProjectsPidCrosses(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ProjectsPidCrossesCid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidCrossesCid(w, r)
	case "PUT":
		putProjectsPidCrossesCid(w, r)
	case "DELETE":
		deleteProjectsPidCrossesCid(w, r)
	}
}
func getProjectsPidCrossesCid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	c := entity.Cross{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(c)
}
func putProjectsPidCrossesCid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func deleteProjectsPidCrossesCid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ProjectsPidCrossesCidCandidates(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidCrossesCidCandidates(w, r)
	case "POST":
		postProjectsPidCrossesCidCandidates(w, r)
	}
}
func getProjectsPidCrossesCidCandidates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	c := [10]entity.Candidate{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(c)
}
func postProjectsPidCrossesCidCandidates(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ProjectsPidCrossesCidCandidatesCnid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsPidCrossesCidCandidatesCnid(w, r)
	case "PUT":
		putProjectsPidCrossesCidCandidatesCnid(w, r)
	case "DELETE":
		deleteProjectsPidCrossesCidCandidatesCnid(w, r)
	}
}
func getProjectsPidCrossesCidCandidatesCnid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	c := entity.Candidate{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(c)
}
func putProjectsPidCrossesCidCandidatesCnid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func deleteProjectsPidCrossesCidCandidatesCnid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ProjectsPidTreview(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsTreview(w, r)
	}
}
func getProjectsTreview(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ProjectsPidTreviewCid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProjectsTreviewCid(w, r)
	}
}
func getProjectsTreviewCid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
