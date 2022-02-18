package db

import (
	"componentmod/internal/utils/log"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"componentmod/internal/utils/db/model"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		Logger:                                   logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal(fmt.Sprintf("%+v", errors.WithStack(err)))
	}

	mySqlDB = MySqlDB{db}

	initTableAndProcedure()
}

func initTableAndProcedure() {

	// create table
	mySqlDB.AutoMigrate(&model.Production{})
	mySqlDB.AutoMigrate(&model.ProductionTemp{})
	mySqlDB.AutoMigrate(&model.Carousel{})
	mySqlDB.AutoMigrate(&model.Picture{})
	mySqlDB.AutoMigrate(&model.User{})
	mySqlDB.AutoMigrate(&model.Role{})
	mySqlDB.AutoMigrate(&model.Menu{})

	// create procedure
	mySqlDB.Exec(model.DROP_PROCEDURE_IF_EXISTS)
	mySqlDB.Exec(model.PROCEDURE_GET_PROD_CATEGORIES)

	// create initial
	initialData("carousels", "resources/initialdata/carousels.sql")
	initialData("pictures", "resources/initialdata/pictures.sql")
	initialData("carousel_picture", "resources/initialdata/carousel_picture.sql")
	initialData("users", "resources/initialdata/users.sql")
	initialData("roles", "resources/initialdata/roles.sql")
	initialData("menus", "resources/initialdata/menus.sql")
	initialData("role_menu", "resources/initialdata/role_menu.sql")
	initialData("user_role", "resources/initialdata/user_role.sql")

}

func initialData(tableName, sqlFilePath string) {
	if mySqlDB.Migrator().HasTable(tableName) {
		var value int64
		mySqlDB.Table(tableName).Count(&value)
		if value == 0 {
			mydir, _ := os.Getwd()
			query, err := ioutil.ReadFile(mydir + "/" + sqlFilePath)
			if err != nil {
				log.Fatal(fmt.Sprintf("%+v", errors.WithStack(err)))
			}

			sqlAll := string(query)
			for _, v := range strings.Split(sqlAll, ";") {
				if v != "" {
					mySqlDB.Exec(v)
				}
			}

		}
	}
}
