package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	restapi "github.com/Coop25/the-meme-index-api/gen/openapi/memeapi"
)

// Autocomplete tags
// (GET /tags/autocomplete)
func (c *Controller) GetTagsAutocomplete(w http.ResponseWriter, r *http.Request, params restapi.GetTagsAutocompleteParams) {
	tags, err := c.managers.Tags.AutocompleteTags(params.Query)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Unable to autocomplete tags", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the tags array to JSON and write it to the response
	if err := json.NewEncoder(w).Encode(tags); err != nil {
		http.Error(w, "Unable to encode response", http.StatusInternalServerError)
	}
}
