package api

import (
	errorCode "componentmod/internal/api/error_code"
	"componentmod/internal/dto"
	"componentmod/internal/utils"
	"componentmod/internal/utils/db"
	"componentmod/internal/utils/db/model"
	"componentmod/internal/utils/log"
	"fmt"
	"strings"
	"time"
)

var (
	colume = map[string]string{"PName": "name", "PId": "product_id", "PCategory": "categories", "PCreTime": "create_time"}
)

const (
	CACHE_CATEGORY      = "cache_category"
	CACHE_CATEGORY_TIME = 60 * time.Minute
)

func (s *Shopee) Production(shProduction *dto.ShopeeProductionInDTO) (map[string]interface{}, error) {

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

	//get Category 先從cache拿 看看有沒有資料
	var category []string
	cacheRDB := db.GetCacheRDB()
	err := cacheRDB.Get(cacheRDB.Ctx, CACHE_CATEGORY, &category)

	if err == nil {
		resMap := make(map[string]interface{}, 0)
		resMap["category"] = category
		return resMap, nil
	}

	if err.Error() != db.CACHE_MISS {
		log.Error(fmt.Sprintf("cache %s not save,%+v", CACHE_CATEGORY, err))
	}

	sqldb := db.GetMySqlDB()
	sqldb.Raw(model.GET_PROD_CATEGORIES).Scan(&category)

	err = cacheRDB.SetItemByCache(cacheRDB.Ctx, CACHE_CATEGORY, category, CACHE_CATEGORY_TIME)
	// err = rdb.Set(rdb.Ctx, CACHE_CATEGORY, category, CACHE_CATEGORY_TIME).Err()

	if err != nil {
		log.Error(fmt.Sprintf("cache %s not save,%+v", CACHE_CATEGORY, err))
	}

	resMap := make(map[string]interface{}, 0)
	resMap["category"] = category

	return resMap, nil
}
