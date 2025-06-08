package db

import (
	"fmt"
	"log"
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
	db, _ := strconv.Atoi(os.Getenv("DB_NUMBER"))
	log.Printf("Creating connection: %s:6379", host)
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
