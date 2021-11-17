package api

import (
	"componentmod/internal/utils/db"
	"componentmod/internal/utils/log"
	"fmt"
	"time"
)

const (
	CACHE_CAROUSELS      = "cache_carousels"
	CACHE_CAROUSELS_TIME = 10 * time.Minute
)

func (s *Shopee) Carousels() (map[string]interface{}, error) {

	//get Carousels 先從cache拿 看看有沒有資料
	//以下為完成
	var category []string
	cacheRDB := db.GetCacheRDB()
	err := cacheRDB.Get(cacheRDB.Ctx, CACHE_CAROUSELS, &category)

	if err == nil {
		resMap := make(map[string]interface{}, 0)
		resMap["category"] = category
		return resMap, nil
	}

	if err.Error() != db.CACHE_MISS {
		log.Error(fmt.Sprintf("cache %s not save,%+v", CACHE_CAROUSELS, err))
	}

	err = cacheRDB.SetItemByCache(cacheRDB.Ctx, CACHE_CAROUSELS, category, CACHE_CAROUSELS_TIME)

	if err != nil {
		log.Error(fmt.Sprintf("cache %s not save,%+v", CACHE_CAROUSELS, err))
	}

	return resMap, nil
}
