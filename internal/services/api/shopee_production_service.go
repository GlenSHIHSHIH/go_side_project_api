package api

import (
	errorCode "componentmod/internal/api/error_code"
	"componentmod/internal/dto"
	"componentmod/internal/utils"
	"componentmod/internal/utils/db"
	"componentmod/internal/utils/db/model"
	"fmt"
	"strings"
)

var (
	colume = map[string]string{"PName": "name", "PId": "product_id", "PCategory": "categories", "PCreTime": "create_time"}
)

func (s *Shopee) Production(shProduction *dto.ShopeePageDTO) (interface{}, error) {

	//頁數預設 矯正
	page, pageLimit := pageParameter(shProduction.Page, shProduction.PageLimit, 1, 10)

	shProduction.Page = page
	shProduction.PageLimit = pageLimit
	productionList, count, err := s.getProductionData(shProduction)
	if err != nil {
		return nil, err
	}

	shProduction.Count = count

	res := &dto.ShopeeProductionOutDTO{
		ProductionList: productionList,
		PageData:       shProduction,
	}
	return res, nil
}

func (s *Shopee) getProductionData(shProduction *dto.ShopeePageDTO) ([]*dto.ShopeeProductionData, int64, error) {
	sqldb := db.GetMySqlDB()

	sql := sqldb.Model(&model.Production{})

	//搜尋 產品名稱
	if PSearch := shProduction.Search; PSearch != "" {
		sql = sql.Where("name LIKE ?", "%"+PSearch+"%")
	}

	//篩選 產品分類
	if PSearchOption := shProduction.SearchCategory; PSearchOption != "" {
		sql = sql.Where("categories LIKE ?", "%"+PSearchOption+"%")
	}

	//筆數 count
	var count int64 = 0
	sql.Count(&count)

	//分頁
	// page, pageLimit := pageParameter(shProduction.Page, shProduction.PageLimit, 1, 10)
	sql = sql.Limit(shProduction.PageLimit).Offset((shProduction.Page - 1) * shProduction.PageLimit)

	//排序 依照所選欄位
	if shProduction.SortColumn != "" && (strings.EqualFold(shProduction.Sort, "asc") || strings.EqualFold(shProduction.Sort, "desc")) {
		scolumne := colume[shProduction.SortColumn]
		if scolumne == "" {
			return nil, 0, utils.CreateApiErr(errorCode.PARAMETER_ERROR_CODE, errorCode.PARAMETER_ERROR)
		}

		sql = sql.Order(fmt.Sprintf("%v %v", scolumne, shProduction.Sort))
	}

	var ShopeeProductionData []*dto.ShopeeProductionData
	sql.Select("product_id,name,description,options,categories,image,images,url,price,price_min,create_time").Scan(&ShopeeProductionData)

	return ShopeeProductionData, count, nil
}

func (s *Shopee) ProductionById(id string) (interface{}, error) {
	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.Production{})
	sql = sql.Where("id = ?", id)
	var ShopeeProductionData []*dto.ShopeeProductionData
	sql.Select("product_id,name,description,options,categories,image,images,url,price,price_min,create_time").First(&ShopeeProductionData)

	if len(ShopeeProductionData) == 0 {
		return nil, nil
	}

	return ShopeeProductionData, nil
}
