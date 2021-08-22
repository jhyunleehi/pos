package device

import (
	"pos/client/ibofos"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AddRoutes(r *gin.RouterGroup) {
	route := r.Group("/devices")
	{
		route.GET("", GetDevice)
		route.GET("/all/scan", GetDeviceScan)
		route.GET("/:devicename/smart", GetDeviceSmart)
	}
}

func GetDevice(c *gin.Context) {
	param := ibofos.DeviceParam{}
	err := ibofos.SendIbofos(c, "LISTDEVICE", param)
	if err != nil {
		log.Error(err.Error())
	}
}

func GetDeviceScan(c *gin.Context) {
	param := ibofos.DeviceParam{}
	err := ibofos.SendIbofos(c, "SCANDEVICE", param)
	if err != nil {
		log.Error(err.Error())
	}
}

func GetDeviceSmart(c *gin.Context) {
	param := ibofos.DeviceParam{
		Name: c.Param("devicename"),
	}
	err := ibofos.SendIbofos(c, "SMART", param)
	if err != nil {
		log.Error(err.Error())
	}
}
