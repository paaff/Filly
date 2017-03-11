package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/paaff/Filly/api"
	"github.com/paaff/Filly/error"
	"github.com/spf13/viper"
)

func main() {
	// Load configuration file
	loadConfig()

	// Create a simple file server
	fs := http.FileServer(http.Dir("./web/dist"))
	http.Handle("/", fs)

	// GetDir endpoint
	http.Handle("/browse", errorhandler.AppHandler(api.BrowseHandler))

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
		panic(fmt.Errorf("Fatal error config file: %s", err))
	} else {
		api.RootDir = viper.GetString("root_dir")
	}
}
