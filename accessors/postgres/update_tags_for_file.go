package postgres

import "github.com/google/uuid"

func (a *postgresAccessor) UpdateTagsForFile(fileID string, tags []string) error {
    // Start a transaction
    tx, err := a.db.Begin()
    if err != nil {
        return err
    }

    // Defer a rollback in case anything fails
    defer tx.Rollback()

    // Delete existing tag mappings for the file
    _, err = tx.Exec("DELETE FROM file_tag_mappings WHERE file_id = $1", fileID)
    if err != nil {
        return err
    }

    // Insert new tag mappings for the file
    for _, tag := range tags {
        tagUUID := uuid.New()
        var tagID string
        err = tx.QueryRow("INSERT INTO tags (id, tag) VALUES ($1, $2) ON CONFLICT (tag) DO UPDATE SET tag = EXCLUDED.tag RETURNING id;", tagUUID, tag).Scan(&tagID)
        if err != nil {
            return err
        }

        _, err = tx.Exec("INSERT INTO file_tag_mappings (file_id, tag_id) VALUES ($1, $2) ON CONFLICT DO NOTHING;", fileID, tagID)
        if err != nil {
            return err
        }
    }

    // Cleanup unused tags
    _, err = tx.Exec("DELETE FROM tags WHERE id NOT IN (SELECT tag_id FROM file_tag_mappings)")
    if err != nil {
        return err
    }

    // Commit the transaction
    if err = tx.Commit(); err != nil {
        return err
    }

    return nil
}
