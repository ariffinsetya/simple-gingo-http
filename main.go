package main

import (
	"github.com/ariffinsetya/simple-gingo-http/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	r := gin.Default()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": controller.Pong(),
		})
	})

	r.GET("/super1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": controller.RedisGet(rdb, "superOne"),
		})
	})

	r.GET("/super2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": controller.RedisGet(rdb, "superTwo"),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
