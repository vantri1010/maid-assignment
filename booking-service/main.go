package main

import (
	"booking-service/component/appCtx"
	"booking-service/middleware"
	"booking-service/transport"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	env := appCtx.NewEnv()
	mongoClient := appCtx.NewMongoDatabase(env)

	db := mongoClient.Database(env.DBName)
	defer appCtx.CloseMongoDBConnection(mongoClient)

	router := gin.Default()
	router.Use(middleware.Recover())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("", transport.CreateJob(*db))

	router.Run()
}
