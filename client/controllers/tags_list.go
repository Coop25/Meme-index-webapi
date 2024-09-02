package controllers

import (
	"encoding/json"
	"net/http"

	restapi "github.com/Coop25/the-meme-index-api/gen/openapi/memeapi"
)

// List tags with pagination
// (GET /tags/list)
func (c *Controller) GetTagsList(w http.ResponseWriter, r *http.Request, params restapi.GetTagsListParams) {
	tags, err := c.managers.Tags.ListAllTags(*params.Page, *params.PageSize)
	if err != nil {
		http.Error(w, "Unable to list tags", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the tags array to JSON and write it to the response
	if err := json.NewEncoder(w).Encode(tags); err != nil {
		http.Error(w, "Unable to encode response", http.StatusInternalServerError)
	}
}
