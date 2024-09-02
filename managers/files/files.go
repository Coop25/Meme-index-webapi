package files

import (
	"github.com/Coop25/the-meme-index-api/accessors"
	"github.com/Coop25/the-meme-index-api/config"
)

type fileManager struct {
	config    *config.Config
	accessors *accessors.Accessor
}

func New(config *config.Config, accessors *accessors.Accessor) *fileManager {
	return &fileManager{
		config:    config,
		accessors: accessors,
	}
}
