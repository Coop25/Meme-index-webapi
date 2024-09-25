package files

func (m *fileManager) ListAllMemes(page int, limit int) (ListMemes, error) {
	aListMemes, err := m.accessors.Postgres.ListAllMemes(page, limit)
	if err != nil {
		return ListMemes{}, err
	}
	return ListMemes{
		Memes:     m.toMemes(aListMemes.Memes),
		Page:      aListMemes.Page,
		PageCount: aListMemes.TotalCount,
	}, nil
}
