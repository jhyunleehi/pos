package volume

import (
	"pos/client/ibofos"

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
		route.DELETE("/:volumename/mount", DeleteVolumeMount)
		route.DELETE("/:volumename", DeleteVolume)
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
	err = ibofos.SendIbofos(c, "LISTVOLUME", volumeparam)
	if err != nil {
		log.Error(err.Error())
	}
}

func GetVolumeMaxcount(c *gin.Context) {
	volumeparam := ibofos.VolumeParam{}
	err := ibofos.SendIbofos(c, "GETMAXVOLUMECOUNT", volumeparam)
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
	err = ibofos.SendIbofos(c, "CREATEVOLUME", volumeparam)
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
	err = ibofos.SendIbofos(c, "DELETEVOLUME", volumeparam)
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
	err = ibofos.SendIbofos(c, "MOUNTVOLUME", volumeparam)
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
	err = ibofos.SendIbofos(c, "UNMOUNTVOLUME", volumeparam)
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
	err = ibofos.SendIbofos(c, "UPDATEVOLUMEQOS", volumeparam)
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
	err = ibofos.SendIbofos(c, "RENAMEVOLUME", volumeparam)
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
	err = ibofos.SendIbofos(c, "GETHOSTNQN", volumeparam)
	if err != nil {
		log.Error(err.Error())
	}
}
