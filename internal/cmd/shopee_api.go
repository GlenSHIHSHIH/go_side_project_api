package cmd

import (
	"componentmod/internal/utils/db"

	"github.com/urfave/cli/v2"
)

func SetShopeeApiCommand() *cli.Command {
	Command := &cli.Command{
		Name:   "shopee-api",
		Usage:  "Api data for frontend",
		Flags:  BuildUpFlag(db.DBConfig), //參數
		Action: execShopeeApi,            //執行logic
	}

	return Command
}

func execShopeeApi(c *cli.Context) error {

	//建置
	// 1.db
	// 2.gin

	return nil
}
