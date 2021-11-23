package api

type BaseApiService struct {
}

func GetBaseApiService() *BaseApiService {
	return &BaseApiService{}
}

//分頁預設值
func (b *BaseApiService) PageParameter(page, pageLimit, defaultPage, defaultPageLimit int) (int, int) {
	if defaultPage <= 1 {
		defaultPage = 1
	}

	if defaultPageLimit <= 15 {
		defaultPageLimit = 15
	}

	if defaultPage < page {
		defaultPage = page
	}

	if defaultPageLimit < pageLimit {
		defaultPageLimit = pageLimit
	}
	return defaultPage, defaultPageLimit
}
