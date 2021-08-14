package api

import (
	"pos/api/ibofos"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		ibofos.ApplyRoutes(api)
	}
}
