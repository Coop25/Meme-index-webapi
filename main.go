package main

import (
	"log"
	"net/http"

	"github.com/Coop25/the-meme-index-api/accessors"
	restapi "github.com/Coop25/the-meme-index-api/client"
	"github.com/Coop25/the-meme-index-api/config"
	"github.com/Coop25/the-meme-index-api/managers"
)

func main() {
	var err error
	config := config.LoadConfig()
	accessors := accessors.New(&config)
	manager := managers.New(&config, accessors)

	api, err := restapi.New(&config, manager)
	if err != nil {
		log.Fatalf("Failed to initialize API: %s", err)
	}

	// Start the server
	log.Fatal(http.ListenAndServe(":"+config.Port, api.Router))
}
