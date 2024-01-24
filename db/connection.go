package db

import (
	"context"
	"log"
	"os"

	// "github.com/joho/godotenv"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func Init() (*redis.Client, error) {
	// Load the environment variables
	REDISHOST := os.Getenv("REDISHOST")
	if REDISHOST == "" {
		log.Fatal("REDISHOST environment variable is not set")
	}

	REDISPASSWORD := os.Getenv("REDISPASSWORD")
	if REDISPASSWORD == "" {
		log.Fatal("REDISPASSWORD environment variable is not set")
	}

	REDISPORT := os.Getenv("REDISPORT")
	if REDISPORT == "" {
		log.Fatal("REDISPORT environment variable is not set")
	}

    rdb = redis.NewClient(&redis.Options{
        Addr:     REDISHOST + ":" + REDISPORT,
        Password: REDISPASSWORD,
        DB:       0,
    })

    _, err := rdb.Ping(rdb.Context()).Result()
    if err != nil {
        return nil, err
    }

    return rdb, nil
}

func GetDB() *redis.Client {
	return rdb
}

// // save kompas news struct to redis
// func SaveKompasNews(url string, news kompas.KompasNewsStruct) error {
// 	return rdb.Set(ctx, url, news, 0).Err()
// }

// // get kompas news struct from redis
// func GetKompasNews(url string) (kompas.KompasNewsStruct, error) {
// 	var news kompas.KompasNewsStruct
// 	err := rdb.Get(ctx, url).Scan(&news)
// 	return news, err
// }
