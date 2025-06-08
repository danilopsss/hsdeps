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

func GetRedisClient(r RedisCreds) *redis.Client {
	redisOpt := &redis.Options{
		Addr:     r.Host,
		Username: r.User,
		Password: r.Pass,
		DB:       r.Db,
	}
	return redis.NewClient(redisOpt)
}
