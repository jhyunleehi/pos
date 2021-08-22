package v1

import (
	"pos/api/ibofos/v1/array"
	"pos/api/ibofos/v1/device"
	"pos/api/ibofos/v1/internal"
	"pos/api/ibofos/v1/system"
	"pos/api/ibofos/v1/volume"
	"pos/api/ibofos/v1/qos"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.RouterGroup) {
	route := r.Group("/v1")
	{
		array.AddRoutes(route)
		device.AddRoutes(route)
		internal.AddRoutes(route)
		system.AddRoutes(route)
		volume.AddRoutes(route)
		qos.AddRoutes(route)
	}
}
