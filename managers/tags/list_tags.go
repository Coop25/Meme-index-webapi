package tags

import "math"

func (m *tagsManager) ListAllTags(page int, limit int) (ListTags, error) {
	tags, err := m.accessors.Postgres.ListAllTags(page, limit)
	if err != nil {
		return ListTags{}, err
	}
	return ListTags{
		Tags:      tags.Tags,
		Page:      tags.Page,
		PageCount: int(math.Ceil(float64(tags.TotalPages) / float64(limit))),
	}, nil
}
