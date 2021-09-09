package main

import (
	"componentmod/internal/utils/log"
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
	// shopeeService.RunShopeeService(shopId, skipCount)

	//log 測試
	// errorData := errors.New("１２３４")
	// err := errors.WithMessage(errorData, errorData.Error())
	// log.Fatal("1234")

	for {
		// if x >= 100 {
		// 	break
		// }
		// x++
		// wg.Add(1)
		go func() {
			// defer wg.Done()
			x := aa()
			log.Info("asd")
			if x > 1000 {
				return
			}
		}()
	}

	// file.WriteFile(mydir+"/log", "log.log", "testGlen12345678945654231123", 0)
	//todo 寫入 exce  utils/excel

	// 	cause := errors.New("whoops")
	// err := errors.WithMessage(cause, err.Error())
	// log.Fatal(fmt.Sprintf("%+v", err))
}

var x = 0

func aa() int {
	x = x + 1
	return x
}
