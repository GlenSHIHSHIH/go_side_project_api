package validate

import (
	"componentmod/internal/api/errorcode"
	"componentmod/internal/dto/backstagedto"
	"componentmod/internal/services/api/backstage"
	"componentmod/internal/utils"

	"github.com/gin-gonic/gin"
)

func AuthorityMenuValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo, err := UserInfoValidate(c)

		if err != nil {
			ErrHandler(c, nil, err)
			return
		}

		menuService := backstage.GetMenuService()
		menuList, err := menuService.GetMenuListByUserId(userInfo.Id)

		if err != nil {
			ErrHandler(c, nil, err)
			return
		}

		menuIsInUrl := false
		url := c.Request.URL.Path
		menuDTO := menuList.(*backstagedto.MenuDTO)

		for _, v := range menuDTO.Menu {
			if url == v.Url {
				menuIsInUrl = true
				break
			}
		}

		if !menuIsInUrl {
			err = utils.CreateApiErr(errorcode.AUTHORITY_INSUFFICINET, errorcode.USER_AUTHORITY_INSUFFICINET)
			ErrHandler(c, nil, err)
			return
		}

		if err == nil && menuIsInUrl {
			c.Next()
		} else {
			ErrHandler(c, nil, err)
		}
	}
}
