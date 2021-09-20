package excel

import (
	"componentmod/internal/models"
	"componentmod/internal/utils/log"
	"encoding/json"
	"fmt"
	"os"

	"github.com/golang-module/carbon"
	"github.com/tidwall/gjson"
	"github.com/xuri/excelize/v2"

	"componentmod/internal/utils/file"
)

var (
	extensionName = ".xls"
	sheetName     = "shopee"
	// header        = []string{"產品ID(蝦皮ID)", "產品名稱", "敘述", "其他(勿調整)", "分類", "主圖片", "圖片(逗號區隔)", "蝦皮連結"}
	headerList = []map[string]string{ //依照寫入順序 對照代號
		{"ProductId": "產品ID(蝦皮ID)"},
		{"Name": "產品名稱"},
		{"Description": "敘述"},
		{"Options": "其他選項"},
		{"Categories": "分類"},
		{"Image": "主圖片"},
		{"Images": "圖片(逗號區隔)"},
		{"Url": "蝦皮連結"}}
)

func WriteExcel(folderPath string, DataModelList []*models.ShopeeDataModel) error {
	// func WriteExcel(folderPath string) error {
	dateTime := carbon.Now(carbon.Taipei).Format("Y-m-d")
	filePath := folderPath + "/" + dateTime + extensionName

	//建立
	if file.FileIsExist(folderPath) == false {
		errDir := os.MkdirAll(folderPath, 0755)
		if errDir != nil {
			return errDir
		}
	}

	//移除檔案
	if file.FileIsExist(filePath) == true {
		err := os.Remove(filePath)
		if err != nil {
			return err
		}
	}

	//創建檔案、寫入標題
	f := excelize.NewFile()
	f.SetSheetName("Sheet1", sheetName)

	streamWriter, err := f.NewStreamWriter(sheetName)
	if err != nil {
		return err
	}
	defer streamWriter.Flush()

	//寫入標題
	err = writeExcelData(filePath, streamWriter, func() ([][]string, int, int) {
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

	//資料excel 寫入
	err = writeExcelData(filePath, streamWriter, func() ([][]string, int, int) {
		//寫入順序
		var content [][]string
		var detail []string

		for _, DataMode := range DataModelList {
			jsData, err := json.Marshal(DataMode)

			// fmt.Printf("json: %#v", string(jsData))

			if err != nil {
				log.Warn(fmt.Sprintf("轉型錯誤,ProductId:%v , Name:%v", DataMode.ProductId, DataMode.Name))
				continue
			}

			detail = nil
			for _, mapValue := range headerList {
				for k, _ := range mapValue {
					// if k == "Options" {
					// 	fmt.Printf("\n %v \n", gjson.Get(string(jsData), k).String())
					// 	fmt.Printf("\n %v \n", gjson.Get(string(jsData), k).String())
					// }
					detail = append(detail, gjson.Get(string(jsData), k).String())
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

	if err := f.SaveAs(filePath); err != nil {
		return err
	}

	return nil
}

func writeExcelData(filePath string, streamWriter *excelize.StreamWriter, fun func() ([][]string, int, int)) error {

	content, rowCount, rowStart := fun()
	var rowData []interface{}
	var arrcounter int = -1

	for row := rowStart; row < rowCount; row++ {
		rowData = nil
		arrcounter++
		for _, v := range content[arrcounter] {
			rowData = append(rowData, v)
		}

		cell, _ := excelize.CoordinatesToCellName(1, row)
		if err := streamWriter.SetRow(cell, rowData); err != nil {
			return err
		}
	}

	return nil
}
