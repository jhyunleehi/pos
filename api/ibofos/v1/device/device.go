package device

import (
	"net/http"
	"pos/client/ibofos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func AddRoutes(r *gin.RouterGroup) {
	route := r.Group("/devices")
	{
		route.GET("", GetDevice)
		route.GET("/all/scan", GetDevice)
		route.GET("/:devicename/smart", GetDeviceSmart)
	}
}

func GetDevice(c *gin.Context) {

	param := ibofos.SystemParam{}
	client := ibofos.Requester{
		XrId:           uuid.New().String(),
		IbofServerIP:   viper.GetString("server.ibof.ip"),
		IbofServerPort: viper.GetInt("server.ibof.port"),
		Param:          param,
		ParamType:      ibofos.ArrayParam{},
	}
	req, res, err := client.Send("LISTDEVICE")
	if err != nil {
		log.Errorf("%s", err.Error())
	}
	log.Debugf("%+v", req)
	log.Debugf("%+v", res)
	c.JSON(http.StatusOK, res)
}

func GetScanDevice(c *gin.Context) {

	param := ibofos.SystemParam{}
	client := ibofos.Requester{
		XrId:           uuid.New().String(),
		IbofServerIP:   viper.GetString("server.ibof.ip"),
		IbofServerPort: viper.GetInt("server.ibof.port"),
		Param:          param,
		ParamType:      ibofos.ArrayParam{},
	}
	req, res, err := client.Send("SCANDEVICE")
	if err != nil {
		log.Errorf("%s", err.Error())
	}
	log.Debugf("%+v", req)
	log.Debugf("%+v", res)
	c.JSON(http.StatusOK, res)
}