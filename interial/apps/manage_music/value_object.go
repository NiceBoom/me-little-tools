package manage_music

type MusicType string
type MusicStatus string
type StorageMethod string

const (
	MusicStatusPending  MusicStatus = "pending"
	MusicStatusApproved MusicStatus = "approved"
	MusicStatusRejected MusicStatus = "rejected"
	MusicStatusDeleted  MusicStatus = "deleted"

	MusicTypeMp3  MusicType = "mp3"
	MusicTypeFlac MusicType = "flac"
	MusicTypeApe  MusicType = "ape"
	MusicTypeWav  MusicType = "wav"

	StorageMethodQiNiu StorageMethod = "qiniu"
	StorageMethodLocal StorageMethod = "local"
	StorageMethodAli   StorageMethod = "Ali"
)
