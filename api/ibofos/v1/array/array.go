package array

import (
	"pos/client/ibofos"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AddRoutes(r *gin.RouterGroup) {
	route := r.Group("/array")
	{
		route.POST("", CreateArray)
		route.GET("/:arrayname", GetArrayInfo)
		route.GET("/:arrayname/devices", GetArrayDevice)
		route.DELETE("/:arrayname", DeleteArray)
		route.POST("/:arrayname/mount", CreateArrayMount)
		route.DELETE("/:arrayname/mount", DeleteArrayMount)
		route.POST("/:arrayname/:devices", ArrayDeviceAdd)
		route.DELETE("/:arrayname/:devices", ArrayDeviceRemove)
	}
}

func CreateArray(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		log.Error(err.Error())
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		return
	}
	param := ibofos.ArrayParam{
		Name:     bodyparam.Param.Name,
		RaidType: bodyparam.Param.RaidType,
		Buffer:   bodyparam.Param.Buffer,
		Data:     bodyparam.Param.Data,
		Spare:    bodyparam.Param.Spare,
	}
	err = ibofos.SendIbofos(c, "CREATEARRAY", param)
	if err != nil {
		log.Error(err.Error())
	}
}

func DeleteArray(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		log.Error(err.Error())
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		return
	}
	param := ibofos.ArrayParam{}
	param.Name = c.Param("arrayname")
	err = ibofos.SendIbofos(c, "DELETEARRAY", param)
	if err != nil {
		log.Error(err.Error())
	}
}

func GetArrayDevice(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		log.Error(err.Error())
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		return
	}
	param := ibofos.ArrayParam{}
	param.Name = c.Param("arrayname")
	err = ibofos.SendIbofos(c, "LISTARRAYDEVICE", param)
	if err != nil {
		log.Error(err.Error())
	}
}

func GetArrayInfo(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		log.Error(err.Error())
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		return
	}
	param := ibofos.ArrayParam{}
	param.Name = c.Param("arrayname")
	err = ibofos.SendIbofos(c, "ARRAYINFO", param)
	if err != nil {
		log.Error(err.Error())
	}
}

func CreateArrayMount(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		log.Error(err.Error())
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		return
	}
	param := ibofos.ArrayParam{}
	param.Name = c.Param("arrayname")
	err = ibofos.SendIbofos(c, "MOUNTARRAY", param)
	if err != nil {
		log.Error(err.Error())
	}
}

func DeleteArrayMount(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		log.Error(err.Error())
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		return
	}
	param := ibofos.ArrayParam{}
	param.Name = c.Param("arrayname")
	err = ibofos.SendIbofos(c, "UNMOUNTARRAY", param)
	if err != nil {
		log.Error(err.Error())
	}
}

func ArrayDeviceAdd(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		log.Error(err.Error())
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		return
	}
	param := ibofos.ArrayParam{}
	param.Name = c.Param("arrayname")
	param.Spare = append(param.Spare, ibofos.Device{DeviceName: c.Param("devices")})
	err = ibofos.SendIbofos(c, "ADDDEVICE", param)
	if err != nil {
		log.Error(err.Error())
	}
}

func ArrayDeviceRemove(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		log.Error(err.Error())
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		return
	}
	param := ibofos.ArrayParam{}
	param.Name = c.Param("arrayname")
	param.Spare = append(param.Spare, ibofos.Device{DeviceName: c.Param("devices")})
	err = ibofos.SendIbofos(c, "REMOVEDEVICE", param)
	if err != nil {
		log.Error(err.Error())
	}
}
