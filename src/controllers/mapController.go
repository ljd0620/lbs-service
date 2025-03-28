package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"lbs-service/src/services"
	"lbs-service/src/middlewares"
)

type MapController struct {
	BaiduService   services.BaiduMapService
	GaodeService   services.GaodeMapService
	TencentService services.TencentMapService
	RedisMiddleware middlewares.RedisMiddleware
}

func (mc *MapController) GetMapData(c *gin.Context) {
	mapType := c.Query("type")
	location := c.Query("location")

	var data interface{}
	var err error

	switch mapType {
	case "baidu":
		data, err = mc.BaiduService.GetLocation(location)
	case "gaode":
		data, err = mc.GaodeService.GetLocation(location)
	case "tencent":
		data, err = mc.TencentService.GetLocation(location)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid map type"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mc.RedisMiddleware.CacheResponse(location, data)

	c.JSON(http.StatusOK, data)
}