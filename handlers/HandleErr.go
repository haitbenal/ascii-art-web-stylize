package handlers

import (
	"html/template"
	"net/http"
)

var pages = template.Must(template.ParseGlob("./templates/*.html"))

type data struct {
	Code        int
	Description string
}

func HandleErr(w http.ResponseWriter, descriptioin string, statusCode int) {
	// render the err pop up
	info := data{Code: statusCode, Description: descriptioin}
	w.WriteHeader(statusCode)
	pages.ExecuteTemplate(w, "error.html", info)
}
