package forestage

import (
	"componentmod/internal/api/config"
	"componentmod/internal/dto/forestage"
)

type BaseForestageService struct {
}

func GetBaseForestageService() *BaseForestageService {
	return &BaseForestageService{}
}

// 基礎參數
func (b *BaseForestageService) GetBaseConfig() (interface{}, error) {

	baseConfigDTO := &forestage.BaseConfigDTO{
		ImgUrl: config.ImgUrl,
	}

	baseForestageConfig := &forestage.BaseForestageConfigDTO{
		BaseConfigDTO: baseConfigDTO,
	}

	return baseForestageConfig, nil
}
