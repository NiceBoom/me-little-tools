package main

import (
	"log"
	"mysql/interial/apps/common/persistence"
	"mysql/interial/apps/manage_music"
)

func main() {
	db, err := persistence.NewDb("mysql", "root:we-tools-mysql@tcp(192.168.124.35:3306)/me_little_tools?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("connect to mysql failed, err:#{err}")
		return
	}
	repo := manage_music.NewRepo(db)
	music, err := repo.GetAllMusic()
	if err != nil {
		log.Fatalf("get all music failed, err:#{err}")
		return
	}
	log.Println(music)

}
