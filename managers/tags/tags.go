package tags

import (
	"github.com/Coop25/the-meme-index-api/accessors"
	"github.com/Coop25/the-meme-index-api/config"
)

type tagsManager struct {
	config    *config.Config
	accessors *accessors.Accessor
}

func New(config *config.Config, accessors *accessors.Accessor) *tagsManager {
	return &tagsManager{
		config:    config,
		accessors: accessors,
	}
}
