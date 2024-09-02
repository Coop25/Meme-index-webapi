package minio

import (
	"log"

	"github.com/Coop25/the-meme-index-api/config"
	"github.com/minio/minio-go"
)

type minioAccessor struct {
	config *config.Config
	minio  *minio.Client
}

func New(config *config.Config) *minioAccessor {
	minioClient, err := minio.New(config.MinioEndpoint, config.MinioSecretKeyID, config.MinioAccessKey, config.MinioUseSSL)
	if err != nil {
		log.Fatal("Error connecting to minio: ", err)
	}

	return &minioAccessor{
		config: config,
		minio:  minioClient,
	}
}
