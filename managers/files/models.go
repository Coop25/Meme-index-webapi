package files

import (
	"io"

	"github.com/google/uuid"
)

type FileManager interface {
	UploadFile(in UploadFileRequest) (string, error)
	GetMeme(id string) (Meme, error)
	DeleteFileById(id string) error
	RandomMeme() (Meme, error)
	ListAllMemes(page int, limit int) (ListMemes, error)
	GetSignedURL(id string) (string, error)
}

type UploadFileRequest struct {
	File        io.Reader
	Filename    string
	Tags        []string
	URL         string
	Description string
}

type ListMemes struct {
	Memes     []Meme
	Page      int
	PageCount int
}

type Meme struct {
	ID          uuid.UUID
	Tags        []string
	ContentType string
	FileUrl     string
	URL         *string
	Description *string
}
