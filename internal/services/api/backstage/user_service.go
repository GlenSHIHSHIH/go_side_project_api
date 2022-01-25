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

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
}

func GetUserService() *UserService {
	return &UserService{}
}

// get user data
func (u *UserService) GetUserByLoginName(loginName string) *model.User {
	var user *model.User
	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.User{})
	sql.Where("login_name = ?", loginName).Where("status = ?", true).Find(&user)

	if user.Id == 0 {
		user = nil
	}

	return user
}

// 產出密碼
func (u *UserService) GenUserPwd(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	if err != nil {
		log.Error(fmt.Sprintf("Gnerate password error: %+v", err))
		return "", err
	}

	encodePW := string(hash)
	return encodePW, nil
}

// 確認密碼
func (u *UserService) CheckUserPwd(pwd, hashPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))

	if err != nil {
		return false
	}

	return true
}

func (u *UserService) GetUserViewList(p *dto.PageForMultSearchDTO) (interface{}, error) {

	//頁數預設 矯正
	baseApiService := api.GetBaseApiService()
	page, pageLimit := baseApiService.PageParameter(p.Page, p.PageLimit, 1, 15)

	p.Page = page
	p.PageLimit = pageLimit
	userViewDTO, count, err := u.getUserData(p)
	if err != nil {
		return nil, err
	}

	p.Count = count

	data := &backstagedto.UserListDTO{
		UserList: userViewDTO,
		PageData: p,
	}

	return data, nil
}

func (u *UserService) getUserData(p *dto.PageForMultSearchDTO) ([]*backstagedto.UserViewData, int64, error) {

	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.User{})

	//搜尋條件分類
	if PSearch := p.Search["name"]; PSearch != "" {
		sql = sql.Where("users.name LIKE ?", "%"+PSearch+"%")
	}

	if PSearch := p.Search["loginName"]; PSearch != "" {
		sql = sql.Where("users.login_name LIKE ?", "%"+PSearch+"%")
	}

	if PSearch := p.Search["email"]; PSearch != "" {
		sql = sql.Where("users.email LIKE ?", "%"+PSearch+"%")
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

	var userViewDTO []*backstagedto.UserViewData

	sql = sql.Select("users.*,us.name as CreateUser,u.name as UpdateUser")
	sql = sql.Joins("left join users as us on us.id=users.create_user_id")
	sql = sql.Joins("left join users as u on u.id=users.update_user_id")
	sql.Find(&userViewDTO)

	return userViewDTO, count, nil
}

func (u *UserService) GetUserById(id string) (interface{}, error) {

	var userCreateOrEditDTO *backstagedto.UserCreateOrEditDTO
	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.User{})
	sql = sql.Where("id = ?", id)
	sql.Find(&userCreateOrEditDTO)

	var role_id []int
	sql = sqldb.Table("user_role")
	sql = sql.Where("user_id = ?", id)
	sql.Pluck("role_id", &role_id)

	if userCreateOrEditDTO.Id == 0 {
		userCreateOrEditDTO = nil
	} else {
		userCreateOrEditDTO.Select = role_id
	}

	userIdDTO := &backstagedto.UserIdDTO{
		UserById: userCreateOrEditDTO,
	}

	return userIdDTO, nil
}

// 新增使用者
func (u *UserService) CreateUser(userInfo *backstagedto.JwtUserInfoDTO, userCreateOrEditDTO *backstagedto.UserCreateOrEditDTO) (interface{}, error) {

	//查詢使用者名稱有無重複
	getUser := u.GetUserByLoginName(userCreateOrEditDTO.LoginName)

	if getUser != nil {
		errData := errors.New(errorcode.PARAMETER_LOGIN_NAME_DUPLICATE)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_LOGIN_NAME_DUPLICATE)
	}

	//塞入密碼
	userPwd := userCreateOrEditDTO.Password
	pwd, err := u.GenUserPwd(userPwd)
	userCreateOrEditDTO.Password = pwd

	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	user := model.User{
		Name:         userCreateOrEditDTO.Name,
		LoginName:    userCreateOrEditDTO.LoginName,
		Status:       userCreateOrEditDTO.Status,
		Remark:       userCreateOrEditDTO.Remark,
		Password:     userCreateOrEditDTO.Password,
		Email:        userCreateOrEditDTO.Email,
		UserType:     userCreateOrEditDTO.UserType,
		CreateTime:   time.Now(),
		CreateUserId: userInfo.Id,
	}

	sqldb := db.GetMySqlDB()
	sqldb.Create(&user)

	//儲存 role_menu list
	storeUserRoleTable(user.Id, userCreateOrEditDTO.Select)

	//移除全部人的菜單cache
	menuService := GetMenuService()
	menuService.RemoveCacheMenuNameByAllUser()

	return nil, nil
}

func storeUserRoleTable(id int, selected []int) {

	sqldb := db.GetMySqlDB()
	sqldb.Unscoped().Table("user_role").Where("user_id = ?", id).Delete(&model.User{})

	var userRoleArr []map[string]interface{}
	for _, v := range selected {
		userRole := map[string]interface{}{"user_id": id, "role_id": v}
		userRoleArr = append(userRoleArr, userRole)
	}

	sql := sqldb.Table("user_role")
	sql = sql.Debug()
	sql.Create(userRoleArr)

}

func (u *UserService) EditUser(userInfo *backstagedto.JwtUserInfoDTO, id string, userCreateOrEditDTO *backstagedto.UserCreateOrEditDTO) (interface{}, error) {

	var user *model.User
	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.User{})
	sql.Where("id = ?", id).Find(&user)

	samePws := u.CheckUserPwd(userCreateOrEditDTO.Password, user.Password)
	if !samePws {
		pwd, err := u.GenUserPwd(userCreateOrEditDTO.Password)
		user.Password = pwd
		user.PwdUpdateTime = time.Now()

		if err != nil {
			errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
			log.Error(fmt.Sprintf("%+v", errData))
			return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
		}

	}

	user.Name = userCreateOrEditDTO.Name
	user.LoginName = userCreateOrEditDTO.LoginName
	user.Status = userCreateOrEditDTO.Status
	user.Remark = userCreateOrEditDTO.Remark
	user.Email = userCreateOrEditDTO.Email
	user.UserType = userCreateOrEditDTO.UserType
	user.UpdateTime = time.Now()
	user.UpdateUserId = userInfo.Id

	sqldb.Save(user)

	//儲存 role_menu list
	userId, _ := strconv.Atoi(id)
	storeUserRoleTable(userId, userCreateOrEditDTO.Select)

	//移除全部人的菜單cache
	menuService := GetMenuService()
	menuService.RemoveCacheMenuNameByAllUser()

	return nil, nil
}

func (u *UserService) DeleteUser(ids []string) (interface{}, error) {

	// 從菜單刪除
	sqldb := db.GetMySqlDB()
	sqldb.Where("id in ?", ids).Delete(&model.User{})

	// 從菜單、權限中繼表單 刪除
	sqldb.Unscoped().Table("user_user").Where("user_id in ?", ids).Delete(&model.User{})

	//移除全部人的菜單cache
	menuService := GetMenuService()
	menuService.RemoveCacheMenuNameByAllUser()

	return nil, nil
}
