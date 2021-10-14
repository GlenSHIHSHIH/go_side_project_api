package db

import (
	"componentmod/internal/utils/log"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var (
	redisHost, redisPort, redisPassword, redisDBNumber string
)

const (
	CACHE_MISS = "cache: key is missing"
)

//rdb 參數設定
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
		Destination: &redisDBNumber,
	},
}

func GetRedisDB() RedisDB {
	return redisDB
}

func GetCacheRDB() CacheRDB {
	return cacheRDB
}

var (
	redisDB  RedisDB
	cacheRDB CacheRDB
)

type RedisDB struct {
	*redis.Client
	Ctx context.Context
}

type CacheRDB struct {
	*cache.Cache
	Ctx context.Context
}

var Rdb *redis.Client

func ReidsInit() {
	rdbNumber, err := strconv.Atoi(redisDBNumber)

	if err != nil {
		log.Fatal(fmt.Sprintf("%+v", errors.WithStack(err)))
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort), //"localhost:6379"
		Password: redisPassword,                              // no password set
		DB:       rdbNumber,                                  // use default DB
	})

	var ctx = context.Background()

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(fmt.Sprintf("%+v", errors.WithStack(err)))
	}

	cacher := cache.New(&cache.Options{
		Redis: rdb,
	})

	redisDB = RedisDB{rdb, ctx}
	cacheRDB = CacheRDB{cacher, ctx}

}

func (c *CacheRDB) SetItemByCache(ctx context.Context, key string, value interface{}, time time.Duration) error {

	return c.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   time,
	})
}
