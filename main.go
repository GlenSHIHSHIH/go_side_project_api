package main

import (
	"componentmod/internal/cmd"
	"componentmod/internal/utils/log"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	//－－－－－－－－－－主程式－－－－－－－－－－
	app := cli.NewApp()
	app.Name = "Shopee"
	app.Usage = "Setting basic configuration"
	app.Version = "0.0.1"
	app.Commands = []*cli.Command{
		cmd.SetShopeeCommand(),         //資料擷取到excel (從蝦皮)
		cmd.SetShopeeDataToDBCommand(), //資料擷取到DB (從蝦皮)
		cmd.SetShopeeApiCommand(),      //api 輸出
		cmd.ImportExcelToDB(),          //資料寫入DB
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(fmt.Sprintf("%+v", err))
	}

}
