package handlers

import (
	"fmt"
	"net/http"
)

func UsersRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"This is the Users Root API")
}
