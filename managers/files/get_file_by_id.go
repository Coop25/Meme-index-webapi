package files

import "github.com/Coop25/the-meme-index-api/accessors/postgres"

func (m *fileManager) GetMeme(id string) (Meme, error) {
	meme, err := m.accessors.Postgres.GetMemeById(id)
	if err != nil {
		return Meme{}, err
	}
	return m.toMeme(meme), nil
}

func (m *fileManager) toMemes(in []postgres.Meme) []Meme {
	memes := []Meme{}
	for _, meme := range in {
		memes = append(memes, m.toMeme(meme))
	}
	return memes
}

func (m *fileManager) toMeme(meme postgres.Meme) Meme {
	return Meme{
		ID:          meme.Id,
		Tags:        meme.Tags,
		ContentType: meme.ContentType,
		FileUrl:     "http://" + m.config.MinioEndpoint + "/" + m.config.MinioBucketName + "/" + meme.Name,
		URL:         meme.Url,
		Description: meme.Description,
	}
}
