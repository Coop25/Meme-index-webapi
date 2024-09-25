package postgres

import "github.com/google/uuid"

func (a *postgresAccessor) AddTagToFile(fileID string, tag string) error {
	// Start a transaction
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}
	tagUUID := uuid.New()
	var tagID string
	err = tx.QueryRow("INSERT INTO tags (id, tag) VALUES ($1, $2) ON CONFLICT (tag) DO UPDATE SET tag = EXCLUDED.tag RETURNING id;", tagUUID, tag).Scan(&tagID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("INSERT INTO file_tag_mappings (file_id, tag_id) VALUES ($1, $2) ON CONFLICT DO NOTHING;", fileID, tagID)
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
