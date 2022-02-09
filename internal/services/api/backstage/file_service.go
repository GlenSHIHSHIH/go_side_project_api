package backstage

import "github.com/gin-gonic/gin"

const (
	FILE_PATH = "./resources/"
)

type FileService struct {
}

func GetFileService() *FileService {
	return &FileService{}
}

func (f *FileService) GetFile(fileName string, c *gin.Context) {
	c.File(FILE_PATH + fileName)
}
