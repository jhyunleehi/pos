package system

import (
	"net/http"
	"pos/client/ibofos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
)

func AddRoutes(r *gin.RouterGroup) {
	route := r.Group("/system")
	{
		route.GET("", GetSystem)
	}
}

func GetSystem(c *gin.Context) {

	param := ibofos.SystemParam{}
	client := ibofos.Requester{
		XrId:           uuid.New().String(),
		IbofServerIP:   viper.GetString("server.ibof.ip"),
		IbofServerPort: viper.GetInt("server.ibof.port"),
		Param:          param,
		ParamType:      ibofos.ArrayParam{},
	}	
	req, res, err := client.Send("GETIBOFOSINFO")
	if err != nil {
		log.Errorf("%s", err.Error())
	}
	log.Debugf("%v", req)
	log.Debugf("%v", res)
	c.JSON(http.StatusOK, res)
}
