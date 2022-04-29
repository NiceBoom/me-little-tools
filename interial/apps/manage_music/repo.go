package manage_music

import (
	"github.com/jmoiron/sqlx"
	"mysql/interial/apps/common/persistence"
)

type Repo interface {
	WithUnitOfWork(f persistence.UnitOfWorkFunc) error

	GetAllMusic(page, prePage int) (*[]Music, error)
	GetMusicByCreatorId(creatorId uint64) (*[]Music, error)
	GetMusicByMusicType(musicType MusicType) (*[]Music, error)
	GetMusicByStatus(musicStatus MusicStatus) (*[]Music, error)

	CreateMusic(tx *sqlx.Tx, music *Music) error
}

type RepoImpl struct {
	db  *sqlx.DB
	uow *persistence.UnitOfWork
}

var _ Repo = (*RepoImpl)(nil)

func NewRepo(db *sqlx.DB) Repo {
	uow := persistence.NewUnitOfWork(db)
	return &RepoImpl{
		db:  db,
		uow: uow,
	}
}

func (r *RepoImpl) WithUnitOfWork(f persistence.UnitOfWorkFunc) error {
	return r.uow.Execute(f)
}

func (r *RepoImpl) GetAllMusic(page, prePage int) (*[]Music, error) {
	var music []Music
	err := r.db.Select(&music, "SELECT * FROM music ORDER BY id DESC LIMIT ?,?", (page-1)*prePage, prePage)
	if err != nil {
		return nil, err
	}
	return &music, err
}

func (r *RepoImpl) GetMusicByCreatorId(creatorId uint64) (*[]Music, error) {
	var music []Music
	err := r.db.Select(&music, "SELECT * FROM music WHERE creator_id = ?", creatorId)
	if err != nil {
		return nil, err
	}
	return &music, err
}

func (r RepoImpl) GetMusicByMusicType(musicType MusicType) (*[]Music, error) {
	var music []Music
	err := r.db.Select(&music, "SELECT * FROM music WHERE type = ?", musicType)
	if err != nil {
		return nil, err
	}
	return &music, err
}

func (r *RepoImpl) GetMusicByStatus(status MusicStatus) (*[]Music, error) {
	var music []Music
	err := r.db.Select(&music, "SELECT * FROM music WHERE status = ?", status)
	if err != nil {
		return nil, err
	}
	return &music, err
}

func (r *RepoImpl) CreateMusic(tx *sqlx.Tx, music *Music) error {
	stmt := "INSERT INTO music (id, title, storage_key, type, creator_id, status, creator_at)" +
		"VALUES (:id, :title, :storage_key, :type, :creator_id, :status, :creator_at)"
	_, err := tx.NamedExec(stmt, music)
	return err
}
