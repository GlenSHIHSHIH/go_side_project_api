package forestage

import (
	"componentmod/internal/api/config"
	"componentmod/internal/dto/forestagedto"
	"time"
)

const (
	CACHE_FORESTAGE_CONFIG      = "cache_forestage_config"
	CACHE_FORESTAGE_CONFIG_TIME = 10 * time.Minute
)

type BaseForestageService struct {
}

func GetBaseForestageService() *BaseForestageService {
	return &BaseForestageService{}
}

// 基礎參數
func (b *BaseForestageService) GetBaseConfig() (interface{}, error) {

	//未來從資料庫讀取 需增加cache

	baseConfigDTO := &forestagedto.BaseConfigDTO{
		ImgUrl: config.ImgUrl,
	}

	baseForestageConfig := &forestagedto.BaseForestageConfigDTO{
		BaseConfigDTO: baseConfigDTO,
	}

	return baseForestageConfig, nil
}
