package files

func (m *fileManager) DeleteFileById(id string) error {
	// Delete the file from the database
	err := m.accessors.Postgres.DeleteMeme(id)
	if err != nil {
		return err
	}

	// Delete the file from the storage
	err = m.accessors.Minio.DeleteMeme(id)
	if err != nil {
		return err
	}

	return nil
}
