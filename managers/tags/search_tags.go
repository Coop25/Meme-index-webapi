package tags

import "github.com/Coop25/the-meme-index-api/managers/files"

func (m *tagsManager) SearchFilesByTags(tags []string, page int, limit int) (files.ListMemes, error) {
	memes, err := m.accessors.Postgres.SearchFilesByTags(tags, page, limit)
	if err != nil {
		return files.ListMemes{}, err
	}
	return files.ListMemes{
		Memes:      m.toMemes(memes.Memes),
		Page:       memes.Page,
		TotalPages: memes.TotalPages,
	}, nil
}
