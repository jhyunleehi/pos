package volume

import (
	"net/http"
	"pos/client/ibofos"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
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
	client, err := ibofos.Setup(param)
	if err != nil {
		log.Errorf("%v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	req, res, err := client.Send("LISTVOLUME")
	if err != nil {
		log.Errorf("%s", err.Error())
	}
	log.Debugf("%+v", req)
	log.Debugf("%+v", res)
	c.JSON(http.StatusOK, res)
}

//POST c.ShouldBindJSON(&model.createvol)
