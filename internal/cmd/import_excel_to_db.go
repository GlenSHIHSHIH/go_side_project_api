package cmd

import (
	"componentmod/internal/utils/db"

	"github.com/urfave/cli/v2"
)

func ImportExcelToDB() *cli.Command {
	Command := &cli.Command{
		Name:   "import-excel-to-db",
		Usage:  "Import excel to db",
		Flags:  BuildUpFlag(db.DBConfig), //參數
		Action: execImport,               //執行logic與初始化
	}

	return Command
}

func execImport(c *cli.Context) error {

	//建置
	// 1.db
	db.DBInit()
	// 2.gin

	return nil
}
