package excel

import (
	"os"
	"path/filepath"
)

var Pwd string
var FilePath string
var FileName = "test.csv"

func init() {
	Pwd, _ = os.Getwd()
	FilePath = filepath.Join(Pwd, FileName)
}
