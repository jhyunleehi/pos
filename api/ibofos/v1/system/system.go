package system

import (
	"net/http"
	"pos/client/ibofos"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func AddRoutes(r *gin.RouterGroup) {
	route := r.Group("/system")
	{
		route.GET("", GetIbofosInfo)
		route.POST("/mount", MountIbofos)
		route.POST("", RunIbofos)
		route.DELETE("", ExitIbofos)
		route.DELETE("/mount", UnMountIbofos)
	}
}

func GetIbofosInfo(c *gin.Context) {
	param := ibofos.SystemParam{}
	client, err := ibofos.Setup(param )
	if err != nil {
		log.Errorf("%v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	req, res, err := client.Send("GETIBOFOSINFO")
	if err != nil {
		log.Errorf("%s", err.Error())
	}
	log.Debugf("%v", req)
	log.Debugf("%v", res)
	c.JSON(http.StatusOK, res)
}

func MountIbofos (c *gin.Context) {
	param := ibofos.SystemParam{}
	client, err := ibofos.Setup(param )
	if err != nil {
		log.Errorf("%v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	req, res, err := client.Send("MOUNTIBOFOS")
	if err != nil {
		log.Errorf("%s", err.Error())
	}
	log.Debugf("%v", req)
	log.Debugf("%v", res)
	c.JSON(http.StatusOK, res)
}

func RunIbofos (c *gin.Context) {
	param := ibofos.SystemParam{}
	client, err := ibofos.Setup(param )
	if err != nil {
		log.Errorf("%v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	req, res, err := client.Send("RUNIBOFOS")
	if err != nil {
		log.Errorf("%s", err.Error())
	}
	log.Debugf("%v", req)
	log.Debugf("%v", res)
	c.JSON(http.StatusOK, res)
}


func ExitIbofos (c *gin.Context) {
	param := ibofos.SystemParam{}
	client, err := ibofos.Setup(param )
	if err != nil {
		log.Errorf("%v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	req, res, err := client.Send("EXITIBOFOS")
	if err != nil {
		log.Errorf("%s", err.Error())
	}
	log.Debugf("%v", req)
	log.Debugf("%v", res)
	c.JSON(http.StatusOK, res)
}



func UnMountIbofos (c *gin.Context) {
	param := ibofos.SystemParam{}
	client, err := ibofos.Setup(param )
	if err != nil {
		log.Errorf("%v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	req, res, err := client.Send("UNMOUNTIBOFOS")
	if err != nil {
		log.Errorf("%s", err.Error())
	}
	log.Debugf("%v", req)
	log.Debugf("%v", res)
	c.JSON(http.StatusOK, res)
}

