package main

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"log"
	"me-little-tools/interial/apps/common/persistence"
	"me-little-tools/interial/apps/manage_music"
)

func main() {
	//测试连接
	g := gin.Default()
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//获取mysql连接
	db, err := persistence.NewDb("mysql", "root:we-tools-mysql@tcp(192.168.124.35:3306)/me_little_tools?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("connect to mysql failed, err:%v", err)
		return
	}
	//初始化雪花算法
	node, err := snowflake.NewNode(1)
	log.Println(node)

	if err != nil {
		log.Fatalf("new snowflake node failed, err:%v", err)
		return
	}

	repo := manage_music.NewRepo(db)
	//获取全部音乐列表
	music, err := repo.GetAllMusic(1, 2)
	if err != nil {
		log.Fatalf("get all music failed, err:%v", err)
		return
	}
	log.Println("================")
	log.Println(music)
	log.Println("=========获取全部结束=======")

	musicByCreatorId, err := repo.GetMusicByCreatorId(1)
	if err != nil {
		log.Fatalf("get music by creatorId failed,err:%v", err)
		return
	}
	log.Println(musicByCreatorId)

	//文件上传

	usecase := manage_music.NewUsecase(repo, node)
	musicApi := manage_music.NewApi(usecase)
	g.POST("/upload", musicApi.UploadMusic)
	g.Run("localhost:9090")
}
