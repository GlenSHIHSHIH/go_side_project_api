package db

import (
	"componentmod/internal/utils/log"
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var (
	redisHost, redisPort, redisPassword, redisDB string
)

//db 參數設定
var RedisConfig = []cli.Flag{
	&cli.StringFlag{
		Name:        "redis-host",
		Usage:       "redis host",
		Value:       "127.0.0.1",
		Destination: &redisHost,
	},
	&cli.StringFlag{
		Name:        "redis-port",
		Usage:       "redis port",
		Value:       "6379",
		Destination: &redisPort,
	},
	&cli.StringFlag{
		Name:        "redis-password",
		Usage:       "redis password",
		Value:       "1qaz@WSX",
		Destination: &redisPassword,
	},
	&cli.StringFlag{
		Name:        "redis-db",
		Usage:       "redis db",
		Value:       "0",
		Destination: &redisDB,
	},
}

var Rdb *redis.Client

func ReidsInit() {
	redisDb, err := strconv.Atoi(redisDB)

	if err != nil {
		log.Fatal(fmt.Sprintf("%+v", errors.WithStack(err)))
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort), //"localhost:6379"
		Password: redisPassword,                              // no password set
		DB:       redisDb,                                    // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		log.Fatal(fmt.Sprintf("%+v", errors.WithStack(err)))
	}

	Rdb = rdb

}
