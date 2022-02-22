package forestage

import (
	"componentmod/internal/dto/forestagedto"
	"componentmod/internal/utils/db"
	"componentmod/internal/utils/db/model"
	"componentmod/internal/utils/log"
	"fmt"
	"time"

	"github.com/stroiman/go-automapper"
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
	var carouselPicture []*forestagedto.CarouselPictureData
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
	sqlQuery := sqldb.Table("(?) as ca", subQuery1).Joins("join carousel_picture on carousel_id=ca.id ").Joins("join pictures on picture_id=pictures.id")
	sqlQuery = sqlQuery.Order("pictures.weight desc").Select("ca.id,ca.name as CarouselName,ca.start_time,ca.end_time,pictures.name as PictureName,alt,url,pictures.weight")
	sqlQuery.Scan(&carouselPicture)

	pictureData := []*forestagedto.PictureData{}

	automapper.Map(carouselPicture, &pictureData)

	carouselDTO = &forestagedto.CarouselDTO{
		Carousel: forestagedto.CarouselData{
			Id:           carouselPicture[0].Id,
			CarouselName: carouselPicture[0].CarouselName,
			StartTime:    carouselPicture[0].StartTime,
			EndTime:      carouselPicture[0].EndTime,
		},
		Picture: pictureData,
	}

	err = cacheRDB.SetItemByCache(cacheRDB.Ctx, CACHE_CAROUSEL, carouselDTO, CACHE_CAROUSEL_TIME)

	if err != nil {
		log.Error(fmt.Sprintf("cache %s not save,%+v", CACHE_CAROUSEL, err))
	}

	return carouselDTO, nil
}
