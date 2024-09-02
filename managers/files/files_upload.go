package files

import (
	"github.com/Coop25/the-meme-index-api/accessors/minio"
	"github.com/Coop25/the-meme-index-api/accessors/postgres"
	"github.com/google/uuid"
)

func (m *fileManager) UploadFile(in UploadFileRequest) (string, error) {
	meme, err := m.accessors.Minio.UploadMeme(minio.NewMeme{
		Id:       uuid.New(),
		FileName: in.Filename,
		Content:  in.File,
	})
	if err != nil {
		// If there was an error uploading the meme to the storage, try to clean up anything that was created in the storage and return the error
		_ = m.DeleteFileById(meme.Id.String())
		return "", err
	}

	id, err := m.accessors.Postgres.UploadMeme(postgres.NewMeme{
		Id:          meme.Id,
		Name:        meme.FileName,
		Tags:        in.Tags,
		Url:         in.URL,
		ContentType: meme.ContentType,
		Description: in.Description,
	})
	if err != nil {
		// If the there was an error uploading the meme to the database, delete the meme from the storage and clean up anything that was created in postgres
		_ = m.DeleteFileById(meme.Id.String())
		return "", err
	}

	return id, nil
}
