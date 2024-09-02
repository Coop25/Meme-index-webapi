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
	files, err := c.managers.Files.ListAllMemes(*params.Page, *params.PageSize)
	if err != nil {
		log.Printf("Unable to list files: %v", err)
		http.Error(w, "Unable to list files", http.StatusInternalServerError)
		return
	}

	// Respond with the files
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(files); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
