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

var (
	// wg        sync.WaitGroup
	shopeeId  string
	skipCount string
)

//cli v2 參數設定
var shopeeConfig = []cli.Flag{
	&cli.StringFlag{
		Name:        "shopee-id",
		Usage:       "shopee id",
		Aliases:     []string{"Id", "I"},
		Value:       "32286362",
		Destination: &shopeeId,
	},
	&cli.StringFlag{
		Name:        "skip-count",
		Usage:       "page skip count",
		Aliases:     []string{"Skip", "S"},
		Value:       "0",
		Destination: &skipCount,
	},
}

func SetShopeeCommand() *cli.Command {
	Command := &cli.Command{
		Name:   "shopee-data",
		Usage:  "get shopee data and setting shopee's id and page skip count, export to excel",
		Flags:  BuildUpFlag(shopeeConfig, excel.ExcelConfig), //參數
		Action: execShopee,                                   //執行logic
		// Action: executFackData,                               //執行logic
	}

	return Command
}

func execShopee(c *cli.Context) error {
	err := executGetShopeeData()

	// err := executFackData()  假資料

	if err != nil {
		return errors.WithMessage(errors.WithStack(err), fmt.Sprintf(" Write Excel Paht:"+excel.FILE_PATH, excel.FileName))
	}

	return nil
}

func executGetShopeeData() error {
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
func executFackData() error {

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

	filePath, err := excel.GetExcelPath()
	if err != nil {
		return errors.WithStack(err)
	}

	shopeeExcel := shopee.NewShopeeExcelService()
	err = shopeeExcel.WriteExcel(filePath, excel.SHEET_NAME_SHOPEE, DataModelList, excel.HeaderList)
	if err != nil {
		return errors.WithMessage(errors.WithStack(err), fmt.Sprintf(" Write Excel Paht:"+excel.FILE_PATH, excel.FileName))
	}
	return nil
}
