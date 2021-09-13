package excel

import (
	"componentmod/internal/utils/log"
	"fmt"
	"os"

	"github.com/golang-module/carbon"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"

	"componentmod/internal/utils/file"
)

var (
	folderPath    string
	extensionName = ".xls"
	sheetName     = "shopee"
	header        = []string{"產品ID(蝦皮ID)", "產品名稱", "敘述", "其他(勿調整)", "分類", "主圖片", "圖片(逗號區隔)", "蝦皮連結"}
)

// func WriteExcel(ShopeeDataModel shopeeData) {
func WriteExcel(folderPath string) error {
	dateTime := carbon.Now(carbon.Taipei).Format("Y-m-d")
	filePath := folderPath + "/" + dateTime + extensionName
	if file.FileIsExist(folderPath) == false {
		errDir := os.MkdirAll(folderPath, 0755)
		if errDir != nil {
			errContent := errors.New(fmt.Sprintf("建立資料夾錯誤"))
			errData := errors.WithMessage(errContent, errDir.Error())
			log.Error(fmt.Sprintf("%+v", errData))
			return errData
		}
	}

	if file.FileIsExist(filePath) == false {
		//創建檔案、寫入標題
		f := excelize.NewFile()
		f.SetSheetName("Sheet1", sheetName)
		err := newFileAndHeader(filePath, f, header)
		if err != nil {
			return err
		}
	}

	// f, err := excelize.OpenFile(filePath)
	// if err != nil {
	// 	log.Warn(err)
	// 	return
	// }

	// shopeeData
	return nil
}

func newFileAndHeader(filePath string, f *excelize.File, content []string) error {
	streamWriter, err := f.NewStreamWriter(sheetName)
	if err != nil {
		errContent := errors.New(fmt.Sprintf("串流寫入excel錯誤"))
		errData := errors.WithMessage(err, errContent.Error())
		return errData
	}

	rowData := make([]interface{}, len(content))
	for i, s := range content {
		rowData[i] = s
	}

	cell, _ := excelize.CoordinatesToCellName(1, 1)
	if err := streamWriter.SetRow(cell, rowData); err != nil {
		errContent := errors.New(fmt.Sprintf("row寫入excel錯誤"))
		errData := errors.WithMessage(err, errContent.Error())
		return errData
	}

	if err := f.SaveAs(filePath); err != nil {
		errContent := errors.New(fmt.Sprintf("excel存檔錯誤"))
		errData := errors.WithMessage(err, errContent.Error())
		return errData
	}
	return nil
}
