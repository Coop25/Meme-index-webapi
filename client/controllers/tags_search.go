package controllers

import (
	"net/http"
	restapi "github.com/Coop25/the-meme-index-api/gen/openapi/memeapi"
)

// List files with pagination and tag filtering
// (GET /tags/search)
func (c *Controller) GetTagsSearch(w http.ResponseWriter, r *http.Request, params restapi.GetTagsSearchParams) {
	panic("not implemented") // TODO: Implement
}
