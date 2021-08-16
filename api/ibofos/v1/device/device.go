package device

import (
	"net/http"
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
	client, err := ibofos.Setup(param)
	if err != nil {
		log.Errorf("%v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	req, res, err := client.Send("LISTDEVICE")
	if err != nil {
		log.Errorf("%s", err.Error())
	}
	log.Debugf("%+v", req)
	log.Debugf("%+v", res)
	c.JSON(http.StatusOK, res)
}

func GetDeviceScan(c *gin.Context) {
	param := ibofos.DeviceParam{}
	client, err := ibofos.Setup(param)
	if err != nil {
		log.Errorf("%v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	req, res, err := client.Send("SCANDEVICE")
	if err != nil {
		log.Errorf("%s", err.Error())
	}
	log.Debugf("%+v", req)
	log.Debugf("%+v", res)
	c.JSON(http.StatusOK, res)
}

func GetDeviceSmart(c *gin.Context) {
	param := ibofos.DeviceParam{
		Name: c.Param("devicename"),
	}
	client, err := ibofos.Setup(param)
	if err != nil {
		log.Errorf("%v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	req, res, err := client.Send("SMART")
	if err != nil {
		log.Errorf("%s", err.Error())
	}
	log.Debugf("%+v", req)
	log.Debugf("%+v", res)
	c.JSON(http.StatusOK, res)
}
