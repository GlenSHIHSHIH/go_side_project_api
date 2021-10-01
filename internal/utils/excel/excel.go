package excel

import (
	"componentmod/internal/utils/file"
	"fmt"
	"os"

	"github.com/golang-module/carbon"
	"github.com/urfave/cli/v2"
	"github.com/xuri/excelize/v2"
)

const (
	FOLDER_PATH       = "./excel"
	FILE_PATH         = FOLDER_PATH + "/%s.xlsx"
	SHEET_NAME_SHOPEE = "shopee"
)

var (
	FileName string
	// header        = []string{"產品ID(蝦皮ID)", "產品名稱", "敘述", "其他(勿調整)", "分類", "主圖片", "圖片(逗號區隔)", "蝦皮連結"}
	HeaderList = []map[string]string{ //依照寫入"順序" 對照代號,map 本身無順序
		{"ProductId": "產品ID(蝦皮ID)"},
		{"Name": "產品名稱"},
		{"Description": "敘述"},
		{"Options": "其他選項"},
		{"Categories": "分類"},
		{"Image": "主圖片"},
		{"Images": "圖片(逗號區隔)"},
		{"Url": "蝦皮連結"}}
)

//file name 設定
var ExcelConfig = []cli.Flag{
	&cli.StringFlag{
		Name:        "excel-file-name",
		Usage:       "excel file name",
		Value:       "",
		Destination: &FileName,
	},
}

func filePath() string {
	if len(FileName) == 0 {
		dateTime := carbon.Now(carbon.Taipei).Format("Y-m-d")
		return fmt.Sprintf(FILE_PATH, dateTime)
	} else {
		return fmt.Sprintf(FILE_PATH, FileName)
	}
}

func GetExcelDataBySheet(sheetName string) ([][]string, error) {
	filePath := filePath()

	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	// 獲取 shopee 上所有儲存格
	rows, err := file.GetRows(sheetName)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func GetExcelPath() (string, error) {

	filePath := filePath()
	//建立
	if file.FileIsExist(FOLDER_PATH) == false {
		err := os.MkdirAll(FOLDER_PATH, 0755)
		if err != nil {
			return "", err
		}
	}

	//移除檔案
	if file.FileIsExist(filePath) == true {
		err := os.Remove(filePath)
		if err != nil {
			return "", err
		}
	}

	return filePath, nil
}
