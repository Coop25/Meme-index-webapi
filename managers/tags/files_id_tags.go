package tags

import (
	"github.com/Coop25/the-meme-index-api/accessors/postgres"
	"github.com/Coop25/the-meme-index-api/managers/files"
)

func (m *tagsManager) AddTagsToFile(id string, tags []string) (files.Meme, error) {
	for _, tag := range tags {
		if err := m.accessors.Postgres.AddTagToFile(id, tag); err != nil {
			return files.Meme{}, err
		}
	}
	meme, err := m.accessors.Postgres.GetMemeById(id)
	if err != nil {
		return files.Meme{}, err
	}
	return m.toMeme(meme), nil
}

func (m *tagsManager) RemoveTagFromFile(id string, tags []string) error {
	for _, tag := range tags {
		if err := m.accessors.Postgres.RemoveTagFromFile(id, tag); err != nil {
			return err
		}
	}
	return nil
}

func (m *tagsManager) toMemes(in []postgres.Meme) []files.Meme {
	memes := []files.Meme{}
	for _, meme := range in {
		memes = append(memes, m.toMeme(meme))
	}
	return memes
}

func (m *tagsManager) toMeme(meme postgres.Meme) files.Meme {
	return files.Meme{
		ID:          meme.Id,
		Tags:        meme.Tags,
		ContentType: meme.ContentType,
		FileUrl:     "http://" + m.config.MinioEndpoint + "/" + m.config.MinioBucketName + "/" + meme.Name,
		URL:         meme.Url,
		Description: meme.Description,
	}
}
