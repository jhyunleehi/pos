package array

import (
	"encoding/json"
	"pos/api/ibofos"
	"pos/model"
	"pos/common/displaymgr"
	"pos/common/socketmgr"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ApplyRoutes( r *gin.Engine){
	parray  :=  r.Group("/array")
	parray.GET("", GetArray) 
}

func GetArray (c *gin.Context) {

		var command = "GETIBOFOSINFO"

		systemInfoReq := model.Request{
			RID:     "fromfakeclient",
			COMMAND: command,
		}

		reqJSON, err := json.Marshal(systemInfoReq)
		if err != nil {
			log.Debug("error:", err)
		}

		displaymgr.PrintRequest(string(reqJSON))

		socketmgr.Connect()
		resJSON := socketmgr.SendReqAndReceiveRes(string(reqJSON))
		socketmgr.Close()

		displaymgr.PrintResponse(command, resJSON, globals.IsDebug, globals.IsJSONRes)
}