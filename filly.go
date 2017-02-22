package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/paaff/Filly/files"
	"github.com/spf13/viper"
)

func main() {
	// Load configuration file
	loadConfig()

	// Create a simple file server
	fs := http.FileServer(http.Dir("./web/dist"))
	http.Handle("/", fs)

	// GetDir endpoint
	http.Handle("/browse", appHandler(browseHandler))

	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

// Loads a config
func loadConfig() {
	viper.SetConfigName("fillyconf")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	} else {
		content.ROOT_DIR = viper.GetString("root_dir")
	}
}

// Custom error created for handling HTTP errors more fluently.
type appError struct {
	Error   error
	Message string
	Code    int
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		http.Error(w, e.Message, e.Code)
	}
}

func browseHandler(w http.ResponseWriter, r *http.Request) *appError {
	if r.Method == "POST" {
		path := r.FormValue("path")
		// Browse from the POST form variable
		if cont, err := content.GetDirectoryContentInJSON(path); err != nil {
			return &appError{err, "Content cannot be found", 404} // Status: Not Found
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(cont)
		}
	}
	return nil
}
