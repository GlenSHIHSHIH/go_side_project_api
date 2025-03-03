package middleware

import (
	"componentmod/internal/api/config"
	"componentmod/internal/utils/log"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// func GetGinEngine() GinEngine {
// 	return ginEngine
// }

// type GinEngine struct {
// 	*gin.Engine
// }

// var ginEngine GinEngine

func WebApiInit() {

	gin.SetMode(config.WebGin)

	r := gin.New()

	//middleware
	middlewareInit(r)

	//router group
	Router(r)

	// ginEngine = GinEngine{r}

	err := r.Run(fmt.Sprintf(":%s", config.WebPort))
	if err != nil {
		log.Fatal(fmt.Sprintf("%+v", errors.WithStack(err)))
	}

	log.Info(fmt.Sprintf("default setting listen 0.0.0.0:%s is localhost", config.WebPort))
}

type HandlerFunc gin.HandlerFunc

func middlewareInit(r *gin.Engine) {
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.NoRoute(Direction404)

	//Cors 設定
	r.Use(cors.New(CorsConfig()))

}
