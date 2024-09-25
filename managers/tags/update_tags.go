package tags

func (m *tagsManager) UpdateTagsForFile(fileID string, tags []string) error {
	if err := m.accessors.Postgres.UpdateTagsForFile(fileID, tags); err != nil {
		return err
	}
	return nil
}
