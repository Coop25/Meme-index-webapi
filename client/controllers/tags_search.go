package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	restapi "github.com/Coop25/the-meme-index-api/gen/openapi/memeapi"
)

// List files with pagination and tag filtering
// (GET /tags/search)
func (c *Controller) GetTagsSearch(w http.ResponseWriter, r *http.Request, params restapi.GetTagsSearchParams) {
	// Convert the URL-encoded array of strings back into a slice
	inTags := strings.Split(params.Tags, ",")

	tags, err := c.managers.Tags.SearchFilesByTags(inTags, *params.Page, *params.PageSize)
	if err != nil {
		http.Error(w, "Unable to search tags", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the tags array to JSON and write it to the response
	if err := json.NewEncoder(w).Encode(tags); err != nil {
		http.Error(w, "Unable to encode response", http.StatusInternalServerError)
	}
}
