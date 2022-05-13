package config

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

var (
	WebHost, WebPort, ImgUrl, WebGin string
)

//web 參數設定
var WebConfig = []cli.Flag{
	&cli.StringFlag{
		Name:        "web-host",
		Usage:       "web host",
		Value:       "https://kumkum.ga/api",
		Destination: &WebHost,
		EnvVars:     []string{"web_host"},
	},
	&cli.StringFlag{
		Name:        "web-port",
		Usage:       "web port",
		Value:       "80",
		Destination: &WebPort,
		EnvVars:     []string{"web_port"},
	},
	&cli.StringFlag{
		Name:        "web-imgUrl",
		Usage:       "web imgUrl",
		Value:       "https://cf.shopee.tw/file/",
		Destination: &ImgUrl,
		EnvVars:     []string{"web_imgUrl"},
	},
	&cli.StringFlag{
		Name:        "web-gin",
		Usage:       "web gin",
		Value:       gin.DebugMode,
		Destination: &WebGin,
		EnvVars:     []string{"web_gin"},
	},
}

func IsProduction() bool {
	return gin.ReleaseMode == WebGin
}
