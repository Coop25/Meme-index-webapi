package accessors

import (
	"github.com/Coop25/the-meme-index-api/accessors/minio"
	"github.com/Coop25/the-meme-index-api/accessors/postgres"
	"github.com/Coop25/the-meme-index-api/config"
)

type Accessor struct {
	Postgres postgres.PostgresAccessor
	Minio    minio.MinioAccessor
}

func New(config *config.Config) *Accessor {
	return &Accessor{
		Postgres: postgres.New(config),
		Minio:    minio.New(config),
	}
}
