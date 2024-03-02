package main

import (
	"log"
	"strconv"

	"golang-url-shortener/src/config"
	"golang-url-shortener/src/databases"
	"golang-url-shortener/src/repositories"
	"golang-url-shortener/src/routes"
	"golang-url-shortener/src/services"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func main() {
	config, err := config.New()

	if err != nil {
		log.Fatalf("Failed loading config frome env: %v\n", err)
	}

	r := gin.Default()

	redis := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_ADDRESS,
		Password: config.REDIS_PASSWORD,
		DB:       config.REDIS_DATABASE,
	})

	if err := redis.Ping().Err(); err != nil {
		log.Fatal("Failed to connect redis", err)
	}

	redisDatabase := databases.NewRedisDatabase(redis)
	shortenerRepository := repositories.NewShortenerRepository(redisDatabase)
	shortenerService := services.NewShortenerService(shortenerRepository)

	routes.AddShortenerRoutes(&r.RouterGroup, shortenerService)

	if err := r.Run(":" + strconv.Itoa(config.PORT)); err != nil {
		log.Fatalf("Failed to run")
	}
}
