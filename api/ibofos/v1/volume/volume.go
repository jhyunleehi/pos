package volume

import (
	"net/http"
	"pos/client/ibofos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func AddRoutes(r *gin.RouterGroup) {
	route := r.Group("/volumes") 
	{
		route.GET("", GetVolume)
		route.GET("/:volumename", GetVolume)
	}
}

func GetVolume(c *gin.Context) {

	param := ibofos.VolumeParam{
		Name:  c.Param("volumename"),
		Array: c.Query("array"),
	}
	client := ibofos.Requester{
		XrId:           uuid.New().String(),
		IbofServerIP:   viper.GetString("server.ibof.ip"),
		IbofServerPort: viper.GetInt("server.ibof.port"),
		Param:          param,
		ParamType:      ibofos.ArrayParam{},
	}
	req, res, err := client.Send("LISTVOLUME")
	if err != nil {
		log.Errorf("%s", err.Error())
	}
	log.Debugf("%+v", req)
	log.Debugf("%+v", res)
	c.JSON(http.StatusOK, res)
}
