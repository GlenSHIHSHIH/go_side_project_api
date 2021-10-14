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

func (s *Shopee) Production(shProduction *dto.ShopeeProductionInDTO) (map[string]interface{}, error) {
	// rdb := db.GetRedisDB()
	// rdb.Set()
	//頁數預設 矯正
	page, pageLimit := pageParameter(shProduction.Page, shProduction.PageLimit, 1, 10)

	shProduction.Page = page
	shProduction.PageLimit = pageLimit
	productionList, count, err := s.getProductionData(shProduction)
	if err != nil {
		return nil, err
	}
	resMap := make(map[string]interface{}, 0)
	resMap["productionList"] = productionList
	resMap["count"] = count
	resMap["page"] = page
	resMap["pageLimit"] = pageLimit

	return resMap, nil
}

func (s *Shopee) getProductionData(shProduction *dto.ShopeeProductionInDTO) ([]*dto.ShopeeProductionOutDTO, int64, error) {
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

	var ShopeeProductionOutDTO []*dto.ShopeeProductionOutDTO
	sql.Select("product_id,name,description,options,categories,image,images,url,price,price_min,create_time").Scan(&ShopeeProductionOutDTO)

	return ShopeeProductionOutDTO, count, nil
}

func (s *Shopee) Category() (map[string]interface{}, error) {

	sqldb := db.GetMySqlDB()
	var category []string
	sqldb.Raw(model.GET_PROD_CATEGORIES).Scan(&category)
	fmt.Printf("%+v\n", category)

	resMap := make(map[string]interface{}, 0)
	resMap["category"] = category

	return resMap, nil
}
