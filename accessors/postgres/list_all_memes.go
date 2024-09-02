package postgres

import (
	"math"

	"github.com/google/uuid"
)

func (a *postgresAccessor) ListAllMemes(page int, limit int) (ListMemes, error) {

	// Count the total number of records
	var totalRecords int
	countQuery := `SELECT COUNT(*) FROM files;`
	err := a.db.QueryRow(countQuery).Scan(&totalRecords)
	if err != nil {
		return ListMemes{}, err
	}

	// Calculate the total number of pages
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	if page > totalPages {
		page = totalPages
	}

	query := `SELECT id, name, contenttype, url, description FROM files ORDER BY name ASC LIMIT $1 OFFSET $2;`

	rows, err := a.db.Query(query, limit, (page-1)*limit)
	if err != nil {
		return ListMemes{}, err
	}
	defer rows.Close()

	memes := []Meme{}
	for rows.Next() {
		var meme Meme
		err = rows.Scan(&meme.Id, &meme.Name, &meme.ContentType, &meme.Url, &meme.Description)
		if err != nil {
			return ListMemes{}, err
		}

		// Skip rows with default zero values for ID
		if meme.Id == uuid.Nil {
			continue
		}

		// Retrieve associated tags using the new function
		tags, err := a.GetTagsByFileID(meme.Id.String())
		if err != nil {
			return ListMemes{}, err
		}
		meme.Tags = tags

		memes = append(memes, meme)
	}

	return ListMemes{
		Memes:      memes,
		Page:       page,
		TotalPages: totalPages,
	}, nil
}
