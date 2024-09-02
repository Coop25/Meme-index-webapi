package postgres

func (a *postgresAccessor) DeleteMeme(id string) error {
	_, err := a.db.Exec("DELETE FROM files WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
