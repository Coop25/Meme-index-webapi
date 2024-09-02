package managers

import (
	"github.com/Coop25/the-meme-index-api/accessors"
	"github.com/Coop25/the-meme-index-api/config"
	"github.com/Coop25/the-meme-index-api/managers/files"
	"github.com/Coop25/the-meme-index-api/managers/tags"
)

type Managers struct {
	Tags  tags.TagsManager
	Files files.FileManager
}

func New(config *config.Config, accessors *accessors.Accessor) *Managers {
	return &Managers{
		Tags:  tags.New(config, accessors),
		Files: files.New(config, accessors),
	}
}
