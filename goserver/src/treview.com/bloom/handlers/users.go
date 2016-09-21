package handlers

import (
	"net/http"
)

func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsers(w, r)
	case "POST":
		postUsers(w, r)
	}
}
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func postUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func UsersUid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsersUid(w, r)
	case "PUT":
		putUsersUid(w, r)
	}
}
func getUsersUid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func putUsersUid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func UsersUidProjects(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsersUidProjects(w, r)
	}
}
func getUsersUidProjects(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func UsersUidMail(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsersUidMail(w, r)
	case "POST":
		postUsersUidMail(w, r)
	}
}
func getUsersUidMail(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func postUsersUidMail(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func UsersUidMailMid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsersUidMailMid(w, r)
	case "PUT":
		putUsersUidMailMid(w, r)
	case "DELETE":
		deleteUsersUidMailMid(w, r)
	}
}
func getUsersUidMailMid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func putUsersUidMailMid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func deleteUsersUidMailMid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
