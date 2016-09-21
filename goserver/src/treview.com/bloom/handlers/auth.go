package handlers

import (
	//	"fmt"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAuth(w, r)
	case "POST":
		postAuth(w, r)
	case "DELETE":
		deleteAuth(w, r)
	}
}

func getAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func postAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func deleteAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
