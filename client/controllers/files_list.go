package controllers

import (
	"net/http"

	restapi "github.com/Coop25/the-meme-index-api/gen/openapi/memeapi"
)

// List files with pagination
// (GET /files/list)
func (c *Controller) GetFilesList(w http.ResponseWriter, r *http.Request, params restapi.GetFilesListParams) {
	panic("not implemented") // TODO: Implement
}
