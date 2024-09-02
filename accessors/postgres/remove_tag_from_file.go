package postgres

func (a *postgresAccessor) RemoveTagFromFile(fileID string, tag string) error {
	// Start a transaction
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	// Retrieve the tag ID
	var tagID string
	err = tx.QueryRow("SELECT id FROM tags WHERE tag = $1;", tag).Scan(&tagID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Remove the tag from the file
	_, err = tx.Exec("DELETE FROM file_tags WHERE file_id = $1 AND tag_id = $2;", fileID, tagID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
