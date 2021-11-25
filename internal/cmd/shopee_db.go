package cmd

import (
	"componentmod/internal/dto"
	"componentmod/internal/services/shopee"
	"componentmod/internal/utils/excel"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func SetShopeeDataToDBCommand() *cli.Command {
	Command := &cli.Command{
		Name:  "shopee-data-db",
		Usage: "get shopee data and setting shopee's id and page skip count, save to db",
		Flags: BuildUpFlag(shopeeConfig, excel.ExcelConfig), //參數
		// Action: execShopeeSaveDB,                             //執行logic
		Action: FackData, //執行logic
	}

	return Command
}

func execShopeeSaveDB(c *cli.Context) error {
	err := GetShopeeData()

	// err := executFackData()  假資料

	if err != nil {
		return errors.WithMessage(errors.WithStack(err), fmt.Sprintf(" Write Excel Paht:"+excel.FILE_PATH, excel.FileName))
	}

	return nil
}

func GetShopeeData() error {
	id, err := strconv.Atoi(shopeeId)
	if err != nil {
		return errors.WithStack(err)
	}

	skip, err := strconv.Atoi(skipCount)
	if err != nil {
		return errors.WithStack(err)
	}

	shopeeService := shopee.NewShopeeService()
	shopeeModelGroup, err := shopeeService.RunShopeeService(id, skip)
	if err != nil {
		// 寫入 log 紀錄
		return errors.WithMessage(errors.WithStack(err), "Shopee 網址錯誤")
	}

	filePath, err := excel.GetExcelPath()
	if err != nil {
		return errors.WithStack(err)
	}

	shopeeExcel := shopee.NewShopeeExcelService()
	err = shopeeExcel.WriteExcel(filePath, excel.SHEET_NAME_SHOPEE, shopeeModelGroup, excel.HeaderList)
	if err != nil {
		return errors.WithMessage(errors.WithStack(err), fmt.Sprintf(" Write Excel Paht:"+excel.FILE_PATH, excel.FileName))
	}

	return nil
}

// 假資料
func FackData(c *cli.Context) error {

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

	return nil
}
