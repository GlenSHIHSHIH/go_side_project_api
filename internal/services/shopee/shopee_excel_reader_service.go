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
		dataModel := &model.ProductionTemp{
			ProductId:   uint32(id),
			Name:        columnValue[1],
			Description: columnValue[2],
			Options:     columnValue[3],
			Categories:  columnValue[4],
			Image:       columnValue[5],
			Images:      columnValue[6],
			Url:         columnValue[7],
		}
		shopeeModelList = append(shopeeModelList, dataModel)
	}

	myDB := db.GetMySqlDB()

	//新增到 temp table
	myDB.Exec("truncate table  production_temps")
	myDB.CreateInBatches(shopeeModelList, 1000)

	//insert 與 update
	//update
	myDB.Exec(`SET SQL_SAFE_UPDATES=0`)
	myDB.Exec(`update productions as pd  INNER join production_temps as temp on pd.product_id = temp.product_id 
			set 
				pd.name=temp.name,
				pd.description=temp.description,
				pd.options=temp.options,
				pd.categories=temp.categories,
				pd.image=temp.image,
				pd.images=temp.images,
				pd.url=temp.url,
				pd.update_time=temp.update_time `)
	myDB.Exec(`SET SQL_SAFE_UPDATES=1`)

	// insert
	myDB.Exec(`insert into productions(product_id,name,description,options,categories,image,images,url,create_time,update_time)
		select 
			temp.product_id,temp.name,temp.description,temp.options,temp.categories,
			temp.image,temp.images,temp.url,temp.create_time,temp.update_time
		from production_temps as temp left join productions as pd on pd.product_id = temp.product_id 
		where  pd.id is null`)

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
