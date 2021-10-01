package shopee

import (
	"componentmod/internal/dto"
	"componentmod/internal/utils/log"
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/xuri/excelize/v2"
)

type ShopeeExcelService struct {
}

func NewShopeeExcelService() *ShopeeExcelService {
	return &ShopeeExcelService{}
}

func (ses *ShopeeExcelService) WriteExcel(filePath, sheetName string, DataModelList []*dto.ShopeeDataDTO, headerList []map[string]string) error {

	//創建檔案、寫入標題
	f := excelize.NewFile()
	f.SetSheetName("Sheet1", sheetName)

	streamWriter, err := f.NewStreamWriter(sheetName)
	if err != nil {
		return err
	}
	defer streamWriter.Flush()

	//寫入標題
	err = ses.writeExcelData(streamWriter, func() ([][]string, int, int) {
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
	err = ses.writeExcelData(streamWriter, func() ([][]string, int, int) {
		//寫入順序
		var content [][]string
		var detail []string

		for _, DataMode := range DataModelList {
			jsData, err := json.Marshal(DataMode)

			if err != nil {
				log.Warn(fmt.Sprintf("轉型錯誤,ProductId:%v , Name:%v", DataMode.ProductId, DataMode.Name))
				continue
			}

			detail = nil
			for _, mapValue := range headerList {
				for k, _ := range mapValue {
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

func (ses *ShopeeExcelService) writeExcelData(streamWriter *excelize.StreamWriter, fun func() ([][]string, int, int)) error {

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
