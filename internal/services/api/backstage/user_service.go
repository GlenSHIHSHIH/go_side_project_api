package backstage

import (
	errorCode "componentmod/internal/api/errorcode"
	"componentmod/internal/utils"
	"componentmod/internal/utils/db"
	"componentmod/internal/utils/db/model"
	"componentmod/internal/utils/log"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
}

func GetUserService() *UserService {
	return &UserService{}
}

// 新增使用者
func (u *UserService) CreateUser(user *model.User) (interface{}, error) {

	//查詢使用者名稱有無重複
	getUser := u.GetUserByLoginName(user.LoginName)

	if getUser != nil {
		errData := errors.New(errorCode.PARAMETER_LOGIN_NAME_DUPLICATE)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorCode.PARAMETER_ERROR_CODE, errorCode.PARAMETER_LOGIN_NAME_DUPLICATE)
	}

	//塞入密碼
	userPwd := user.Password
	pwd, err := u.GenUserPwd(userPwd)
	user.Password = pwd

	if err != nil || len(user.Name) <= 3 || len(user.LoginName) <= 3 || len(userPwd) <= 5 {
		otherErr := ""
		if len(user.Name) <= 3 || len(user.LoginName) <= 3 || len(userPwd) <= 5 {
			otherErr = fmt.Sprintf(", or user.Name,user.LoginName lens <=3 and len(userPwd) <= 5, len(user.Name)=%v,len(user.LoginName)=%v,len(userPwd)=%v", len(user.Name), len(user.LoginName), len(userPwd))
		}

		errData := errors.WithMessage(errors.WithStack(err), errorCode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v"+otherErr, errData))
		return nil, utils.CreateApiErr(errorCode.PARAMETER_ERROR_CODE, errorCode.PARAMETER_ERROR)
	}

	//寫入 db
	//todo 新增人員

	sqldb := db.GetMySqlDB()
	result := sqldb.Select("name", "login_name", "password", "email", "status", "user_type", "create_user_id", "remark").Create(user)
	err = result.Error
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorCode.SERVER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorCode.SERVER_ERROR_CODE, errorCode.SERVER_ERROR)
	}

	return nil, nil
}

// get user data
func (u *UserService) GetUserByLoginName(loginName string) *model.User {
	var user *model.User
	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.User{})
	sql.Where("login_name = ?", loginName).Find(&user)

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
