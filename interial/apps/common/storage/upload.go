package storage

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	"github.com/qiniu/go-sdk/v7/storage"
	"strconv"
)

type Upload interface {
	UploadMusicToQiNiu(fileByte *[]byte, fileName string, creatorID uint64) (string, error)
	//UploadMusicToLocal()()
	//UploadMusicToAli()()
}

type UploadImpl struct {
	QiNiuAccessKey string
	QiNiuSecretKey string
	QiNiuBucket    string
}

var _ Upload = (*UploadImpl)(nil)

func NewUpload(qiNiuAccessKey, qiNiuSecretKey, qiNiuBucket string) Upload {
	return &UploadImpl{
		QiNiuAccessKey: qiNiuAccessKey,
		QiNiuSecretKey: qiNiuSecretKey,
		QiNiuBucket:    qiNiuBucket,
	}
}

func (u *UploadImpl) UploadMusicToQiNiu(fileByte *[]byte, fileName string, creatorId uint64) (string, error) {

	fmt.Println("===========start upload music==========")
	putPolicy := storage.PutPolicy{
		Scope: u.QiNiuBucket,
	}
	mac := qbox.NewMac(u.QiNiuAccessKey, u.QiNiuSecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	creatorIdString := strconv.FormatUint(creatorId, 10)

	uploadPath := "music_upload/" + creatorIdString + "/" + fileName

	dataLen := int64(len(*fileByte))
	err := formUploader.Put(context.Background(), &ret, upToken, uploadPath, bytes.NewReader(*fileByte), dataLen, nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(ret.Key, ret.Hash)
	fmt.Println("========stop upload music=========")
	return ret.Key, err
}
