package http

import (
	"componentmod/internal/utils/log"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type UtilHttp struct {
}

func NewUtilHttp() *UtilHttp {
	return &UtilHttp{}
}

const (
	Domain        = "https://shopee.tw"
	TimeoutSecond = 3
	WaitTime      = 100
)

//http get 擷取api
func (uh UtilHttp) HttpGet(url string) string {
	client := &http.Client{
		Timeout: TimeoutSecond * time.Second,
	}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Origin", Domain)

	//短暫等待 不要瘋狂送封包 免得被鎖IP
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(10) + 1
	time.Sleep(time.Duration(random) * WaitTime * time.Millisecond)

	resp, err := client.Do(request)

	if err != nil {
		// 寫入 log 紀錄
		errData := errors.New(fmt.Sprintf("連線錯誤,url:%v"))
		errors := errors.WithMessage(err, errData.Error())
		log.Warn(errors)
		// return "", err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)

}
