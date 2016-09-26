package handlers

import (
	"net/http"
	"encoding/json"
	"treview.com/bloom/entity"
)

func Breeds(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getBreeds(w, r)
	case "POST":
		postBreeds(w, r)
	}
}

func getBreeds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	b := [10]entity.Candidate{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(b)
}
func postBreeds(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func BreedsBid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getBreedsBid(w, r)
	case "PUT":
		putBreedsBid(w, r)
	}
}
func getBreedsBid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	b := entity.Candidate{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(b)
}
func putBreedsBid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func BreedsBidTraits(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getBreedsBidTraits(w, r)
	case "POST":
		postBreedsBidTraits(w, r)
	}
}
func getBreedsBidTraits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	t := [10]entity.Trait{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(t)
}
func postBreedsBidTraits(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func BreedsBidTraitsTid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getBreedsBidTraitsTid(w, r)
	case "PUT":
		putBreedsBidTraitsTid(w, r)
	case "DELETE":
		deleteBreedsBidTraitsTid(w, r)
	}
}
func getBreedsBidTraitsTid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//token := r.Header.Get("Authorization")
	t := entity.Trait{}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(t)
}
func putBreedsBidTraitsTid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func deleteBreedsBidTraitsTid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
