package main

import (
	"log"

	"golang-url-shortener/src/databases"
	"golang-url-shortener/src/repositories"
	"golang-url-shortener/src/routes"
	"golang-url-shortener/src/services"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func main() {
	r := gin.Default()

	redis := redis.NewClient(&redis.Options{
		Addr:     "locahost:6379",
		Password: "123456",
	})

	if err := redis.Ping().Err(); err != nil {
		log.Fatal("Failed to connect redis", err)
	}

	redisDatabase := databases.NewRedisDatabase(redis)
	shortenerRepository := repositories.NewShortenerRepository(redisDatabase)
	shortenerService := services.NewShortenerService(shortenerRepository)

	routes.AddShortenerRoutes(&r.RouterGroup, shortenerService)

	if err := r.Run("0.0.0.0:80"); err != nil {
		log.Fatalf("Failed to run")
	}
}
