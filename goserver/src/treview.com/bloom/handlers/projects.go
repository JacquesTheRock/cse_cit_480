package handlers

import (
	"net/http"
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
	w.WriteHeader(http.StatusOK)
}
func postProjects(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
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
	w.WriteHeader(http.StatusOK)
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
	w.WriteHeader(http.StatusOK)
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
	w.WriteHeader(http.StatusOK)
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
	w.WriteHeader(http.StatusOK)
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
	w.WriteHeader(http.StatusOK)
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
	w.WriteHeader(http.StatusOK)
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
	w.WriteHeader(http.StatusOK)
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
