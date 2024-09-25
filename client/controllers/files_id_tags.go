package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	restapi "github.com/Coop25/the-meme-index-api/gen/openapi/memeapi"
	"github.com/Coop25/the-meme-index-api/managers/files"
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

	if err := c.managers.Tags.UpdateTagsForFile(id, tagsRequest.Tags); err != nil {
		http.Error(w, "Unable to add tags to file", http.StatusInternalServerError)
		return
	}
	
	// Send a 204 No Content status
	w.WriteHeader(http.StatusNoContent)
}


func toFileListResponse(meme files.ListMemes) restapi.FileListResponse {
	memes := toMemesResponse(meme.Memes)
	return restapi.FileListResponse{
		Memes:     &memes,
		Page:      &meme.Page,
		PageCount: &meme.PageCount,
	}
}

func toMemesResponse(memes []files.Meme) []restapi.Meme {
	memeResponses := []restapi.Meme{}
	for _, meme := range memes {
		memeResponses = append(memeResponses, toMemeResponse(meme))
	}
	return memeResponses
}

func toMemeResponse(meme files.Meme) restapi.Meme {
	id := meme.ID.String()
	return restapi.Meme{
		Id:          &id,
		Tags:        &meme.Tags,
		ContentType: &meme.ContentType,
		FileUrl:     &meme.FileUrl,
		Url:         meme.URL,
		Description: meme.Description,
	}
}
