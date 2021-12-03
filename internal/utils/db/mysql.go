package db

import (
	"componentmod/internal/utils/log"
	"fmt"

	"componentmod/internal/utils/db/model"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbHost, dbPort, dbName, dbUserName, dbPassword string
)

//db 參數設定
var DBConfig = []cli.Flag{
	&cli.StringFlag{
		Name:        "db-host",
		Usage:       "db host",
		Value:       "127.0.0.1",
		Destination: &dbHost,
		EnvVars:     []string{"db_host"},
	},
	&cli.StringFlag{
		Name:        "db-port",
		Usage:       "db port",
		Value:       "3306",
		Destination: &dbPort,
		EnvVars:     []string{"db_port"},
	},
	&cli.StringFlag{
		Name:        "db-name",
		Usage:       "db name",
		Value:       "jiyoung_shopee",
		Destination: &dbName,
		EnvVars:     []string{"db_name"},
	},
	&cli.StringFlag{
		Name:        "db-username",
		Usage:       "db username",
		Value:       "glen",
		Destination: &dbUserName,
		EnvVars:     []string{"db_username"},
	},
	&cli.StringFlag{
		Name:        "db-password",
		Usage:       "db password",
		Value:       "1qaz@WSX",
		Destination: &dbPassword,
		EnvVars:     []string{"db_password"},
	},
}

func GetMySqlDB() MySqlDB {
	return mySqlDB
}

var mySqlDB MySqlDB

type MySqlDB struct {
	*gorm.DB
}

func DBInit() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUserName, dbPassword, dbHost, dbPort, dbName)
	// log.Fatal(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal(fmt.Sprintf("%+v", errors.WithStack(err)))
	}

	mySqlDB = MySqlDB{db}

	initTableAndProcedure()
}

func initTableAndProcedure() {

	//create table
	mySqlDB.AutoMigrate(&model.Production{})
	mySqlDB.AutoMigrate(&model.ProductionTemp{})
	mySqlDB.AutoMigrate(&model.Carousel{})
	mySqlDB.AutoMigrate(&model.User{})
	mySqlDB.AutoMigrate(&model.Role{})
	mySqlDB.AutoMigrate(&model.Menu{})
	//create procedure
	mySqlDB.Exec(model.DROP_PROCEDURE_IF_EXISTS)
	mySqlDB.Exec(model.PROCEDURE_GET_PROD_CATEGORIES)
}
