package cmd

import (
	"componentmod/internal/dto"
	"componentmod/internal/services/shopee"
	"componentmod/internal/utils/db"
	"componentmod/internal/utils/excel"
	"fmt"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func SetShopeeDataToDBCommand() *cli.Command {
	Command := &cli.Command{
		Name:   "shopee-data-to-db",
		Usage:  "get shopee data and setting shopee's id and page skip count, save to db",
		Flags:  BuildUpFlag(db.DBConfig, shopeeConfig), //參數
		Action: execShopeeSaveDB,                       //執行logic
		// Action: FackData, //執行logic
	}

	return Command
}

func execShopeeSaveDB(c *cli.Context) error {

	//建置
	db.DBInit() // 1.db

	ShopeeDataDTO, err := GetShopeeData()

	if err != nil {
		return errors.WithMessage(errors.WithStack(err), fmt.Sprintf("Get Shopee Data error"))
	}

	shopeeSaveDBService := shopee.NewShopeeSaveDBService()
	shopeeSaveDBService.ShopeeSaveDBService(ShopeeDataDTO)

	if err != nil {
		return errors.WithMessage(errors.WithStack(err), fmt.Sprintf(" Write Excel Paht:"+excel.FILE_PATH, excel.FileName))
	}

	return nil
}

// 假資料
func FackData(c *cli.Context) error {

	//建置
	db.DBInit() // 1.db

	var DataModelList []*dto.ShopeeDataDTO

	attributeData := "[{\"name\": \"產地\",\"value\": \"韓國\",\"id\": 100037,\"is_timestamp\": false,\"brand_option\": null,\"val_id\": null},{\"name\": \"有機\",\"value\": \"是\",\"id\": 100126,\"is_timestamp\": false,\"brand_option\": null,\"val_id\": null}]"

	data := &dto.ShopeeDataDTO{
		ProductId:      1248984949,
		Name:           "name",
		Description:    "description",
		Options:        []dto.Options{{Name: "size", Option: []string{"x", "m", "l"}}},
		Image:          "image",
		Images:         "images",
		Categories:     "Categories",
		Url:            "github.com.tw",
		Attribute:      attributeData,
		LikedCount:     3,
		HistoricalSold: 10,
		Stock:          20,
	}

	DataModelList = append(DataModelList, data)

	//寫入db 測試
	shopeeSaveDBService := shopee.NewShopeeSaveDBService()
	shopeeSaveDBService.ShopeeSaveDBService(DataModelList)
	return nil
}
