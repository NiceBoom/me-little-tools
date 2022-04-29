package manage_music

import (
	"database/sql"
	"time"
)

type Music struct {
	ID          uint64       `db:"id"`
	Title       string       `db:"title"`
	StorageKey  string       `db:"storage_key"`
	Type        MusicType    `db:"type"`
	CreatorID   uint64       `db:"creator_id"`
	Status      MusicStatus  `db:"status"`
	LikeCount   uint64       `db:"like_count"`
	UnLikeCount uint64       `db:"unlike_count"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdateAt    time.Time    `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
	Notes       string       `db:"notes"`
}
