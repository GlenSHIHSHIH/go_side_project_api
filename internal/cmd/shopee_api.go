package cmd

import (
	"componentmod/internal/api/config"
	"componentmod/internal/api/middleware"
	"componentmod/internal/utils"
	"componentmod/internal/utils/db"

	"github.com/urfave/cli/v2"
)

func SetShopeeApiCommand() *cli.Command {
	Command := &cli.Command{
		Name:   "shopee-api",
		Usage:  "Api data for frontend",
		Flags:  BuildUpFlag(db.DBConfig, db.RedisConfig, config.WebConfig, utils.JwtConfig), //參數
		Action: execShopeeApi,                                                               //執行logic與初始化
	}

	return Command
}

func execShopeeApi(c *cli.Context) error {

	//建置

	// 1.db
	db.DBInit()

	// 2.redis
	db.ReidsInit()

	// 3.gin
	middleware.WebApiInit()
	// middleware.WebRun()

	return nil
}
