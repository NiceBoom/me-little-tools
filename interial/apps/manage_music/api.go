package manage_music

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type Api struct {
	usecase Usecase
}

func NewApi(usecase Usecase) *Api {
	return &Api{usecase}
}

func (api *Api) UploadMusic(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "file upload fail")
		return
	}
	allowedSuffixes := map[string]bool{
		"mp3":  true,
		"flac": true,
		"wav":  true,
		"ape":  true,
	}
	filename := file.Filename
	filenameSplit := strings.Split(filename, ".")
	filenameSplitLower := strings.ToLower(filenameSplit[len(filenameSplit)-1])
	if !allowedSuffixes[filenameSplitLower] {
		c.String(http.StatusBadRequest, "file suffix is not allowed")
		return
	}
	f, err := file.Open()
	if err != nil {
		c.String(http.StatusBadRequest, "file open error"+err.Error())
		return
	}
	fileContent := make([]byte, file.Size)
	_, err = f.Read(fileContent)
	if err != nil {
		c.String(http.StatusBadRequest, "file Serialization err"+err.Error())
	}
	//TODO get creatorId
	var creatorId uint64 = 183522510
	inputDto := &CreateMusicInputDto{
		CreatorID:      creatorId,
		Title:          filename,
		FileContent:    fileContent,
		Type:           MusicType(filenameSplitLower),
		FilenameSuffix: filenameSplitLower,
	}
	outputDto, err := api.usecase.CreateMusic(inputDto)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	log.Println(outputDto)

}
