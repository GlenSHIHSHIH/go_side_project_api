package backstage

import (
	errorcode "componentmod/internal/api/errorcode"
	"componentmod/internal/services/api/forestage"
	"componentmod/internal/utils"
	"componentmod/internal/utils/db"
	"componentmod/internal/utils/log"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

type CacheService struct {
	carousel       string
	category       string
	productionBank string
	production     string
}

func GetCacheService() *CacheService {
	return &CacheService{
		carousel:       forestage.CACHE_CAROUSEL,
		category:       forestage.CACHE_CATEGORY,
		production:     forestage.CACHE_PRODUCTION,
		productionBank: forestage.CACHE_PRODUCTION_RANK,
	}
}

func (c *CacheService) DeleteCache(cacheName string) (interface{}, error) {

	// reflect.Indirect(reflect.ValueOf(forestage).)
	rmCacheName := reflect.ValueOf(c).Elem().FieldByName(cacheName).String()

	if rmCacheName == "<invalid Value>" {
		errData := errors.New(errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	redisRDB := db.GetRedisDB()
	if "production" == cacheName {
		keys := redisRDB.Keys(redisRDB.Ctx, rmCacheName+"*").Val()
		cacheNames := append([]interface{}{"unlink"}, utils.ChangeStringToInterfaceArr(keys)...)
		redisRDB.Do(redisRDB.Ctx, cacheNames...)
		return nil, nil
	}

	redisRDB.Del(redisRDB.Ctx, rmCacheName)

	return nil, nil
}
