package v1

import (
	"pos/api/ibofos/array"
	"pos/api/ibofos/device"
	"pos/api/ibofos/internal"
	"pos/api/ibofos/system"
	"pos/api/ibofos/volume"

	"gihub.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	array.ApplyRoutes(v1)
	device.ApplyRoutes(v1)
	internal.ApplyRoutes(v1)
	system.ApplyRoutes(v1)
	volume.ApplyRoutes(v1)
}
