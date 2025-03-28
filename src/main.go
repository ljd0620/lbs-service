package main

import (
	"github.com/gin-gonic/gin"
	"lbs-service/src/middlewares"
	"lbs-service/src/controllers"
)

func main() {
	r := gin.Default()

	redisMiddleware := middlewares.NewRedisMiddleware()
	r.Use(redisMiddleware.Connect())

	mapController := controllers.NewMapController()
	
	r.GET("/map/data", mapController.GetMapData)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}