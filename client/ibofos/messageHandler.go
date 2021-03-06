package ibofos

import (
	"encoding/json"
	"errors"
	"net/http"
	"pos/common/events"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	errLocked    = errors.New("Locked out buddy")
	ErrSending   = errors.New("Sending error")
	ErrReceiving = errors.New("Receiving error")
	ErrJson      = errors.New("Json error")
	ErrRes       = errors.New("Response error")
	ErrConn      = errors.New("iBoF Connection Error")
	ErrJsonType  = errors.New("Json Type Validation Error")
	//mutex        = &sync.Mutex{}
)

type Requester struct {
	XrId           string
	Param          interface{}
	//ParamType      interface{}
	IbofServerIP   string
	IbofServerPort int
}
func init(){
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)	
}

func Setup(param interface{}) (Requester, error ){
	req := Requester {
		XrId:           uuid.New().String(),
		
		IbofServerIP:   viper.GetString("server.ibof.ip"),
		IbofServerPort: viper.GetInt("server.ibof.port"),
		Param:          param,
		//ParamType:      paramtype,
	}
	return req, nil
}

func SendIbofos(c *gin.Context, command string, param interface{}) error {
	log.Debugf("[%s][%+v]", command, param)
	client, err := Setup(param)
	if err != nil {
		log.Errorf("%s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	req, res, err := client.Send(command)
	if err != nil {
		log.Errorf("%s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	if res.Result.Status.Code != 0 {
		status, err := events.GetStatusInfo(res.Result.Status.Code)
		if err != nil {
			log.Errorf("%s", err.Error())
			c.JSON(http.StatusInternalServerError, res)
			return err
		}
		log.Errorf("%+v", status)
		c.JSON(http.StatusInternalServerError, status)
		return err
	}
	log.Debugf("%+v", req)
	log.Debugf("%+v", res)
	c.JSON(http.StatusOK, res)
	return nil
}


func (rq Requester) Send(command string) (Request, Response, error) {
	iBoFRequest := Request{
		Command: command,
		Rid:     rq.XrId,
	}

	err := rq.checkJsonType()
	if err != nil {
		return iBoFRequest, Response{}, err
	} else {
		iBoFRequest.Param = rq.Param
		res, err := rq.sendIBoF(iBoFRequest)
		return iBoFRequest, res, err
	}
}

func (rq Requester) checkJsonType() error {
	var err error
	marshalled, _ := json.Marshal(rq.Param)

	switch param := rq.Param.(type) {
	case ArrayParam:
		err = json.Unmarshal(marshalled, &param)
	case DeviceParam:
		err = json.Unmarshal(marshalled, &param)
	case VolumeParam:
		err = json.Unmarshal(marshalled, &param)
	case InternalParam:
		err = json.Unmarshal(marshalled, &param)
	case SystemParam:
		err = json.Unmarshal(marshalled, &param)
	case RebuildParam:
		err = json.Unmarshal(marshalled, &param)
	case LoggerParam:
		err = json.Unmarshal(marshalled, &param)
	case WBTParam:
		err = json.Unmarshal(marshalled, &param)
	}

	if err != nil {
		log.Debugf("checkJsonType : ", ErrJsonType.Error())
		err = ErrJsonType
	}

	return err
}

func (rq Requester) sendIBoF(iBoFRequest Request) (Response, error) {
	conn, err := rq.ConnectToIBoFOS()
	if err != nil {
		return Response{}, ErrConn
	}
	defer rq.DisconnectToIBoFOS(conn)

	log.Debugf("sendIBoF : %+v", iBoFRequest)

	marshaled, _ := json.Marshal(iBoFRequest)
	err = rq.WriteToIBoFSocket(conn, marshaled)
	if err != nil {
		log.Infof("sendIBoF write error : %v", err)
		return Response{}, ErrSending
	}

	for {
		temp, err := rq.ReadFromIBoFSocket(conn)
		if err != nil {
			log.Errorf("sendIBoF read error : %v", err)
			return Response{},ErrReceiving
		} else {
			log.Infof("Response From iBoF : %s", temp.String())
		}

		response := Response{}

		d := json.NewDecoder(&temp)
		d.UseNumber()

		if err = d.Decode(&response); err != nil {
			log.Fatal(err)
		}

		if err != nil {
			log.Infof("Response Unmarshal Error : %v", err)
			return Response{}, ErrJson
		} else {
			response.LastSuccessTime = time.Now().UTC().Unix()
			return response, nil
		}
	}
}