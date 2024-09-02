package postgres

import (
	"fmt"
	"strings"
)

func (a *postgresAccessor) SearchFilesByTags(tags []string, page int, limit int) (ListMemes, error) {
	// Start a transaction
	tx, err := a.db.Begin()
	if err != nil {
		return ListMemes{}, err
	}

	// Create the base query
	queryString := "SELECT files.id, files.name, files.contenttype, files.url, files.description FROM files"
	countQuery := "SELECT COUNT(DISTINCT files.id) FROM files"

	if len(tags) > 0 {
		queryString += " JOIN file_tags ON files.id = file_tags.file_id JOIN tags ON file_tags.tag_id = tags.id WHERE"
		countQuery += " JOIN file_tags ON files.id = file_tags.file_id JOIN tags ON file_tags.tag_id = tags.id WHERE"

		tagConditions := make([]string, len(tags))
		for i, _ := range tags {
			tagConditions[i] = fmt.Sprintf("tags.tag = $%d", i+1)
		}
		queryString += " " + strings.Join(tagConditions, " OR ")
		countQuery += " " + strings.Join(tagConditions, " OR ")
	}

	queryString += " GROUP BY files.id, files.name, files.contenttype, files.url, files.description"
	queryString += " ORDER BY files.name ASC LIMIT $%d OFFSET $%d"

	// Add pagination parameters
	offset := (page - 1) * limit
	queryString = fmt.Sprintf(queryString, len(tags)+1, len(tags)+2)

	// Execute the count query
	countStmt, err := tx.Prepare(countQuery)
	if err != nil {
		tx.Rollback()
		return ListMemes{}, err
	}
	defer countStmt.Close()

	// Convert tags to []any
	tagArgs := make([]any, len(tags))
	for i, tag := range tags {
		tagArgs[i] = tag
	}

	var totalCount int
	err = countStmt.QueryRow(tagArgs...).Scan(&totalCount)
	if err != nil {
		tx.Rollback()
		return ListMemes{}, err
	}

	// Execute the main query
	stmt, err := tx.Prepare(queryString)
	if err != nil {
		tx.Rollback()
		return ListMemes{}, err
	}
	defer stmt.Close()

	args := append(tagArgs, limit, offset)
	rows, err := stmt.Query(args...)
	if err != nil {
		tx.Rollback()
		return ListMemes{}, err
	}
	defer rows.Close()

	// Parse the results
	var memes []Meme
	for rows.Next() {
		var meme Meme
		err = rows.Scan(&meme.Id, &meme.Name, &meme.ContentType, &meme.Url, &meme.Description)
		if err != nil {
			tx.Rollback()
			return ListMemes{}, err
		}

		// Retrieve associated tags using the new function
		tags, err := a.GetTagsByFileID(meme.Id.String())
		if err != nil {
			return ListMemes{}, err
		}
		meme.Tags = tags

		memes = append(memes, meme)
	}

	if err := rows.Err(); err != nil {
		tx.Rollback()
		return ListMemes{}, err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return ListMemes{}, err
	}

	return ListMemes{
		Memes:      memes,
		Page:       page,
		TotalPages: totalCount,
	}, nil
}
