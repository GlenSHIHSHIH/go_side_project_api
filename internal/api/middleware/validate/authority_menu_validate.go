package validate

import (
	"componentmod/internal/api/errorcode"
	"componentmod/internal/services/api/backstage"
	"componentmod/internal/utils"
	"fmt"
	"regexp"
	"strings"

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
		menuData := menuService.GetMenuListByUserId(userInfo.Id)

		menuIsInUrl := false
		url := c.Request.URL.Path

		for _, v := range menuData {

			if v.Url == "" {
				continue
			}

			r, _ := regexp.Compile(fmt.Sprintf("^%s[/]*", v.Url))
			if r.MatchString(url) {
				compareUrl := strings.Replace(url, r.FindString(url), "", 1)
				if !strings.Contains(compareUrl, "/") {
					menuIsInUrl = true
					break
				}

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
