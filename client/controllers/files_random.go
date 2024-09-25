package controllers

import (
	"encoding/json"
	"net/http"
)

// Get a random file
// (GET /files/random)
func (c *Controller) GetFilesRandom(w http.ResponseWriter, r *http.Request) {
	meme, err := c.managers.Files.RandomMeme()
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusInternalServerError)
		return
	}

	// Respond with the file
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(toMemeResponse(meme)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
