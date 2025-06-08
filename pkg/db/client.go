package db

import (
	"github.com/redis/go-redis/v9"
)

type RedisCreds struct {
	Host string
	User string
	Pass string
	Db   int
}

type Clients interface {
	GetRedisClient(creds RedisCreds) (*redis.Client, error)
}

type Databases struct {
	Db Clients
}

func GetRedisClient(creds RedisCreds) *redis.Client {
	redisOpt := &redis.Options{
		Addr:     creds.Host,
		Username: creds.User,
		Password: creds.Pass,
		DB:       creds.Db,
	}
	return redis.NewClient(redisOpt)
}
