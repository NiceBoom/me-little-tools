package manage_music

import (
	"github.com/jmoiron/sqlx"
	"mysql/interial/apps/common/persistence"
)

type Repo interface {
	GetAllMusic() (*[]Music, error)
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

func (r *RepoImpl) GetAllMusic() (*[]Music, error) {
	var music []Music
	err := r.db.Select(&music, "SELECT * FROM music ORDER BY id")
	if err != nil {
		return nil, err
	}
	return &music, err
}
