package cmd

import (
	"componentmod/internal/services/shopee"
	"componentmod/internal/utils/excel"
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var (
	// wg        sync.WaitGroup
	shopeeId  string
	skipCount string
)

func SetShopeeCommand() *cli.Command {
	Command := &cli.Command{
		Name:   "Shopee Data",
		Usage:  "Get Shopee Data and setting shopee's Id and page skip count",
		Flags:  BuildUpFlag(shopeeConfig), //參數
		Action: execShopee,                //執行logic
	}

	return Command
}

//cli v2 參數設定
var shopeeConfig = []cli.Flag{
	&cli.StringFlag{
		Name:        "shopeeId",
		Usage:       "shopee id",
		Aliases:     []string{"Id", "I"},
		Value:       "32286362",
		Destination: &shopeeId,
	},
	&cli.StringFlag{
		Name:        "skipCount",
		Usage:       "page skip count",
		Aliases:     []string{"Skip", "S"},
		Value:       "0",
		Destination: &skipCount,
	},
}

func execShopee(c *cli.Context) error {

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

	mydir, _ := os.Getwd()
	err = excel.WriteExcel(mydir+"/excel", shopeeModelGroup)
	if err != nil {
		return errors.WithMessage(errors.WithStack(err), fmt.Sprintf(" Write Excel Paht:%s/excel", mydir))
	}

	return nil
}
