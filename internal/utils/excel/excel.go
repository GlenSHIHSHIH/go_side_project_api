package excel

import (
	"componentmod/internal/models"
	"componentmod/internal/utils/log"
	"encoding/json"
	"fmt"
	"os"

	"github.com/golang-module/carbon"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"github.com/xuri/excelize/v2"

	"componentmod/internal/utils/file"
)

var (
	extensionName = ".xls"
	sheetName     = "shopee"
	// header        = []string{"產品ID(蝦皮ID)", "產品名稱", "敘述", "其他(勿調整)", "分類", "主圖片", "圖片(逗號區隔)", "蝦皮連結"}
	headerList = []map[string]string{ //依照寫入順序 對照代號
		1: {"ProductId": "產品ID(蝦皮ID)"},
		2: {"Name": "產品名稱"},
		3: {"Description": "敘述"},
		4: {"Option": "其他選項"},
		5: {"Categories": "分類"},
		6: {"Image": "主圖片"},
		7: {"Images": "圖片(逗號區隔)"},
		8: {"Url": "蝦皮連結"}}
)

func WriteExcel(folderPath string, DataModelList []*models.ShopeeDataModel) error {
	// func WriteExcel(folderPath string) error {
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

	//創建檔案、寫入標題
	if file.FileIsExist(filePath) == false {
		f := excelize.NewFile()
		f.SetSheetName("Sheet1", sheetName)

		err := writeExcelData(filePath, f, func() ([][]string, int, int) {
			//寫入順序
			var content [][]string
			var detail []string

			for _, mapValue := range headerList {
				for _, v := range mapValue {
					detail = append(detail, v)
				}
			}

			content = append(content, detail)
			rowCount := 1
			rowStart := 1
			return content, rowCount + rowStart, rowStart
		})

		if err != nil {
			return err
		}
	}

	//資料excel 寫入
	f, err := excelize.OpenFile(filePath)
	err = writeExcelData(filePath, f, func() ([][]string, int, int) {
		//寫入順序
		var content [][]string
		var detail []string

		for _, DataMode := range DataModelList {
			jsData, err := json.Marshal(DataMode)
			if err != nil {
				log.Warn(fmt.Sprintf("轉型錯誤,ProductId:%v , Name:%v", DataMode.ProductId, DataMode.Name))
				continue
			}
			for _, mapValue := range headerList {
				for _, v := range mapValue {
					// detail = append(detail, v)
					detail = append(detail, gjson.Get(string(jsData), v).String())
				}
			}

			content = append(content, detail)
		}
		rowCount := len(DataModelList)
		rowStart := 2
		return content, rowCount + rowStart, rowStart
	})
	if err != nil {
		return err
	}

	return nil
}

func writeExcelData(filePath string, f *excelize.File, fun func() ([][]string, int, int)) error {
	fmt.Println(f.GetSheetName(0))
	streamWriter, err := f.NewStreamWriter("shopee")
	if err != nil {
		errContent := errors.New(fmt.Sprintf("串流寫入excel錯誤"))
		errData := errors.WithMessage(err, errContent.Error())
		return errData
	}

	defer streamWriter.Flush()

	content, rowCount, rowStart := fun()
	arrcounter := -1
	for row := rowStart; row < rowCount; row++ {
		var rowData []interface{}
		arrcounter++
		for _, s := range content[arrcounter] {
			for _, value := range s {
				rowData = append(rowData, value)
			}
		}

		cell, _ := excelize.CoordinatesToCellName(1, row)
		if err := streamWriter.SetRow(cell, rowData); err != nil {
			errContent := errors.New(fmt.Sprintf("row寫入excel錯誤"))
			errData := errors.WithMessage(err, errContent.Error())
			return errData
		}
	}

	if err := f.SaveAs(filePath); err != nil {
		errContent := errors.New(fmt.Sprintf("excel存檔錯誤"))
		errData := errors.WithMessage(err, errContent.Error())
		return errData
	}
	return nil
}
