package storage

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	"github.com/qiniu/go-sdk/v7/storage"
	"mysql/interial/apps/manage_music"
)

type Upload interface {
	UploadMusicToQiNiu(fileByte *[]byte) (string, error)
	//UploadMusicToLocal()()
	//UploadMusicToAli()()
}

type UploadImpl struct {
}

var _ Upload = (*UploadImpl)(nil)

func NewUpload() Upload {
	return &UploadImpl{}
}

func (u *UploadImpl) UploadMusicToQiNiu(fileByte *[]byte) (string, error) {

	fmt.Println("===========start upload music==========")
	putPolicy := storage.PutPolicy{
		Scope: manage_music.QiNiuBucket,
	}
	mac := qbox.NewMac(manage_music.QiNiuAccessKey, manage_music.QiNiuSecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	dataLen := int64(len(*fileByte))
	err := formUploader.Put(context.Background(), &ret, upToken, "", bytes.NewReader(*fileByte), dataLen, nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(ret.Key, ret.Hash)
	fmt.Println("========stop upload music=========")
	return ret.Key, err
}
