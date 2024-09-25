package files

func (m *fileManager) GetSignedURL(id string) (string, error) {
	// Get the file by ID
	file, err := m.accessors.Postgres.GetMemeById(id)
	if err != nil {
		return "", err
	}
	
	// Generate a signed URL for the file
	signedURL, err := m.accessors.Minio.GenerateDownloadURL(file.Name)
	if err != nil {
		return "", err
	}

	return signedURL, nil
}