package volume

import (
	"net/http"
	"pos/client/ibofos"
	"pos/common/events"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func AddRoutes(r *gin.RouterGroup) {
	route := r.Group("/volumes")
	{
		route.GET("", GetVolume)
		route.GET("/", GetVolume)
		route.GET("/:volumename", GetVolume)
		route.GET("/maxcount", GetVolumeMaxcount)
		route.GET("/:volumename/hostnqn", GetVolumeHostNQN)
		route.POST("", CreateVolume)
		route.POST("/:volumename/mount", CreateVolumeMount)
		route.DELETE("/:volumename", DeleteVolume)
		route.DELETE("/:volumename/mount", DeleteVolumeMount)
		route.PATCH("/:volumename/qos", UpdateVolumeQos)
		route.PATCH("/:volumename", UpdateVolumeRename)
	}
}

func GetVolume(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		log.Error(err.Error())
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		return
	}
	volumeparam := ibofos.VolumeParam{
		Array: bodyparam.Param.Array,
	}
	err = SendIbofosAPI(c, "LISTVOLUME", volumeparam)
	if err != nil {
		log.Error(err.Error())
	}
}

func GetVolumeMaxcount(c *gin.Context) {
	volumeparam := ibofos.VolumeParam{}
	err := SendIbofosAPI(c, "GETMAXVOLUMECOUNT", volumeparam)
	if err != nil {
		log.Error(err.Error())
	}
}

//POST c.ShouldBindJSON(&model.createvol)
func CreateVolume(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		log.Error(err.Error())
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		return
	}
	volumeparam := ibofos.VolumeParam{
		Array:   bodyparam.Param.Array,
		Name:    bodyparam.Param.Name,
		Size:    bodyparam.Param.Size,
		Maxbw:   bodyparam.Param.Maxbw,
		Maxiops: bodyparam.Param.Maxiops,
	}
	err = SendIbofosAPI(c, "CREATEVOLUME", volumeparam)
	if err != nil {
		log.Error(err.Error())
	}
}

func DeleteVolume(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		log.Error(err.Error())
		return
	}
	volumeparam := ibofos.VolumeParam{}
	volumeparam.Array = bodyparam.Param.Array
	volumeparam.Name = c.Param("volumename")
	err = SendIbofosAPI(c, "DELETEVOLUME", volumeparam)
	if err != nil {
		log.Error(err.Error())
	}
}

func CreateVolumeMount(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		log.Error(err.Error())
		return
	}
	volumeparam := ibofos.VolumeParam{}
	volumeparam.Array = bodyparam.Param.Array
	volumeparam.Name = c.Param("volumename")
	err = SendIbofosAPI(c, "MOUNTVOLUME", volumeparam)
	if err != nil {
		log.Error(err.Error())
	}
}

func DeleteVolumeMount(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		log.Error(err.Error())
		return
	}
	volumeparam := ibofos.VolumeParam{}
	volumeparam.Array = bodyparam.Param.Array
	volumeparam.Name = c.Param("volumename")
	err = SendIbofosAPI(c, "UNMOUNTVOLUME", volumeparam)
	if err != nil {
		log.Error(err.Error())
	}
}

func UpdateVolumeQos(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		log.Error(err.Error())
		return
	}
	volumeparam := ibofos.VolumeParam{}
	volumeparam.Array = bodyparam.Param.Array
	volumeparam.Maxiops = bodyparam.Param.Maxiops
	volumeparam.Maxbw = bodyparam.Param.Maxbw
	volumeparam.Name = c.Param("volumename")
	err = SendIbofosAPI(c, "UPDATEVOLUMEQOS", volumeparam)
	if err != nil {
		log.Error(err.Error())
	}
}

func UpdateVolumeRename(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		log.Error(err.Error())
		return
	}
	volumeparam := ibofos.VolumeParam{}
	volumeparam.Array = bodyparam.Param.Array
	volumeparam.NewName = bodyparam.Param.NewName
	volumeparam.Name = c.Param("volumename")
	err = SendIbofosAPI(c, "RENAMEVOLUME", volumeparam)
	if err != nil {
		log.Error(err.Error())
	}
}

func GetVolumeHostNQN(c *gin.Context) {
	bodyparam := ibofos.BodyParam{}
	err := c.ShouldBindJSON(&bodyparam)
	if err != nil && err.Error() != "EOF" {
		c.JSON(400, gin.H{"error": "Request JSON Parsing Error", "origin": err.Error()})
		log.Error(err.Error())
		return
	}
	volumeparam := ibofos.VolumeParam{}
	volumeparam.Array = bodyparam.Param.Array	
	volumeparam.Name = c.Param("volumename")
	err = SendIbofosAPI(c, "GETHOSTNQN", volumeparam)
	if err != nil {
		log.Error(err.Error())
	}
}

func SendIbofosAPI(c *gin.Context, command string, param interface{}) error {
	log.Debugf("[%s][%+v]", command, param)
	client, err := ibofos.Setup(param)
	if err != nil {
		log.Errorf("%s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	req, res, err := client.Send(command)
	if err != nil {
		log.Errorf("%s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	if res.Result.Status.Code != 0 {
		status, err := events.GetStatusInfo(res.Result.Status.Code)
		if err != nil {
			log.Errorf("%s", err.Error())
			c.JSON(http.StatusInternalServerError, res)
			return err
		}
		log.Errorf("%+v", status)
		c.JSON(http.StatusInternalServerError, status)
		return err
	}
	log.Debugf("%+v", req)
	log.Debugf("%+v", res)
	c.JSON(http.StatusOK, res)
	return nil
}