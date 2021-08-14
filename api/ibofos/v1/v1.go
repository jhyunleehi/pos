package v1

import (
	"pos/api/ibofos/v1/array"
	"pos/api/ibofos/v1/device"
	"pos/api/ibofos/v1/internal"
	"pos/api/ibofos/v1/system"
	"pos/api/ibofos/v1/volume"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	array.ApplyRoutes(v1)
	device.ApplyRoutes(v1)
	internal.ApplyRoutes(v1)
	system.ApplyRoutes(v1)
	volume.ApplyRoutes(v1)
}
