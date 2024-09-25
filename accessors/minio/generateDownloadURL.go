package minio

import (
    "net/url"
    "time"
)

func (m *minioAccessor) GenerateDownloadURL(fileName string) (string, error) {
    // Set expiration duration to 1 hour
    expiration := time.Hour

    // Define custom request parameters
    reqParams := make(url.Values)
    reqParams.Set("response-content-disposition", "attachment; filename=\""+fileName+"\"")

    // Generate a pre-signed URL for the specified file with custom request parameters
    presignedURL, err := m.minio.PresignedGetObject(m.config.MinioBucketName, fileName, expiration, reqParams)
    if err != nil {
        return "", err
    }
    return presignedURL.String(), nil
}