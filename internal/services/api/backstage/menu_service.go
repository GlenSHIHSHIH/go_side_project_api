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

func (M *MenuService) GetMenuListByUserId(id int) (interface{}, error) {

	//get Carousels 先從cache拿 看看有沒有資料
	var menu []*backstagedto.MenuData
	var menuDTO *backstagedto.MenuDTO
	cacheName := CACHE_MENU + "_" + strconv.Itoa(id)
	cacheRDB := db.GetCacheRDB()
	err := cacheRDB.Get(cacheRDB.Ctx, cacheName, &menuDTO)

	if err == nil {
		return menuDTO, nil
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
	sql.Scan(&menu)

	menuDTO = &backstagedto.MenuDTO{
		Menu: menu,
	}

	err = cacheRDB.SetItemByCache(cacheRDB.Ctx, cacheName, menuDTO, CACHE_MENU_TIME)

	if err != nil {
		log.Error(fmt.Sprintf("cache %s not save,%+v", cacheName, err))
	}

	return menuDTO, nil
}

func (M *MenuService) GetMenuNestList() {

}
