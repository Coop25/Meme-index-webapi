package postgres

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func (a *postgresAccessor) UploadMeme(in NewMeme) (string, error) {
	// Start a transaction
	tx, err := a.db.Begin()
	if err != nil {
		return "", err
	}

	// Step 1: Insert file data
	query, values, err := buildInsertQuery(in)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	// Execute the insert query
	var fileID string
	err = tx.QueryRow(query, values...).Scan(&fileID)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	// Step 2: Insert tags and relationships
	for _, tag := range in.Tags {
		tagUUID := uuid.New()
		var tagID string
		err := tx.QueryRow("INSERT INTO tags (id, tag) VALUES ($1, $2) ON CONFLICT (tag) DO UPDATE SET tag = EXCLUDED.tag RETURNING id;", tagUUID, tag).Scan(&tagID)
		if err != nil {
			tx.Rollback()
			return "", err
		}

		_, err = tx.Exec("INSERT INTO file_tags (file_id, tag_id) VALUES ($1, $2) ON CONFLICT DO NOTHING;", fileID, tagID)
		if err != nil {
			tx.Rollback()
			return "", err
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return "", err
	}

	return fileID, nil
}

// buildInsertQuery dynamically builds the insert query based on content, URL, and tags presence
func buildInsertQuery(in NewMeme) (string, []interface{}, error) {
	var queryBuilder strings.Builder
	var values []interface{}

	queryBuilder.WriteString("INSERT INTO files (id, name, contenttype")
	values = append(values, in.Id)
	values = append(values, in.Name)
	values = append(values, in.ContentType)

	if in.Url != "" {
		queryBuilder.WriteString(", url")
		values = append(values, in.Url)
	}

	if in.Description != "" {
		queryBuilder.WriteString(", description")
		values = append(values, in.Description)
	}

	queryBuilder.WriteString(") VALUES (")

	for i := 1; i <= len(values); i++ {
		queryBuilder.WriteString("$" + strconv.Itoa(i))
		if i < len(values) {
			queryBuilder.WriteString(", ")
		}
	}

	queryBuilder.WriteString(") RETURNING id;")

	return queryBuilder.String(), values, nil
}
