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

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type CarouselService struct {
}

func GetCarouselService() *CarouselService {
	return &CarouselService{}
}

func (r *CarouselService) GetCarouselViewList(p *dto.PageForMultSearchDTO) (interface{}, error) {

	//頁數預設 矯正
	baseApiService := api.GetBaseApiService()
	page, pageLimit := baseApiService.PageParameter(p.Page, p.PageLimit, 1, 15)

	p.Page = page
	p.PageLimit = pageLimit
	roleViewDTO, count, err := r.getCarouselData(p)
	if err != nil {
		return nil, err
	}

	p.Count = count

	data := &backstagedto.RoleListDTO{
		RoleList: roleViewDTO,
		PageData: p,
	}

	return data, nil
}

func (r *CarouselService) getCarouselData(p *dto.PageForMultSearchDTO) ([]*backstagedto.RoleViewData, int64, error) {

	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.Role{})

	//搜尋條件分類
	if PSearch := p.Search["name"]; PSearch != "" {
		sql = sql.Where("roles.name LIKE ?", "%"+PSearch+"%")
	}

	if PSearch := p.Search["key"]; PSearch != "" {
		sql = sql.Where("roles.key LIKE ?", "%"+PSearch+"%")
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

	var roleViewDTO []*backstagedto.RoleViewData

	sql = sql.Select("roles.*,users.name as CreateUser,u.name as UpdateUser")
	sql = sql.Joins("left join users on users.id=roles.create_user_id")
	sql = sql.Joins("left join users as u on u.id=roles.update_user_id")
	sql.Find(&roleViewDTO)

	return roleViewDTO, count, nil
}

func (r *CarouselService) GetCarouselList() (interface{}, error) {

	var roleCreateOrEditDTO []*backstagedto.RoleOptionList
	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.Role{})
	sql = sql.Order("weight desc")
	sql.Find(&roleCreateOrEditDTO)

	roleOptionListDTO := &backstagedto.RoleOptionListDTO{
		RoleList: roleCreateOrEditDTO,
	}

	return roleOptionListDTO, nil
}

func (r *CarouselService) GetCarouselById(id string) (interface{}, error) {

	var roleCreateOrEditDTO *backstagedto.RoleCreateOrEditDTO
	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.Role{})
	sql = sql.Where("id = ?", id)
	sql.Find(&roleCreateOrEditDTO)

	var menu_id []int
	sql = sqldb.Table("role_menu")
	sql = sql.Where("role_id = ?", id)
	sql.Pluck("menu_id", &menu_id)

	if roleCreateOrEditDTO.Id == 0 {
		roleCreateOrEditDTO = nil
	} else {
		roleCreateOrEditDTO.Select = menu_id
	}

	roleIdDTO := &backstagedto.RoleIdDTO{
		RoleById: roleCreateOrEditDTO,
	}

	return roleIdDTO, nil
}

func (r *CarouselService) DeleteCarousel(ids []string) (interface{}, error) {

	// 從菜單刪除
	sqldb := db.GetMySqlDB()
	err := sqldb.Transaction(func(tx *gorm.DB) error {
		// 從菜單刪除
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Where("id in ?", ids).Delete(&model.Role{}).Error; err != nil {
			// return any error will rollback
			return err
		}

		// 從菜單、權限中繼表單 刪除
		if err := tx.Unscoped().Table("role_menu").Where("role_id in ?", ids).Delete(&model.Role{}).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Table("user_role").Where("role_id in ?", ids).Delete(&model.Role{}).Error; err != nil {
			return err
		}

		//移除全部人的菜單cache
		menuService := GetMenuService()
		menuService.RemoveCacheMenuNameByAllUser()

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.SQL_DELETE_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.SERVER_ERROR_CODE, errorcode.SQL_DELETE_ERROR)
	}

	return nil, nil
}

func (r *CarouselService) CreateCarousel(userInfo *backstagedto.JwtUserInfoDTO, roleCreateOrEditDTO *backstagedto.RoleCreateOrEditDTO) (interface{}, error) {

	role := model.Role{
		Name:         roleCreateOrEditDTO.Name,
		Key:          roleCreateOrEditDTO.Key,
		Weight:       roleCreateOrEditDTO.Weight,
		Status:       roleCreateOrEditDTO.Status,
		Remark:       roleCreateOrEditDTO.Remark,
		CreateTime:   time.Now(),
		CreateUserId: userInfo.Id,
	}

	sqldb := db.GetMySqlDB()
	err := sqldb.Transaction(func(tx *gorm.DB) error {

		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Create(&role).Error; err != nil {
			// return any error will rollback
			return err
		}

		//儲存 role_menu list
		storeRoleMenuTable(role.Id, roleCreateOrEditDTO.Select)
		//移除全部人的菜單cache
		menuService := GetMenuService()
		menuService.RemoveCacheMenuNameByAllUser()
		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {

		errData := errors.WithMessage(errors.WithStack(err), errorcode.SQL_INSERT_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))

		if strings.Contains(err.Error(), "Duplicate") {
			return nil, utils.CreateApiErr(errorcode.SERVER_ERROR_CODE, errorcode.SQL_INSERT_ERROR+": 識別碼(key) 重複,請重新輸入")
		}

		return nil, utils.CreateApiErr(errorcode.SERVER_ERROR_CODE, errorcode.SQL_INSERT_ERROR)
	}

	return nil, nil
}

func (r *CarouselService) EditCarousel(userInfo *backstagedto.JwtUserInfoDTO, id string, roleCreateOrEditDTO *backstagedto.RoleCreateOrEditDTO) (interface{}, error) {

	var role *model.Role
	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.Role{})
	sql.Where("id = ?", id).Find(&role)

	role.Name = roleCreateOrEditDTO.Name
	role.Key = roleCreateOrEditDTO.Key
	role.Weight = roleCreateOrEditDTO.Weight
	role.Status = roleCreateOrEditDTO.Status
	role.Remark = roleCreateOrEditDTO.Remark
	role.UpdateTime = time.Now()
	role.UpdateUserId = userInfo.Id

	err := sqldb.Transaction(func(tx *gorm.DB) error {

		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Save(role).Error; err != nil {
			// return any error will rollback
			return err
		}

		//儲存 role_menu list
		roleId, _ := strconv.Atoi(id)
		storeRoleMenuTable(roleId, roleCreateOrEditDTO.Select)
		//移除全部人的菜單cache
		menuService := GetMenuService()
		menuService.RemoveCacheMenuNameByAllUser()
		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {

		errData := errors.WithMessage(errors.WithStack(err), errorcode.SQL_UPDATE_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))

		if strings.Contains(err.Error(), "Duplicate") {
			return nil, utils.CreateApiErr(errorcode.SERVER_ERROR_CODE, errorcode.SQL_UPDATE_ERROR+": 識別碼(key) 重複,請重新輸入")
		}

		return nil, utils.CreateApiErr(errorcode.SERVER_ERROR_CODE, errorcode.SQL_UPDATE_ERROR)
	}

	return nil, nil
}

func storeCarouselMenuTable(id int, selected []int) {

	menuService := GetMenuService()
	menu := menuService.GetMenuAll()

	sqldb := db.GetMySqlDB()
	sqldb.Unscoped().Table("role_menu").Where("role_id = ?", id).Delete(&model.Role{})

	if len(selected) == 0 {
		return
	}

	var nodes, addParentNode []int
	nodes = append(nodes, selected...)
	addParentNode = append(addParentNode, selected...)

	//搜尋父節點
	for {
		if len(nodes) == 0 {
			break
		}

		fmt.Println("default nodes:")
		fmt.Println(nodes)
		for i := len(nodes) - 1; i >= 0; i-- {
			for _, v := range menu {
				if nodes[i] == v.Id {

					if v.Parent == 0 {
						break
					}

					if !utils.ValueIsInIntArray(addParentNode, v.Parent) {
						addParentNode = append(addParentNode, v.Parent)
					}

					if !utils.ValueIsInIntArray(nodes, v.Parent) {
						nodes = append(nodes, v.Parent)
					}

					break
				}
			}
			nodes = append(nodes[0:i], nodes[i+1:]...)
		}

	}

	var roleMenuArr []map[string]interface{}
	for _, v := range addParentNode {
		roleMenu := map[string]interface{}{"role_id": id, "menu_id": v}
		roleMenuArr = append(roleMenuArr, roleMenu)
	}

	sql := sqldb.Table("role_menu")
	sql.Create(roleMenuArr)

}
