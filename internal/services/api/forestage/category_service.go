package forestage

import (
	"componentmod/internal/dto/forestagedto"
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

type CategoryService struct {
}

func GetCategoryService() *CategoryService {
	return &CategoryService{}
}

func (c *CategoryService) GetCategoryList() (interface{}, error) {

	//get Category 先從cache拿 看看有沒有資料
	var category []string
	var categoryDTO *forestagedto.CategoryDTO
	cacheRDB := db.GetCacheRDB()
	err := cacheRDB.Get(cacheRDB.Ctx, CACHE_CATEGORY, &categoryDTO)

	if err == nil {
		return categoryDTO, nil
	}

	if err.Error() != db.CACHE_MISS {
		log.Error(fmt.Sprintf("cache %s not save,%+v", CACHE_CATEGORY, err))
	}

	sqldb := db.GetMySqlDB()
	sqldb.Raw(model.GET_PROD_CATEGORIES).Scan(&category)

	categoryDTO = &forestagedto.CategoryDTO{
		Category: category,
	}

	err = cacheRDB.SetItemByCache(cacheRDB.Ctx, CACHE_CATEGORY, categoryDTO, CACHE_CATEGORY_TIME)

	if err != nil {
		log.Error(fmt.Sprintf("cache %s not save,%+v", CACHE_CATEGORY, err))
	}

	return categoryDTO, nil
}
