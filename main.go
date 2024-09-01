package main

import (
	"log"
	"net/http"
	"github.com/Coop25/the-meme-index-api/config"
	restapi "github.com/Coop25/the-meme-index-api/client"
)

func main() {
	config := config.LoadConfig()

	api := restapi.New(&config)

	// Start the server
	log.Fatal(http.ListenAndServe(":"+config.Port, api.Router))
}
