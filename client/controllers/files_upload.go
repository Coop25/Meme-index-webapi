package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Coop25/the-meme-index-api/managers/files"
)

// Upload a File
// (POST /files/upload)
func (c *Controller) PostFilesUpload(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data
	err := r.ParseMultipartForm(0)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the file from the form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Extract tags, URL, and description from the form data (adjust field names accordingly)
	tags := strings.Split(r.FormValue("tags"), ",")
	for i, tag := range tags {
		tags[i] = strings.TrimSpace(tag) // Trim spaces around each tag
	}

	// Extract tags, URL, and description from the form data (adjust field names accordingly)
	newMeme := files.UploadFileRequest{
		Filename:    handler.Filename,
		File:        file,
		Tags:        tags,
		URL:         r.FormValue("url"),
		Description: r.FormValue("description"),
	}

	id, err := c.managers.Files.UploadFile(newMeme)
	if err != nil {
		http.Error(w, "Unable to upload file", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	response := map[string]interface{}{
		"fileID":  id,
		"message": "File uploaded successfully, and information stored in the database",
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
