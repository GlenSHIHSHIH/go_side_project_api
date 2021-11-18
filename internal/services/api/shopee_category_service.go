package api

import (
	"componentmod/internal/dto"
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

func (s *Shopee) GetCategoryList() (interface{}, error) {

	//get Category 先從cache拿 看看有沒有資料
	var category []string
	cacheRDB := db.GetCacheRDB()
	err := cacheRDB.Get(cacheRDB.Ctx, CACHE_CATEGORY, &category)

	if err == nil {
		categoryDTO := &dto.ShopeeCategoryDTO{
			Category: category,
		}
		return categoryDTO, nil
	}

	if err.Error() != db.CACHE_MISS {
		log.Error(fmt.Sprintf("cache %s not save,%+v", CACHE_CATEGORY, err))
	}

	sqldb := db.GetMySqlDB()
	sqldb.Raw(model.GET_PROD_CATEGORIES).Scan(&category)

	err = cacheRDB.SetItemByCache(cacheRDB.Ctx, CACHE_CATEGORY, category, CACHE_CATEGORY_TIME)

	if err != nil {
		log.Error(fmt.Sprintf("cache %s not save,%+v", CACHE_CATEGORY, err))
	}

	categoryDTO := &dto.ShopeeCategoryDTO{
		Category: category,
	}

	return categoryDTO, nil
}
