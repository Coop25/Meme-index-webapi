package postgres

import (
	"github.com/google/uuid"
)

type PostgresAccessor interface {
	UploadMeme(in NewMeme) (string, error)
	GetMemeById(id string) (Meme, error)
	ListAllMemes(page int, limit int) (ListMemes, error)
	RandomMeme() (Meme, error)
	DeleteMeme(id string) error

	ListAllTags(page int, limit int) (ListTags, error)
	SearchFilesByTags(tags []string, page int, limit int) (ListMemes, error)
	// returns a top 50 list of tags that match the query
	AutocompleteTags(query string) ([]string, error)
	AddTagToFile(memeId string, tag string) error
	RemoveTagFromFile(memeId string, tag string) error
	UpdateTagsForFile(fileID string, tags []string) error
}

type ListTags struct {
	Tags       []string
	Page       int
	TotalPages int
}

type ListMemes struct {
	Memes      []Meme
	Page       int
	TotalCount int
}

type NewMeme struct {
	Id          uuid.UUID
	Name        string
	Tags        []string
	Url         string
	ContentType string
	Description string
}

type Meme struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ContentType string    `json:"contenttype"`
	Tags        []string  `json:"tags"`
	Url         *string   `json:"url"`
	Description *string   `json:"description"`
}
