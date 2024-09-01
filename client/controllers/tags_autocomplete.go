package controllers

import (
	"net/http"
	restapi "github.com/Coop25/the-meme-index-api/gen/openapi/memeapi"
)

// Autocomplete tags
// (GET /tags/autocomplete)
func (c *Controller) GetTagsAutocomplete(w http.ResponseWriter, r *http.Request, params restapi.GetTagsAutocompleteParams) {
	panic("not implemented") // TODO: Implement
}
