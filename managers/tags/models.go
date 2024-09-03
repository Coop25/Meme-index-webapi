package tags

import "github.com/Coop25/the-meme-index-api/managers/files"

type TagsManager interface {
	RemoveTagFromFile(id string, tags []string) error
	AddTagsToFile(id string, tags []string) (files.Meme, error)
	AutocompleteTags(tag string) ([]string, error)
	ListAllTags(page int, limit int) (ListTags, error)
	SearchFilesByTags(tags []string, page int, limit int) (SearchTags, error)
}

type ListTags struct {
	Tags       []string
	Page       int
	TotalPages int
}

type SearchTags struct {
	Memes      []files.Meme `json:"memes"`
	Page       int          `json:"page"`
	PageCount  int          `json:"page_count"`
	InputTags  []string     `json:"input_tags"`
}
