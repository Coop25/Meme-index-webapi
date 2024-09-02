package files

func (m *fileManager) RandomMeme() (Meme, error) {
	meme, err := m.accessors.Postgres.RandomMeme()
	if err != nil {
		return Meme{}, err
	}
	return m.toMeme(meme), nil
}
