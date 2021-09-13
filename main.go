package main

import (
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
	err := excel.WriteExcel(mydir + "/excel")
	if err != nil {
		log.Error(fmt.Sprintf("%+v", err))
	}
}
