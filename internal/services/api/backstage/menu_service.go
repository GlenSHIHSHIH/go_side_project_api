package backstage

import (
	errorcode "componentmod/internal/api/errorcode"
	"componentmod/internal/dto"
	"componentmod/internal/dto/backstagedto"
	"componentmod/internal/services/api"
	"componentmod/internal/utils"
	"componentmod/internal/utils/db"
	"componentmod/internal/utils/db/model"
	"componentmod/internal/utils/log"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	CACHE_MENU      = "cache_menu"
	CACHE_MENU_TIME = 10 * time.Minute
)

type MenuService struct {
}

func GetMenuService() *MenuService {
	return &MenuService{}
}

func (m *MenuService) GetMenuViewList(p *dto.PageForMultSearchDTO) (interface{}, error) {

	//頁數預設 矯正
	baseApiService := api.GetBaseApiService()
	page, pageLimit := baseApiService.PageParameter(p.Page, p.PageLimit, 1, 10)

	p.Page = page
	p.PageLimit = pageLimit
	menuViewDTO, count, err := m.getMenuData(p)
	if err != nil {
		return nil, err
	}

	p.Count = count

	data := &backstagedto.MenuViewListDTO{
		MenuViewList: menuViewDTO,
		PageData:     p,
	}

	return data, nil
}

func (m *MenuService) getMenuData(p *dto.PageForMultSearchDTO) ([]*backstagedto.MenuViewDTO, int64, error) {

	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.Menu{})

	//搜尋條件分類
	if PSearch := p.Search["name"]; PSearch != "" {
		sql = sql.Where("menus.name LIKE ?", "%"+PSearch+"%")
	}

	if PSearch := p.Search["url"]; PSearch != "" {
		sql = sql.Where("menus.url LIKE ?", "%"+PSearch+"%")
	}

	if PSearch := p.Search["key"]; PSearch != "" {
		sql = sql.Where("menus.key LIKE ?", "%"+PSearch+"%")
	}

	if PSearch := p.Search["parent"]; PSearch != "" {
		sql = sql.Where("menus.parent = ?", PSearch)
	}

	if PSearch := p.Search["feature"]; PSearch != "" {
		sql = sql.Where("menus.feature = ?", PSearch)
	}

	//筆數 count
	var count int64 = 0
	sql.Count(&count)

	//分頁 page, pageLimit := pageParameter(p.Page, p.PageLimit, 1, 10)
	sql = sql.Limit(p.PageLimit).Offset((p.Page - 1) * p.PageLimit)

	//排序 依照所選欄位
	if p.SortColumn != "" && (strings.EqualFold(p.Sort, "asc") || strings.EqualFold(p.Sort, "desc")) {
		scolumne := p.SortColumn
		if scolumne == "" {
			return nil, 0, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
		}

		sql = sql.Order(fmt.Sprintf("%v %v", scolumne, p.Sort))
	}

	var menuViewDTO []*backstagedto.MenuViewDTO
	sql.Select("menus.id,menus.name,menus.key,menus.url,menus.weight,menus.status," +
		"(case when menus.feature ='T' then '標題' when  menus.feature ='P' then '頁面' when  menus.feature ='F' then '按鍵功能' END)as feature," +
		"m.name as parent")
	sql.Joins("left join menus as m on m.id=menus.parent")
	sql.Find(&menuViewDTO)

	return menuViewDTO, count, nil
}

func getMenuNameByUserId(id int) string {
	return CACHE_MENU + "_" + strconv.Itoa(id)
}

func (m *MenuService) RemoveMenuCache(id int) error {
	cacheRDB := db.GetCacheRDB()
	cacheName := getMenuNameByUserId(id)
	return cacheRDB.Delete(cacheRDB.Ctx, cacheName)
}

func (m *MenuService) GetMenuListByUserId(id int) []*backstagedto.MenuData {

	//get Carousels 先從cache拿 看看有沒有資料
	var menu []*backstagedto.MenuData
	cacheName := getMenuNameByUserId(id)
	cacheRDB := db.GetCacheRDB()
	err := cacheRDB.Get(cacheRDB.Ctx, cacheName, &menu)

	if err == nil {
		return menu
	}

	if err.Error() != db.CACHE_MISS {
		log.Error(fmt.Sprintf("cache %s not save,%+v", cacheName, err))
	}

	//get menu (many to many)
	sqldb := db.GetMySqlDB()
	sql := sqldb.Table("users")
	sql = sql.Joins("join user_role on users.id=user_role.user_id and user_role.user_id = ?", id)
	sql = sql.Joins("join role_menu on user_role.role_id= role_menu.role_id")
	sql = sql.Joins("join menus on role_menu.menu_id = menus.id")
	sql = sql.Order("parent asc").Order("weight desc")
	sql.Scan(&menu)

	err = cacheRDB.SetItemByCache(cacheRDB.Ctx, cacheName, menu, CACHE_MENU_TIME)

	if err != nil {
		log.Error(fmt.Sprintf("cache %s not save,%+v", cacheName, err))
	}

	return menu
}

func (m *MenuService) GetMenuNestList(id int) (interface{}, error) {
	menuData := m.GetMenuListByUserId(id)

	menuNestData := nestList(menuData, 0)

	menuNestDTO := &backstagedto.MenuDTO{
		Menu: menuNestData,
	}

	return menuNestDTO, nil
}

func nestList(menuData []*backstagedto.MenuData, parent int) []*backstagedto.MenuNestData {
	var menuNestList []*backstagedto.MenuNestData

	for _, v := range menuData {
		if v.Parent == parent {
			data := &backstagedto.MenuNestData{
				Id:      v.Id,
				Name:    v.Name,
				Key:     v.Key,
				Url:     v.Url,
				Feature: v.Feature,
				Parent:  v.Parent,
			}
			data.Child = nestList(menuData, data.Id)
			menuNestList = append(menuNestList, data)
		}
	}

	return menuNestList
}
