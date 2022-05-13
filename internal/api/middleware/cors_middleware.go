package middleware

import (
	"componentmod/internal/api/config"
	"time"

	"github.com/gin-contrib/cors"
)

func CorsConfig() cors.Config {

	corsConf := cors.Config{
		MaxAge:                 12 * time.Hour,
		AllowBrowserExtensions: true,
	}

	if !config.IsProduction() {
		// 在開發環境時，允許所有 origins、所有 methods 和多數的 headers
		corsConf.AllowAllOrigins = true
		corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
		corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host"}
		config.WebHost = "http://localhost"
	} else {
		// 在正式環境時則根據設定檔調整
		corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
		corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host"}
		corsConf.AllowOrigins = []string{"http://172.21.0.12"}
	}

	return corsConf
}
