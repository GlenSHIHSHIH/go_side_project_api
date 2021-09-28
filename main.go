package main

import (
	"componentmod/internal/cmd"
	"componentmod/internal/utils/log"
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
		cmd.SetShopeeCommand(),    //資料擷取 (從蝦皮)
		cmd.SetShopeeApiCommand(), //api 輸出
		cmd.ImportExcelToDB(),     //資料寫入DB
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	//－－－－－－－－－－主程式資料 測試－－－－－－－－－－

	// shopeeId := 60323372
	// shopeeId := 32286362
	// skipCount := 0

	//－－－－－－－－－－主程式 測試－－－－－－－－－－
	// shopeeService := shopee.NewShopeeService()
	// shopeeModelGroup, err := shopeeService.RunShopeeService(shopeeId, skipCount)
	// if err != nil {
	// 	// 寫入 log 紀錄
	// 	errContent := errors.New(fmt.Sprintf("Shopee 網址錯誤"))
	// 	errorData := errors.WithMessage(err, errContent.Error())
	// 	log.Error(fmt.Sprintf("%+v", errorData))
	// 	return
	// }

	// mydir, _ := os.Getwd()
	// err = excel.WriteExcel(mydir+"/excel", shopeeModelGroup)
	// if err != nil {
	// 	log.Error(fmt.Sprintf("%+v", err))
	// }

	// －－－－－－－－－－error 測試－－－－－－－－－－

	// err := errors.New(fmt.Sprintf("錯誤"))
	// errorData := errors.WithStack(err)
	// log.Error(fmt.Sprintf("%+v", errorData))

	// －－－－－－－－－－excel 測試－－－－－－－－－－

	// var options []models.Options
	// options = append(options, models.Options{Name: "size", Option: []string{"xs", "s", "m", "l", "xl"}})
	// options = append(options, models.Options{Name: "color", Option: []string{"yello", "red", "black"}})
	// shopeeModel := &models.ShopeeDataModel{}
	// shopeeModel.ProductId = 123
	// shopeeModel.Name = "5165"
	// shopeeModel.Description = "asdf"
	// shopeeModel.Options = options
	// shopeeModel.Categories = "asdf"
	// shopeeModel.Image = "www://asdf"
	// shopeeModel.Images = "www://asdf,www://adfa,www://asdfaksdfj,www://asdfasdf"
	// shopeeModel.Url = "www://asdf"

	// mydir, _ := os.Getwd()
	// err = excel.WriteExcel(mydir+"/excel", shopeeModelGroup)
	// if err != nil {
	// 	log.Error(fmt.Sprintf("%+v", err))
	// }

}
