package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func UsersRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the Users Root API")
}

func Users(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	//Get the User info at this point
	fmt.Fprintln(w, "This is the Users API"+userId)
}
