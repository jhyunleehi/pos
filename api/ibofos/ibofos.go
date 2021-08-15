package ibofos

import (
	"pos/api/ibofos/v1"
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.RouterGroup) {
	route := r.Group("/ibofos")
	{
		v1.AddRoutes(route)
	}
}
