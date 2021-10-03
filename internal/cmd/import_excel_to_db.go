package cmd

import (
	"componentmod/internal/services/shopee"
	"componentmod/internal/utils/db"
	"componentmod/internal/utils/excel"

	"github.com/urfave/cli/v2"
)

func ImportExcelToDB() *cli.Command {
	Command := &cli.Command{
		Name:   "import-excel-to-db",
		Usage:  "Import excel to db",
		Flags:  BuildUpFlag(db.DBConfig, excel.ExcelConfig), //參數
		Action: execImport,                                  //執行logic與初始化
	}

	return Command
}

func execImport(c *cli.Context) error {

	//建置
	db.DBInit() // 1.db

	// 2.執行匯入
	shopeeExcelReader := shopee.NewShopeeExcelReaderService()
	err := shopeeExcelReader.ImportExcelShopeeDataToDB(excel.SHEET_NAME_SHOPEE)

	if err != nil {
		return err
	}

	return nil
}
