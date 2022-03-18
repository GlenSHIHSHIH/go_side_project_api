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
	CACHE_CAROUSEL      = "cache_carousel"
	CACHE_CAROUSEL_TIME = 10 * time.Minute
)

type CarouselService struct {
}

func GetCarouselService() *CarouselService {
	return &CarouselService{}
}

func (c *CarouselService) GetCarouselList() (interface{}, error) {

	//get Carousels 先從cache拿 看看有沒有資料
	var carouselData *forestagedto.CarouselData
	var pictureData []*forestagedto.PictureData
	var carouselDTO *forestagedto.CarouselDTO
	cacheRDB := db.GetCacheRDB()
	err := cacheRDB.Get(cacheRDB.Ctx, CACHE_CAROUSEL, &carouselDTO)

	if err == nil {
		return carouselDTO, nil
	}

	if err.Error() != db.CACHE_MISS {
		log.Error(fmt.Sprintf("cache %s not save,%+v", CACHE_CAROUSEL, err))
	}

	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.Carousel{})
	subQuery1 := sql.Where("status = ?", true).Where("start_time <= now()").Where("end_time >= now()").Order("weight desc").Limit(1)
	sqlQuery := sqldb.Table("(?) as ca", subQuery1)
	sqlQuery.Find(&carouselData)

	sqlPic := sqldb.Model(&model.Picture{}).Where("pictures.status = ?", true).Order("pictures.weight desc")
	sqlPic = sqlPic.Joins("join carousel_picture on pictures.id=picture_id")
	sqlPic = sqlPic.Joins("join carousels on carousels.id=carousel_id and carousels.id = ?", carouselData.Id)
	sqlPic.Find(&pictureData)

	carouselDTO = &forestagedto.CarouselDTO{
		Carousel: carouselData,
		Picture:  pictureData,
	}

	err = cacheRDB.SetItemByCache(cacheRDB.Ctx, CACHE_CAROUSEL, carouselDTO, CACHE_CAROUSEL_TIME)

	if err != nil {
		log.Error(fmt.Sprintf("cache %s not save,%+v", CACHE_CAROUSEL, err))
	}

	return carouselDTO, nil
}
