package backstage

import (
	"componentmod/internal/dto/backstagedto"
	"componentmod/internal/utils/db"
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

func getMenuNameByUserId(id int) string {
	return CACHE_MENU + "_" + strconv.Itoa(id)
}

func (M *MenuService) RemoveMenuCache(id int) error {
	cacheRDB := db.GetCacheRDB()
	cacheName := getMenuNameByUserId(id)
	return cacheRDB.Delete(cacheRDB.Ctx, cacheName)
}

func (M *MenuService) GetMenuListByUserId(id int) []*backstagedto.MenuData {

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

func (M *MenuService) GetMenuNestList(id int) (interface{}, error) {
	menuData := M.GetMenuListByUserId(id)

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
