package controllers

import (
	"encoding/json"
	"net/http"

	restapi "github.com/Coop25/the-meme-index-api/gen/openapi/memeapi"
)

// List tags with pagination
// (GET /tags/list)
func (c *Controller) GetTagsList(w http.ResponseWriter, r *http.Request, params restapi.GetTagsListParams) {
	page := 1
	if params.Page != nil {
		page = *params.Page
	}

	pageSize := 50
	if params.PageSize != nil {
		pageSize = *params.PageSize
	}
	tags, err := c.managers.Tags.ListAllTags(page, pageSize)
	if err != nil {
		http.Error(w, "Unable to list tags", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	tagsResponse := restapi.TagListResponse{
		Tags: &tags.Tags,
		Page: &tags.Page,
		PageCount: &tags.PageCount,
	}

	// Encode the tags array to JSON and write it to the response
	if err := json.NewEncoder(w).Encode(tagsResponse); err != nil {
		http.Error(w, "Unable to encode response", http.StatusInternalServerError)
	}
}
