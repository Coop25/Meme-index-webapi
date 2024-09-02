package minio

func (m *minioAccessor) DeleteMeme(id string) error {
	err := m.minio.RemoveObject(m.config.MinioBucketName, id)
	if err != nil {
		return err
	}

	return nil
}
