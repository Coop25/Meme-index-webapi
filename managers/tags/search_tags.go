package tags

import "math"

func (m *tagsManager) SearchFilesByTags(tags []string, page int, limit int) (SearchTags, error) {
	memes, err := m.accessors.Postgres.SearchFilesByTags(tags, page, limit)
	if err != nil {
		return SearchTags{}, err
	}
	return SearchTags{
		Memes:     m.toMemes(memes.Memes),
		Page:      memes.Page,
		PageCount: int(math.Ceil(float64(memes.TotalCount) / float64(limit))),
		InputTags: tags,
	}, nil
}
