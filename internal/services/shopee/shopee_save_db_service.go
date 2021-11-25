package shopee

import (
	"componentmod/internal/dto"
	"componentmod/internal/utils/db/model"
	"fmt"

	mapper "github.com/stroiman/go-automapper"
)

type ShopeeSaveDBService struct {
}

func NewShopeeSaveDBService() *ShopeeSaveDBService {
	return &ShopeeSaveDBService{}
}

func (sDB *ShopeeSaveDBService) ShopeeSaveDBService(ShopeeDataDTO []*dto.ShopeeDataDTO) error {
	var shopeeModelList []*model.ProductionTemp
	// myDB := db.GetMySqlDB()
	mapper.Map(ShopeeDataDTO, shopeeModelList)
	fmt.Printf("%+v", shopeeModelList)
	// err := myDB.Transaction(func(myDB *gorm.DB) error {
	// 	if err := myDB.CreateInBatches(shopeeModelList, 1000).Error; err != nil {
	// 		return err
	// 	}
	// 	return nil
	// })

	return nil
}
