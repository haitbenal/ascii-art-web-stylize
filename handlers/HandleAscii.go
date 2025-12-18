package handlers

import (
	asciiart "asciiart/func"
	"net/http"
)

func HandleAscii(w http.ResponseWriter, r *http.Request) {
	// checking the method of the request
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			HandleErr(w, "Internal Server Error", http.StatusInternalServerError)
			return
		} else {
			form := r.Form
			// cheking if the user manupulate in the form names
			if !form.Has("name") || !form.Has("radio") {
				HandleErr(w, "Bad Request", http.StatusBadRequest)
				return
			} else {
				// extract the data from the form and render the html file and err if existe
				name := r.FormValue("name")
				fName := r.FormValue("radio")
				if fName != "shadow.txt" && fName != "standard.txt" && fName != "thinkertoy.txt" {
					HandleErr(w, "Bad Request", http.StatusBadRequest)
					return
				}
				splited := asciiart.Splite(fName)
				res := asciiart.PrintSymbole(splited, name)
				asciiErr := temp.ExecuteTemplate(w, "ascii-art.html", res)
				if asciiErr != nil {
					HandleErr(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
			}
		}
	} else {
		HandleErr(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

}
