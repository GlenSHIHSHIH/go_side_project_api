package api

type Shopee struct {
}

func GetShopeeService() *Shopee {
	return &Shopee{}
}

//分頁預設值
func pageParameter(page, pageLimit, defaultPage, defaultPageLimit int) (int, int) {
	if defaultPage <= 1 {
		defaultPage = 1
	}

	if defaultPageLimit <= 10 {
		defaultPageLimit = 10
	}

	if defaultPage < page {
		defaultPage = page
	}

	if defaultPageLimit < pageLimit {
		defaultPageLimit = pageLimit
	}
	return defaultPage, defaultPageLimit
}
