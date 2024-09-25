package postgres

// GetTagsByFileID retrieves tags for a given file ID
func (a *postgresAccessor) GetTagsByFileID(fileID string) ([]string, error) {
	tagQuery := "SELECT t.tag FROM tags t JOIN file_tag_mappings ft ON t.id = ft.tag_id WHERE ft.file_id = $1"
	rows, err := a.db.Query(tagQuery, fileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []string
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}
