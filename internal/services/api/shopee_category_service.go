package api

import (
	"componentmod/internal/utils/db"
	"componentmod/internal/utils/db/model"
	"componentmod/internal/utils/log"
	"fmt"
	"time"
)

const (
	CACHE_CATEGORY      = "cache_category"
	CACHE_CATEGORY_TIME = 60 * time.Minute
)

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
