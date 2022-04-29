package manage_music

import (
	"log"
	"mysql/interial/apps/common/persistence"
	"testing"
)

func TestRepo(t *testing.T) {
	//获取mysql连接
	db, err := persistence.NewDb("mysql", "root:we-tools-mysql@tcp(192.168.124.35:3306)/me_little_tools?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("connect to mysql failed, err:%v", err)
		return
	}
	repo := NewRepo(db)
	allMusic, err := repo.GetAllMusic(1, 3)
	if err != nil {
		t.Fatal("get all music failed, err:", err)
	}
	log.Println("=======get all music=======")
	log.Println(allMusic)
	log.Println("=======get all music=======")

	musicById, err := repo.GetMusicByCreatorId(7)
	if err != nil {
		t.Fatal("get music by id failed, err:", err)
	}
	log.Println("======get music by id======")
	log.Println(musicById)
	log.Println("======get music by id======")

	musicType, err := repo.GetMusicByMusicType(MusicTypeFlac)
	if err != nil {
		t.Fatal("get music by music type failed, err:", err)
	}
	log.Println("======get music by music type======")
	log.Println(musicType)
	log.Println("======get music by music type======")

	musicByStatus, err := repo.GetMusicByStatus(MusicStatusRejected)
	if err != nil {
		t.Fatal("get music by music status failed, err:", err)
	}
	log.Println("====get by music status====")
	log.Println(musicByStatus)
	log.Println("====get by music status====")
}
