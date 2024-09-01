package controllers

import (
	"net/http"

	restapi "github.com/Coop25/the-meme-index-api/gen/openapi/memeapi"
)

// List tags with pagination
// (GET /tags/list)
func (c *Controller) GetTagsList(w http.ResponseWriter, r *http.Request, params restapi.GetTagsListParams) {
	panic("not implemented") // TODO: Implement
}
