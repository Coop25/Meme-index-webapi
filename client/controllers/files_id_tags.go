package controllers

import "net/http"

// Add tags to a file by ID
// (PATCH /files/{id}/tags)
func (c *Controller) PatchFilesIdTags(w http.ResponseWriter, r *http.Request, id string) {
	panic("not implemented") // TODO: Implement
}

// Remove tags from a file by ID
// (DELETE /files/{id}/tags)
func (c *Controller) DeleteFilesIdTags(w http.ResponseWriter, r *http.Request, id string) {
	panic("not implemented") // TODO: Implement
}
