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

func (m *MenuService) GetMenuParentList() (interface{}, error) {

	var menuParentListDTO []*backstagedto.MenuParentListDTO
	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.Menu{})
	// sql = sql.Debug()
	sql = sql.Where("feature in ?", []string{"T", "P"})
	sql.Scan(&menuParentListDTO)

	menuParentDTO := &backstagedto.MenuParentDTO{
		MenuParentList: menuParentListDTO,
	}

	return menuParentDTO, nil
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

func (m *MenuService) getMenuData(p *dto.PageForMultSearchDTO) ([]*backstagedto.MenuViewData, int64, error) {

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
	baseApiService := api.GetBaseApiService()
	if p.SortColumn == "" || !baseApiService.Check(p.Sort) {
		return nil, 0, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	if p.SortColumn != "" && baseApiService.Check(p.Sort) {
		sql = sql.Order(fmt.Sprintf("%v %v", p.SortColumn, p.Sort))
	}

	var menuViewDTO []*backstagedto.MenuViewData
	sql.Select("menus.id,menus.name,menus.key,menus.url,menus.weight,menus.status,menus.remark," +
		"(case when menus.feature ='T' then '標題' when  menus.feature ='P' then '頁面' when  menus.feature ='F' then '按鍵功能' END)as feature," +
		"m.name as parent")
	sql.Joins("left join menus as m on m.id=menus.parent")
	sql.Find(&menuViewDTO)

	return menuViewDTO, count, nil
}

func (m *MenuService) GetMenuById(id string) (interface{}, error) {

	var menuViewDTO *backstagedto.MenuViewData
	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.Menu{})
	sql = sql.Where("id = ?", id)
	sql.Find(&menuViewDTO)

	if menuViewDTO.Id == 0 {
		menuViewDTO = nil
	}

	menuIdDTO := &backstagedto.MenuIdDTO{
		MenuById: menuViewDTO,
	}

	return menuIdDTO, nil
}

func (m *MenuService) CreateMenu(userInfo *backstagedto.JwtUserInfoDTO, menuCreateOrEditDTO *backstagedto.MenuCreateOrEditDTO) (interface{}, error) {

	parent, _ := strconv.Atoi(menuCreateOrEditDTO.Parent)

	menu := model.Menu{
		Name:         menuCreateOrEditDTO.Name,
		Key:          menuCreateOrEditDTO.Key,
		Url:          menuCreateOrEditDTO.Url,
		Feature:      menuCreateOrEditDTO.Feature,
		Weight:       menuCreateOrEditDTO.Weight,
		Parent:       parent,
		Status:       menuCreateOrEditDTO.Status,
		Remark:       menuCreateOrEditDTO.Remark,
		CreateTime:   time.Now(),
		CreateUserId: userInfo.Id,
	}
	sqldb := db.GetMySqlDB()
	sqldb.Create(&menu)

	return nil, nil
}
func (m *MenuService) EditMenu(userInfo *backstagedto.JwtUserInfoDTO, id string, menuCreateOrEditDTO *backstagedto.MenuCreateOrEditDTO) (interface{}, error) {

	var menu *model.Menu
	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.Menu{})
	sql.Where("id = ?", id).Find(&menu)

	menu.Name = menuCreateOrEditDTO.Name
	menu.Key = menuCreateOrEditDTO.Key
	menu.Url = menuCreateOrEditDTO.Url
	menu.Feature = menuCreateOrEditDTO.Feature
	menu.Weight = menuCreateOrEditDTO.Weight
	parent, _ := strconv.Atoi(menuCreateOrEditDTO.Parent)
	menu.Parent = parent
	menu.Status = menuCreateOrEditDTO.Status
	menu.Remark = menuCreateOrEditDTO.Remark
	menu.UpdateTime = time.Now()
	menu.UpdateUserId = userInfo.Id

	sqldb.Save(menu)

	//移除全部人的菜單cache
	m.RemoveCacheMenuNameByAllUser()

	return nil, nil
}

func (m *MenuService) DeleteMenu(ids []string) (interface{}, error) {

	// 從菜單刪除
	sqldb := db.GetMySqlDB()
	sqldb.Where("id in ?", ids).Delete(&model.Menu{})

	// 從菜單、權限中繼表單 刪除
	sqldb.Unscoped().Table("role_menu").Where("menu_id in ?", ids).Delete(&model.Menu{})

	//移除全部人的菜單cache
	m.RemoveCacheMenuNameByAllUser()
	return nil, nil
}

//移除全部人的菜單cache
func (m *MenuService) RemoveCacheMenuNameByAllUser() {
	redisRDB := db.GetRedisDB()
	keys := redisRDB.Keys(redisRDB.Ctx, CACHE_MENU+"*").Val()
	cacheNames := append([]interface{}{"unlink"}, utils.ChangeStringToInterfaceArr(keys)...)
	redisRDB.Do(redisRDB.Ctx, cacheNames...)
}

func getCacheMenuNameByUserId(id int) string {
	return CACHE_MENU + "_" + strconv.Itoa(id)
}

func (m *MenuService) RemoveMenuCache(id int) error {
	cacheRDB := db.GetCacheRDB()
	cacheName := getCacheMenuNameByUserId(id)

	return cacheRDB.Delete(cacheRDB.Ctx, cacheName)
}

func (m *MenuService) GetMenuListByUserId(id int) []*backstagedto.MenuData {

	//get Carousels 先從cache拿 看看有沒有資料
	var menu []*backstagedto.MenuData
	cacheName := getCacheMenuNameByUserId(id)
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
	sql = sql.Joins("join roles on roles.id= role_menu.role_id and roles.status = true and roles.deleted is NULL")
	sql = sql.Joins("join menus on role_menu.menu_id = menus.id and menus.status = true and menus.deleted is NULL")
	sql = sql.Where("users.deleted is NULL and users.status = true")
	sql = sql.Select("distinct(menus.id),menus.name,menus.key,menus.url,menus.feature,menus.parent,menus.weight,menus.status ")
	sql = sql.Order("menus.parent asc").Order("menus.weight desc")
	sql.Scan(&menu)

	err = cacheRDB.SetItemByCache(cacheRDB.Ctx, cacheName, menu, CACHE_MENU_TIME)

	if err != nil {
		log.Error(fmt.Sprintf("cache %s not save,%+v", cacheName, err))
	}

	return menu
}

func (m *MenuService) GetMenuAll() []*backstagedto.MenuData {
	var menu []*backstagedto.MenuData
	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.Menu{})
	sql.Find(&menu)
	return menu
}

func (m *MenuService) GetMenuAllList() (interface{}, error) {

	menu := m.GetMenuAll()

	menuNestData := nestList(menu, 0)

	menuNestDTO := &backstagedto.MenuDTO{
		Menu: menuNestData,
	}

	return menuNestDTO, nil
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
				Status:  v.Status,
				Parent:  v.Parent,
			}
			data.Child = nestList(menuData, data.Id)
			menuNestList = append(menuNestList, data)
		}
	}

	return menuNestList
}
