package ibofos

import (
	"pos/api/ibofos/v1"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	ibofos := r.Group("/ibofos")
	v1.ApplyRoutes(ibofos)
}
