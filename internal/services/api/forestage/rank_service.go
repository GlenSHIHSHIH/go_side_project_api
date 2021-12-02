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
	CACHE_PRODUCTION_RANK      = "cache_production_rank"
	CACHE_PRODUCTION_RANK_TIME = 10 * time.Minute
)

type RankService struct {
}

func GetRankService() *RankService {
	return &RankService{}
}

func (PR *RankService) GetProductionRank(count int) (interface{}, error) {

	//get Production Rank 先從cache拿 看看有沒有資料
	var productionRankDTO *forestagedto.ProductionRankDTO
	var productionRank []*forestagedto.ProductionRankData
	cacheRDB := db.GetCacheRDB()
	err := cacheRDB.Get(cacheRDB.Ctx, CACHE_PRODUCTION_RANK, &productionRankDTO)

	if err == nil {
		return productionRankDTO, nil
	}

	if err.Error() != db.CACHE_MISS {
		log.Error(fmt.Sprintf("cache %s not save,%+v", CACHE_PRODUCTION_RANK, err))
	}

	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.Production{})
	sql = sql.Limit(count)
	sql = sql.Order("weight desc").Order("amount desc")
	sql.Select("(liked_count+historical_sold) as amount,id,product_id,name,description,options,categories," +
		"weight,liked_count,historical_sold,stock,image,images,url,price,price_min,create_time").Find(&productionRank)

	productionRankDTO = &forestagedto.ProductionRankDTO{
		ProductionList: productionRank,
	}

	err = cacheRDB.SetItemByCache(cacheRDB.Ctx, CACHE_PRODUCTION_RANK, productionRankDTO, CACHE_PRODUCTION_RANK_TIME)

	if err != nil {
		log.Error(fmt.Sprintf("cache %s not save,%+v", CACHE_PRODUCTION_RANK, err))
	}

	return productionRankDTO, nil
}
