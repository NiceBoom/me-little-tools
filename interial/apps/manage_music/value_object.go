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

	StorageMethodQiNiu StorageMethod = "qiniu"
	StorageMethodLocal StorageMethod = "local"
	StorageMethodAli   StorageMethod = "Ali"

	QiNiuAccessKey string = "D_rr0aEI9JuHRoY_T4xhdYIHhS8bgLbRqfR-Ofgp"
	QiNiuSecretKey string = "qvptjcVO6GX2zgNhpox_8YyOc0_tVgjVTfNNs9I2"
	QiNiuBucket    string = "nice-boom"
)
