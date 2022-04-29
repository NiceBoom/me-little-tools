package manage_music

import (
	"github.com/gin-gonic/gin"
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

}
