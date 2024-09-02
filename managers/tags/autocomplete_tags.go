package tags

func (m *tagsManager) AutocompleteTags(tag string) ([]string, error) {
	tags, err := m.accessors.Postgres.AutocompleteTags(tag)
	if err != nil {
		return []string{}, err
	}
	return tags, nil
}
