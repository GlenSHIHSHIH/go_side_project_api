package shopee

import (
	"componentmod/internal/dto"
	"componentmod/internal/utils/db/model"
	"encoding/json"
)

type ShopeeSaveDBService struct {
}

func NewShopeeSaveDBService() *ShopeeSaveDBService {
	return &ShopeeSaveDBService{}
}

func (sDB *ShopeeSaveDBService) ShopeeSaveDBService(ShopeeDataDTO []*dto.ShopeeDataDTO) error {

	var shopeeModelList []*model.ProductionTemp
	for _, shopeeData := range ShopeeDataDTO {
		option, _ := json.Marshal(shopeeData.Options)
		shopeeMode := &model.ProductionTemp{
			ProductId:      uint32(shopeeData.ProductId),
			Name:           shopeeData.Name,
			Description:    shopeeData.Description,
			Options:        string(option),
			Categories:     shopeeData.Categories,
			Image:          shopeeData.Image,
			Images:         shopeeData.Images,
			Url:            shopeeData.Url,
			Price:          shopeeData.Price,
			PriceMin:       shopeeData.PriceMin,
			Attribute:      shopeeData.Attribute,
			LikedCount:     int(shopeeData.LikedCount),
			HistoricalSold: int(shopeeData.HistoricalSold),
			Stock:          int(shopeeData.Stock),
			Status:         true,
		}

		shopeeModelList = append(shopeeModelList, shopeeMode)
	}

	shopeeExcelReaderService := NewShopeeExcelReaderService()

	return shopeeExcelReaderService.WriteShopeeDataToDB(shopeeModelList)
}
