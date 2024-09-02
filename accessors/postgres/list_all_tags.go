package postgres

import "math"

func (a *postgresAccessor) ListAllTags(page int, limit int) (ListTags, error) {

	// Count the total number of records
	var totalRecords int
	countQuery := `SELECT COUNT(*) FROM tags;`
	err := a.db.QueryRow(countQuery).Scan(&totalRecords)
	if err != nil {
		return ListTags{}, err
	}

	// Calculate the total number of pages
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	if page > totalPages {
		page = totalPages
	}

	query := `SELECT tag FROM tags ORDER BY tag ASC LIMIT $1 OFFSET $2;`

	rows, err := a.db.Query(query, limit, (page-1)*limit)
	if err != nil {
		return ListTags{}, err
	}
	defer rows.Close()

	tags := []string{}
	for rows.Next() {
		var tag string
		err = rows.Scan(&tag)
		if err != nil {
			return ListTags{}, err
		}
		tags = append(tags, tag)
	}

	if err := rows.Err(); err != nil {
		return ListTags{}, err
	}

	return ListTags{
		Tags:       tags,
		Page:       page,
		TotalPages: totalPages,
	}, nil
}
