package db

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host, port, dbname, username, password string
)

//db 參數設定
var DBConfig = []cli.Flag{
	&cli.StringFlag{
		Name:        "db-host",
		Usage:       "db host",
		Value:       "127.0.0.1",
		Destination: &host,
	},
	&cli.StringFlag{
		Name:        "db-port",
		Usage:       "page skip count",
		Value:       "3306",
		Destination: &port,
	},
	&cli.StringFlag{
		Name:        "db-name",
		Usage:       "page skip count",
		Value:       "jiyoung_shopee",
		Destination: &dbname,
	},
	&cli.StringFlag{
		Name:        "db-username",
		Usage:       "page skip count",
		Value:       "glen",
		Destination: &username,
	},
	&cli.StringFlag{
		Name:        "db-password",
		Usage:       "page skip count",
		Value:       "1qaz@WSX",
		Destination: &password,
	},
}

func DBInit() error {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	// db.AutoMigrate()
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

}
