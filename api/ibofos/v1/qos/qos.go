package qos

import (
	"pos/client/ibofos"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AddRoutes(r *gin.RouterGroup) {
	route := r.Group("/qos")
	{
		route.GET("", ListQosPolicies)
		route.POST("", CreateQosVolumePolicy)
		route.DELETE("", ResetQosVolumePolicy)
	}
}

func ListQosPolicies(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		log.Error(err.Error())
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		return
	}
	param := ibofos.QosParam{}
	param.Array = bodyparam.Param.Array
	param.Vol = append(param.Vol, ibofos.Volume{VolumeName: bodyparam.Param.Name})
	err = ibofos.SendIbofos(c, "LISTQOSPOLICIES", param)
	if err != nil {
		log.Error(err.Error())
	}
}

func CreateQosVolumePolicy(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		log.Error(err.Error())
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		return
	}
	param := ibofos.QosParam{}
	param.Array = bodyparam.Param.Array
	param.Vol = append(param.Vol, ibofos.Volume{VolumeName: bodyparam.Param.Name})
	param.Maxbw = bodyparam.Param.Maxbw
	param.Maxiops = bodyparam.Param.Maxiops
	param.Minbw = bodyparam.Param.Minbw
	param.Miniops = bodyparam.Param.Miniops
	err = ibofos.SendIbofos(c, "CREATEQOSVOLUMEPOLICY", param)
	if err != nil {
		log.Error(err.Error())
	}
}

func ResetQosVolumePolicy(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		log.Error(err.Error())
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		return
	}
	param := ibofos.QosParam{}
	param.Array = bodyparam.Param.Array
	param.Vol = append(param.Vol, ibofos.Volume{VolumeName: bodyparam.Param.Name})	
	err = ibofos.SendIbofos(c, "RESETQOSVOLUMEPOLICY", param)
	if err != nil {
		log.Error(err.Error())
	}
}
