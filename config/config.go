package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port             string `envconfig:"PORT" default:"8080"`
	PGDBHost         string `envconfig:"PG_DB_HOST" default:"localhost"`
	PGDBPort         string `envconfig:"PG_DB_PORT" default:"5432"`
	PGDBUser         string `envconfig:"PG_DB_USER" default:"postgres"`
	PGDBPass         string `envconfig:"PG_DB_PASS" default:"password"`
	PGDBName         string `envconfig:"PG_DB_NAME" default:"memeindex"`
	PGDBSSLMode      string `envconfig:"PG_DB_SSL_MODE" default:"disable"`
	MinioEndpoint    string `envconfig:"MINIO_ENDPOINT" default:"localhost:9000"`
	MinioAccessKey   string `envconfig:"MINIO_ACCESS_KEY" default:"minio"`
	MinioSecretKeyID string `envconfig:"MINIO_ACCESS_KEY_ID" default:"minio123"`
	MinioBucketName  string `envconfig:"MINIO_BUCKET_NAME" default:"memeindex"`
	MinioUseSSL      bool   `envconfig:"MINIO_USE_SSL" default:"false"`
	Origin           string `envconfig:"ORIGIN" default:"http://localhost:3000"`
}

func LoadConfig() Config {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}
	return config
}
