package handlers

import (
	"net/http"
	"os"
	"strings"
)

func HandleForbiden(w http.ResponseWriter, r *http.Request) {
	// make the access to files directory forbidden
	file, err := os.Stat(strings.TrimPrefix(r.URL.Path, "/"))
	if err != nil || file.IsDir() {
		HandleErr(w, "Forbidden", http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, "./files/"+r.URL.Path[len("/files/"):])
}
