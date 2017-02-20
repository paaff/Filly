package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/paaff/Filly/files"
)

func main() {

	// Create a simple file server
	fs := http.FileServer(http.Dir("./web/dist"))
	http.Handle("/", fs)

	// GetDir endpoint
	http.HandleFunc("/browse", browseHandler)

	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func browseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		path := r.FormValue("path")
		// Browse from the POST form variable
		cont, status := dirContent.GetDirectoryContentInJSON(path)
		if cont != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(cont)
		} else {
			// Error handling.
			errorHandler(w, r, status)
		}
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == 404 {
		fmt.Fprint(w, "404 - NOT FOUND")
	}
}
