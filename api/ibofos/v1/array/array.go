package array

import (
	"net/http"
	"pos/client/ibofos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
)

func AddRoutes(r *gin.RouterGroup) {
	route := r.Group("/array")
	{
		route.GET("", GetArray)
	}
}

func GetArray(c *gin.Context) {

	param := ibofos.ArrayParam{
		Name: "POSArray",
	}
	client := ibofos.Requester{
		XrId:           uuid.New().String(),
		IbofServerIP:   viper.GetString("server.ibof.ip"),
		IbofServerPort: viper.GetInt("server.ibof.port"),
		Param:          param,
		ParamType:      ibofos.ArrayParam{},
	}	
	req, res, err := client.Send("ARRAYINFO")
	if err != nil {
		log.Errorf("%s", err.Error())
	}
	log.Debugf("%v", req)
	log.Debugf("%v", res)
	c.JSON(http.StatusOK, res)
	// var command = "GETIBOFOSINFO"

	// systemInfoReq := model.Request{
	// 	RID:     "fromfakeclient",
	// 	COMMAND: command,
	// }

	// reqJSON, err := json.Marshal(systemInfoReq)
	// if err != nil {
	// 	log.Debug("error:", err)
	// }
	// log.Debugf("[%+v]", systemInfoReq)
	// log.Debugf("[%+v]", reqJSON)

	// displaymgr.PrintRequest(string(reqJSON))

	// socketmgr.Connect()
	// resJSON := socketmgr.SendReqAndReceiveRes(string(reqJSON))
	// socketmgr.Close()

	// displaymgr.PrintResponse(command, resJSON, true, true, true)

	// c.JSONP(http.StatusOK, resJSON)

}
