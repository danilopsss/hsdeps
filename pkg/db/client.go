package db

import (
	"github.com/redis/go-redis/v9"
)

type RedisCreds struct {
	host string
	user string
	pass string
	db   int
}

type Clients interface {
	GetRedisClient(creds RedisCreds) (*redis.Client, error)
}

type Databases struct {
	db Clients
}

func GetRedisClient(r RedisCreds) *redis.Client {
	redisOpt := &redis.Options{
		Addr:     r.host,
		Username: r.user,
		Password: r.pass,
		DB:       r.db,
	}
	return redis.NewClient(redisOpt)
}
