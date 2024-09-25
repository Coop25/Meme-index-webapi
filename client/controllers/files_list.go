package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	restapi "github.com/Coop25/the-meme-index-api/gen/openapi/memeapi"
)

// List files with pagination
// (GET /files/list)
func (c *Controller) GetFilesList(w http.ResponseWriter, r *http.Request, params restapi.GetFilesListParams) {
	page := 1
	if params.Page != nil {
		page = *params.Page
	}

	pageSize := 50
	if params.PageSize != nil {
		pageSize = *params.PageSize
	}
	files, err := c.managers.Files.ListAllMemes(page, pageSize)
	if err != nil {
		log.Printf("Unable to list files: %v", err)
		http.Error(w, "Unable to list files", http.StatusInternalServerError)
		return
	}

	// Respond with the files
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(toFileListResponse(files)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
