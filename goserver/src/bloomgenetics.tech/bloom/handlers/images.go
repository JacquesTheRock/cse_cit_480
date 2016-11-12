package handlers

import (
	"bloomgenetics.tech/bloom/code"
	"bloomgenetics.tech/bloom/entity"
	"bloomgenetics.tech/bloom/images"
	"bloomgenetics.tech/bloom/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Images(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		postImages(w, r)
	}
}
func postImages(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	img := entity.Image{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&img)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Unable to decode json"
		util.PrintDebug(err)
		util.PrintError("Unable to decode image")
	}
	if out.Code == 0 {
		out.Data, err = images.CreateImage(img)
		if err != nil {
			out.Code = code.UNDEFINED
			out.Status = "Error when writing image"
		}
	}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}

func ImagesIid(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getImagesIid(w, r)
	case "DELETE":
		deleteImagesIid(w, r)
	}
}

func getImagesIid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	iid, err := strconv.ParseInt(vars["iid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a numeric Image ID"
	}
	if out.Code == 0 {
		out.Code = iid
	}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
func deleteImagesIid(w http.ResponseWriter, r *http.Request) {
	out := entity.ApiData{}
	vars := mux.Vars(r)
	iid, err := strconv.ParseInt(vars["iid"], 10, 64)
	if err != nil {
		out.Code = code.INVALIDFIELD
		out.Status = "Not a numeric Image ID"
	}
	if out.Code == 0 {
		out.Code = iid
	}
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(out)
}
