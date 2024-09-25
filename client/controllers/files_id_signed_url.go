package controllers

import (
	"encoding/json"
	"net/http"
)

// Delete a file by ID
// (DELETE /files/{id})
func (c *Controller) GetFilesIdSignedUrl(w http.ResponseWriter, r *http.Request, id string) {
	signedUrl, err := c.managers.Files.GetSignedURL(id)
	if err != nil {
		http.Error(w, "Unable to get signed URL", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	response := map[string]interface{}{
		"signedUrl": signedUrl,
	}

	// Respond with the signed URL
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
