package minio

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
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

	// Compute SHA-256 hash of the file content
	hash := sha256.New()
	if _, err := io.Copy(hash, in.Content); err != nil {
		return MinioMeme{}, err
	}
	hashInBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)

	// Check if a file with the same hash already exists
	exists, err := m.checkIfHashExists(hashString)
	if err != nil {
		return MinioMeme{}, err
	}
	if exists {
		return MinioMeme{}, fmt.Errorf("duplicate file detected")
	}

	// Reset the file content reader to the beginning
	if seeker, ok := in.Content.(io.Seeker); ok {
		if _, err := seeker.Seek(0, io.SeekStart); err != nil {
			return MinioMeme{}, err
		}
	} else {
		return MinioMeme{}, fmt.Errorf("content does not support seeking")
	}

	// Upload the file to MinIO with the UUID filename
	_, err = m.minio.PutObject(m.config.MinioBucketName, uuidFilename, in.Content, -1, minio.PutObjectOptions{
		ContentType: contentType,
		UserMetadata: map[string]string{
			"X-Amz-Meta-File-Hash": hashString,
		},
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
func (m *minioAccessor) checkIfHashExists(hash string) (bool, error) {
	// List all objects in the bucket and check their metadata for the hash
	doneCh := make(chan struct{})
	defer close(doneCh)

	objectCh := m.minio.ListObjectsV2(m.config.MinioBucketName, "", true, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			return false, object.Err
		}
		// Get object metadata
		objInfo, err := m.minio.StatObject(m.config.MinioBucketName, object.Key, minio.StatObjectOptions{})
		if err != nil {
			return false, err
		}
		// Access user-defined metadata
		if fileHash, ok := objInfo.Metadata["X-Amz-Meta-File-Hash"]; ok {
			if fileHash[0] == hash {
				return true, nil
			}
		}
	}
	return false, nil
}

func getContentType(extension string) (string, error) {
	extension = strings.ToLower(strings.TrimPrefix(extension, "."))
	ContentType, exists := ContentTypes[extension]
	if !exists {
		return "", fmt.Errorf("unknown file extension: %s", extension)
	}
	return ContentType, nil
}
