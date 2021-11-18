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
	CACHE_CAROUSEL      = "cache_carousel"
	CACHE_CAROUSEL_TIME = 10 * time.Minute
)

func (s *Shopee) GetCarouselList() (interface{}, error) {

	//get Carousels 先從cache拿 看看有沒有資料
	var carousel []*dto.ShopeeCarouselData
	cacheRDB := db.GetCacheRDB()
	err := cacheRDB.Get(cacheRDB.Ctx, CACHE_CAROUSEL, &carousel)

	if err == nil {
		carouselDTO := dto.ShopeeCarouselDTO{
			Carousel: carousel,
		}
		return carouselDTO, nil
	}

	if err.Error() != db.CACHE_MISS {
		log.Error(fmt.Sprintf("cache %s not save,%+v", CACHE_CAROUSEL, err))
	}

	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.Carousel{})
	sql = sql.Where("status = ?", true)
	sql = sql.Order("weight desc")
	sql.Select("id,name,image,url,weight").Scan(&carousel)

	err = cacheRDB.SetItemByCache(cacheRDB.Ctx, CACHE_CAROUSEL, carousel, CACHE_CAROUSEL_TIME)

	if err != nil {
		log.Error(fmt.Sprintf("cache %s not save,%+v", CACHE_CAROUSEL, err))
	}

	carouselsDTO := dto.ShopeeCarouselDTO{
		Carousel: carousel,
	}

	return carouselsDTO, nil
}
