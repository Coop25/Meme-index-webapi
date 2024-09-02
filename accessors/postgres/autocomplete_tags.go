package postgres

import "strings"

func (a *postgresAccessor) AutocompleteTags(in string) ([]string, error) {
	// Retrieve the file content from the database
	query := `SELECT tag FROM tags WHERE LOWER(tag) LIKE $1 ORDER BY tag ASC LIMIT 50;`

	rows, err := a.db.Query(query, strings.ToLower(in)+"%")
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	tags := []string{}
	for rows.Next() {
		var tag string
		err = rows.Scan(&tag)
		if err != nil {
			return []string{}, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
