package main

import (
	"componentmod/internal/models"
	"componentmod/internal/utils/excel"
	"componentmod/internal/utils/log"
	"fmt"
	"os"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {

	//加入 cli  套件取值

	// shopId := 60323372
	// shopId := 32286362
	// skipCount := 0

	// shopeeService := shopee.NewShopeeService()
	// err:=shopeeService.RunShopeeService(shopId, skipCount)
	// if err != nil {
	// 	// 寫入 log 紀錄
	// 	errContent := errors.New(fmt.Sprintf("Shopee 網址錯誤"))
	// 	errorData:=errors.WithMessage(err, errContent.Error())
	// log.Error(fmt.Sprintf("%+v", errorData))
	// }

	// err := errors.New(fmt.Sprintf("錯誤"))
	// errContent := errors.New(fmt.Sprintf("Shopee 網址錯誤"))
	// errorData := errors.WithMessage(err, errContent.Error())
	// log.Error(fmt.Sprintf("%+v", errorData))

	//todo 寫入 exce  utils/excel
	mydir, _ := os.Getwd()
	var options []models.Options
	options = append(options, models.Options{Name: "size", Option: []string{"xs", "s", "m", "l", "xl"}})
	options = append(options, models.Options{Name: "color", Option: []string{"yello", "red", "black"}})
	shopeeModel := &models.ShopeeDataModel{}
	shopeeModel.ProductId = 123
	shopeeModel.Name = "5165"
	shopeeModel.Description = "asdf"
	shopeeModel.Option = options
	shopeeModel.Categories = "asdf"
	shopeeModel.Image = "www://asdf"
	shopeeModel.Images = "www://asdf,www://adfa,www://asdfaksdfj,www://asdfasdf"
	shopeeModel.Url = "www://asdf"
	err := excel.WriteExcel(mydir+"/excel", []*models.ShopeeDataModel{shopeeModel})
	if err != nil {
		log.Error(fmt.Sprintf("%+v", err))
	}

}
