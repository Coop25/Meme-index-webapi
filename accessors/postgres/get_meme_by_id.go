package postgres

import "github.com/google/uuid"

func (a *postgresAccessor) GetMemeById(id string) (Meme, error) {
	// Retrieve the file content from the database
	query := `SELECT name, url, description, contenttype FROM files WHERE id = $1;`

	// Parse the search UUID
	searchID, err := uuid.Parse(id)
	if err != nil {
		// log.Fatal("Error parsing search UUID:", err)
		return Meme{}, err
	}

	meme := Meme{
		Id: searchID,
	}
	err = a.db.QueryRow(query, searchID).Scan(&meme.Name, &meme.Url, &meme.Description, &meme.ContentType)
	if err != nil {
		return Meme{}, err
	}

	// Step 2: Retrieve associated tags
	// Retrieve associated tags using the new function
	tags, err := a.GetTagsByFileID(id)
	if err != nil {
		return Meme{}, err
	}
	meme.Tags = tags

	return meme, nil
}
