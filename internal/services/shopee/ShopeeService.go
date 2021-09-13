package shopee

import (
	"componentmod/internal/models"
	"componentmod/internal/utils/http"
	"componentmod/internal/utils/log"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

type ShopeeService struct {
}

func NewShopeeService() *ShopeeService {
	return &ShopeeService{}
}

var (
	rwLock         sync.RWMutex
	wg             sync.WaitGroup
	ProductIdGroup []string
)

const (
	ImgUrl    = "https://cf.shopee.tw/file/"
	PUrl      = "https://shopee.tw/%s-i.%s.%s"
	PListApi  = "https://shopee.tw/api/v4/search/search_items?by=pop&limit=30&match_id=%s&newest=%s&order=desc&page_type=shop&scenario=PAGE_OTHERS&version=2"
	PApi      = "https://shopee.tw/api/v4/item/get?itemid=%s&shopid=%s"
	PageCount = 30
)

//蝦皮執行 抓取商品資料
func (Shopee *ShopeeService) RunShopeeService(shopId, skipCount int) error {
	utilHttp := http.NewUtilHttp()

	productListUrl := fmt.Sprintf(PListApi, strconv.Itoa(shopId), strconv.Itoa(skipCount))
	productList, err := utilHttp.HttpGet(productListUrl)
	if err != nil {
		return err
	}

	count := Shopee.GetProductCount(productList)
	productId := Shopee.GetProductIdList(productList)
	ProductIdGroup = append(ProductIdGroup, productId[0:]...)

	for i := PageCount; i < count; i = i + PageCount {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			productListUrl := fmt.Sprintf(PListApi, strconv.Itoa(shopId), strconv.Itoa(i+skipCount))
			productList, err := utilHttp.HttpGet(productListUrl)
			if err != nil {
				// 寫入 log 紀錄
				errContent := errors.New(fmt.Sprintf("連線錯誤,url:%v", productListUrl))
				errData := errors.WithMessage(err, errContent.Error())
				log.Error(fmt.Sprintf("%+v", errData))
				return
			}
			productId := Shopee.GetProductIdList(productList)
			setProductIdToGroup(productId)
		}(i)
	}
	wg.Wait()

	// fmt.Println(ProductIdGroup)

	for _, val := range ProductIdGroup {
		wg.Add(1)
		go func(pId string) {
			defer wg.Done()
			productUrl := fmt.Sprintf(PApi, pId, strconv.Itoa(shopId))
			product, err := utilHttp.HttpGet(productUrl)
			if err != nil {
				// 寫入 log 紀錄
				errContent := errors.New(fmt.Sprintf("連線錯誤,url:%v", productUrl))
				errData := errors.WithMessage(err, errContent.Error())
				log.Error(fmt.Sprintf("%+v", errData))
				return
			}
			productData := Shopee.GetProductData(product)
			// file.WriteExcel()
			fmt.Println(productData)
		}(val)

	}
	wg.Wait()
	return nil
}

// //鎖 多執行緒 讀取後刪除當筆
// func getProductIdToGroup() string {
// 	rwLock.Lock()
// 	defer rwLock.Lock()
// 	ProductId := ProductIdGroup[0]
// 	ProductIdGroup = ProductIdGroup[1:]
// 	return ProductId
// }

//鎖 多執行緒 寫入資料
func setProductIdToGroup(productId []string) {
	rwLock.Lock()
	defer rwLock.Unlock()
	ProductIdGroup = append(ProductIdGroup, productId[0:]...)
}

func (Shopee *ShopeeService) GetProductCount(data string) int {
	value := int(gjson.Get(data, "total_count").Int())
	return value
}

func (Shopee *ShopeeService) GetProductIdList(data string) []string {
	itemid := gjson.Get(data, "items").Array()

	var id []string
	for _, val := range itemid {
		id = append(id, gjson.Get(val.String(), "item_basic.itemid").String())
	}

	return id
}

func (Shopee *ShopeeService) GetProductData(data string) *models.ShopeeDataModel {
	itemid := gjson.Get(data, "data.itemid").String()
	shopid := gjson.Get(data, "data.shopid").String()
	name := gjson.Get(data, "data.name").String()
	description := gjson.Get(data, "data.description").String()
	image := ImgUrl + gjson.Get(data, "data.image").String()
	imgData := gjson.Get(data, "data.images").Array()
	images := ""
	for _, val := range imgData {
		if images != "" {
			images += ","
		}
		images += ImgUrl + val.String()
	}

	variations := gjson.Get(data, "data.tier_variations").Array()
	optionData := models.Options{}
	option := []models.Options{}
	for _, val := range variations {
		optionName := gjson.Get(val.String(), "name").String()
		if optionName == "" {
			continue
		}
		//todo option "顏色"、"尺寸"、"樣式"....
		content := (gjson.Get(val.String(), "options")).String()
		content = strings.Trim(content, "[")
		content = strings.TrimRight(content, "]")
		optionData.Name = optionName
		optionData.Option = strings.Split(content, ",")
		option = append(option, optionData)
	}

	series := gjson.Get(data, "data.categories").Array()
	Categories := ""
	for _, val := range series {
		if Categories != "" {
			Categories += ","
		}
		Categories += (gjson.Get(val.String(), "display_name")).String()
	}

	productId, _ := strconv.ParseInt(itemid, 10, 64)

	// if err != nil {
	// 	fmt.Println("err: itemid not a interger")
	// 	panic("err: itemid not a interger")
	// }

	return &models.ShopeeDataModel{
		ProductId:   productId,
		Name:        name,
		Description: description,
		Option:      option,
		Image:       image,
		Images:      images,
		Categories:  Categories,
		Url:         fmt.Sprintf(PUrl, name, shopid, itemid),
	}
}
