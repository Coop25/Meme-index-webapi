package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// Add tags to a file by ID
// (PATCH /files/{id}/tags)
func (c *Controller) PatchFilesIdTags(w http.ResponseWriter, r *http.Request, id string) {
	// Define a struct to hold the tags
	type TagsRequest struct {
		Tags []string `json:"tags"`
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Unmarshal the request body into the TagsRequest struct
	var tagsRequest TagsRequest
	err = json.Unmarshal(body, &tagsRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	meme, err := c.managers.Tags.AddTagsToFile(id, tagsRequest.Tags)
	if err != nil {
		http.Error(w, "Unable to add tags to file", http.StatusInternalServerError)
		return
	}

	// Respond with the file
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(meme); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Remove tags from a file by ID
// (DELETE /files/{id}/tags)
func (c *Controller) DeleteFilesIdTags(w http.ResponseWriter, r *http.Request, id string) {
	// Define a struct to hold the tags
	type TagsRequest struct {
		Tags []string `json:"tags"`
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Unmarshal the request body into the TagsRequest struct
	var tagsRequest TagsRequest
	err = json.Unmarshal(body, &tagsRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := c.managers.Tags.RemoveTagFromFile(id, tagsRequest.Tags); err != nil {
		http.Error(w, "Unable to remove tags from file", http.StatusInternalServerError)
		return
	}

	// Send a 204 No Content status
	w.WriteHeader(http.StatusNoContent)
}
