package main

import (
	"log"
	"net/http"
	"github.com/Coop25/the-meme-index-api/config"
	restapi "github.com/Coop25/the-meme-index-api/client"
)

func main() {
	var err error
	config := config.LoadConfig()

	api, err := restapi.New(&config)
	if err != nil {
		log.Fatalf("Failed to initialize API: %s", err)
	}

	// Start the server
	log.Fatal(http.ListenAndServe(":"+config.Port, api.Router))
}
