package minio

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/minio/minio-go"
)

func (m *minioAccessor) UploadMeme(in NewMeme) (MinioMeme, error) {
	ext := filepath.Ext(in.FileName)
	uuidFilename := fmt.Sprintf("%s%s", in.Id.String(), ext)
	contentType, err := getContentType(ext)
	if ext == "" {
		return MinioMeme{}, err
	}
	// Upload the file to MinIO with the UUID filename
	_, err = m.minio.PutObject(m.config.MinioBucketName, uuidFilename, in.Content, -1, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return MinioMeme{}, err
	}

	return MinioMeme{
		Id:            in.Id,
		FileName:      uuidFilename,
		ContentType:   contentType,
		FileExtension: ext,
	}, nil
}

func getContentType(extension string) (string, error) {
	extension = strings.ToLower(strings.TrimPrefix(extension, "."))
	ContentType, exists := ContentTypes[extension]
	if !exists {
		return "", fmt.Errorf("unknown file extension: %s", extension)
	}
	return ContentType, nil
}
