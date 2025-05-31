package db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Redis interface {
	DBClient()
}

type DB interface {
	Redis
}

func GetCredentials() *redis.Options {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	port := os.Getenv("DB_PORT")
	db, _ := strconv.Atoi(port)
	return &redis.Options{
		Username: username,
		Addr:     fmt.Sprintf("%s:6379", host),
		Password: pass,
		DB:       db,
	}
}

func DBClient() *redis.Client {
	credentials := GetCredentials()
	return redis.NewClient(credentials)
}
