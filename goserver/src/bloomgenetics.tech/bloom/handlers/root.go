package handlers

import (
	"fmt"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the API Root")
}
