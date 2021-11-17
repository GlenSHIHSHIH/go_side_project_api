package shopee

import (
	"componentmod/internal/utils"
	"componentmod/internal/utils/db"
	"componentmod/internal/utils/db/model"
	"componentmod/internal/utils/excel"
	"strconv"
)

type ShopeeExcelReaderService struct {
}

func NewShopeeExcelReaderService() *ShopeeExcelReaderService {
	return &ShopeeExcelReaderService{}
}

func (sers *ShopeeExcelReaderService) ImportExcelShopeeDataToDB(sheetName string) error {
	rows, err := excel.GetExcelDataBySheet(sheetName)
	if err != nil {
		return err
	}

	content := rows[1:] //內容
	var shopeeModelList []*model.ProductionTemp
	for _, columnValue := range content {

		id, _ := strconv.Atoi(columnValue[0])
		price, _ := strconv.Atoi(columnValue[8])
		priceMin, _ := strconv.Atoi(columnValue[9])
		dataModel := &model.ProductionTemp{
			ProductId:    uint32(id),
			Name:         columnValue[1],
			Description:  columnValue[2],
			Options:      columnValue[3],
			Categories:   columnValue[4],
			Image:        columnValue[5],
			Images:       columnValue[6],
			Url:          columnValue[7],
			Price:        int64(price),
			PriceMin:     int64(priceMin),
			CreateUserId: 0,
			UpdateUserId: 0,
		}
		shopeeModelList = append(shopeeModelList, dataModel)
	}

	myDB := db.GetMySqlDB()

	//新增到 temp table
	myDB.Exec(model.TRUNACTE_PRODUCTION_TEMP)
	myDB.CreateInBatches(shopeeModelList, 1000)

	//insert 與 update
	//update
	myDB.Exec(model.UPDATE_SQL_SAFE_CLOSE)
	myDB.Exec(model.UPDATE_PRODUCTION)
	myDB.Exec(model.UPDATE_SQL_SAFE_OPEN)

	// insert
	myDB.Exec(model.INSERT_PRODUCTION)

	return nil
}

func (sers *ShopeeExcelReaderService) getExcelMap(rows [][]string, headerList *[]map[string]string) {
	var mapSort []string
	receiveHeader := rows[0] //標題

	//整合資料順序(以 headerList 順序為主)
	for _, header := range *headerList {
		for k, v := range header {
			i := utils.GetArrayIndexOf(receiveHeader, v)
			if i >= 0 {
				mapSort[i] = k
			}
		}
	}
}
