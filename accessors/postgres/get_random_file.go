package postgres

func (a *postgresAccessor) RandomMeme() (Meme, error) {
	// Retrieve the file content from the database
	query := `SELECT id, name, url, description, contenttype FROM files ORDER BY random() LIMIT 1;`

	meme := Meme{}
	err := a.db.QueryRow(query).Scan(&meme.Id, &meme.Name, &meme.Url, &meme.Description, &meme.ContentType)
	if err != nil {
		return Meme{}, err
	}

	// Step 2: Retrieve associated tags
	// Retrieve associated tags using the new function
	tags, err := a.GetTagsByFileID(meme.Id.String())
	if err != nil {
		return Meme{}, err
	}
	meme.Tags = tags

	return meme, nil
}
