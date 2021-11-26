package config

import "github.com/urfave/cli/v2"

var (
	WebHost, WebPort, ImgUrl string
)

//web 參數設定
var WebConfig = []cli.Flag{
	&cli.StringFlag{
		Name:        "web-host",
		Usage:       "web host",
		Value:       "kumkum.com",
		Destination: &WebHost,
	},
	&cli.StringFlag{
		Name:        "web-port",
		Usage:       "web port",
		Value:       "80",
		Destination: &WebPort,
	},
	&cli.StringFlag{
		Name:        "web-imgUrl",
		Usage:       "web imgUrl",
		Value:       "https://cf.shopee.tw/file/",
		Destination: &ImgUrl,
	},
}
