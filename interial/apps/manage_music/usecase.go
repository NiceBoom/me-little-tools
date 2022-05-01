package manage_music

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jmoiron/sqlx"
	"me-little-tools/interial/apps/common/storage"
	"time"
)

type CreateMusicInputDto struct {
	CreatorID      uint64
	Title          string
	FileContent    []byte
	Type           MusicType
	FilenameSuffix string
}

type CreateMusicOutputDto struct {
	ID  uint64
	Url string
}

var (
	QiNiuAccessKey string = "D_rr0aEI9JuHRoY_T4xhdYIHhS8bgLbRqfR-Ofgp"
	QiNiuSecretKey string = "qvptjcVO6GX2zgNhpox_8YyOc0_tVgjVTfNNs9I2"
	QiNiuBucket    string = "nice-boom"
)

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

func (u *UsecaseImpl) CreateMusic(musicInputDto *CreateMusicInputDto) (*CreateMusicOutputDto, error) {
	id := uint64(u.idGenerator.Generate())
	content := musicInputDto.FileContent
	upload := storage.NewUpload(QiNiuAccessKey, QiNiuSecretKey, QiNiuBucket)
	fileName, err := upload.UploadMusicToQiNiu(&content, musicInputDto.Title, musicInputDto.CreatorID)
	if err != nil {
		return nil, err
	}

	music := Music{
		ID:         id,
		Title:      musicInputDto.Title,
		StorageKey: fileName,
		Type:       musicInputDto.Type,
		CreatorID:  musicInputDto.CreatorID,
		Status:     MusicStatusPending,
		CreatedAt:  time.Now(),
		UpdateAt:   time.Now(),
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
