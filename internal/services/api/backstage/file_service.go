package backstage

import "github.com/gin-gonic/gin"

const (
	FIXED_FILE_PATH = "./resources/file/"
	FILE_PATH       = "/file/"
)

type FileService struct {
}

func GetFileService() *FileService {
	return &FileService{}
}

func (f *FileService) GetFile(fileName string, c *gin.Context) {
	c.File(FIXED_FILE_PATH + fileName)
}
