package manage_music

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jmoiron/sqlx"
	"mysql/interial/apps/common/storage"
	"time"
)

type CreateMusicInputDto struct {
	Title          string
	FileContent    []byte
	Type           MusicType
	FilenameSuffix string
}

type CreateMusicOutputDto struct {
	ID  uint64
	Url string
}

type Usecase interface {
	CreateMusic(input *CreateMusicInputDto) (*CreateMusicOutputDto, error)
}

type UsecaseImpl struct {
	repo        Repo
	idGenerator *snowflake.Node
}

var _ Usecase = (*UsecaseImpl)(nil)

func NewUsecase(repo Repo, idGenerator *snowflake.Node) Usecase {
	return &UsecaseImpl{
		repo:        repo,
		idGenerator: idGenerator,
	}
}

func (u *UsecaseImpl) CreateMusic(musicByte *CreateMusicInputDto) (*CreateMusicOutputDto, error) {
	id := uint64(u.idGenerator.Generate())
	content := musicByte.FileContent
	upload := storage.NewUpload()
	fileName, err := upload.UploadMusicToQiNiu(&content)
	if err != nil {
		return nil, err
	}

	music := Music{
		ID:         id,
		Title:      musicByte.Title,
		StorageKey: fileName,
		Type:       musicByte.Type,
		CreatorID:  007,
		Status:     MusicStatusPending,
		CreatedAt:  time.Now(),
	}

	err = u.repo.WithUnitOfWork(func(tx *sqlx.Tx) error {
		err := u.repo.CreateMusic(tx, &music)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &CreateMusicOutputDto{
		ID:  id,
		Url: fileName,
	}, nil
}
