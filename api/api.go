package api

import (
	"pos/api/ibofos"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	route := r.Group("/api")
	{
		ibofos.SetRoutes(route)
	}
}
