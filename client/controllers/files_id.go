package controllers

import (
	"encoding/json"
	"net/http"
)

// Delete a file by ID
// (DELETE /files/{id})
func (c *Controller) DeleteFilesId(w http.ResponseWriter, r *http.Request, id string) {
	if err := c.managers.Files.DeleteFileById(id); err != nil {
		http.Error(w, "Unable to delete file", http.StatusInternalServerError)
		return
	}

	// Send a 204 No Content status
	w.WriteHeader(http.StatusNoContent)
}

// Get a file by ID
// (GET /files/{id})
func (c *Controller) GetFilesId(w http.ResponseWriter, r *http.Request, id string) {
	meme, err := c.managers.Files.GetMeme(id)
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusInternalServerError)
		return
	}

	// Respond with the file
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(meme); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
