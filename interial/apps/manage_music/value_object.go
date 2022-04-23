package manage_music

type MusicType string
type MusicStatus string

const (
	MusicStatusPending  MusicStatus = "pending"
	MusicStatusApproved MusicStatus = "approved"
	MusicStatusRejected MusicStatus = "rejected"
	MusicStatusDeleted  MusicStatus = "deleted"

	MusicTypeMp3  MusicType = "mp3"
	MusicTypeFlac MusicType = "Flac"
	MusicTypeApe  MusicType = "ape"
)
