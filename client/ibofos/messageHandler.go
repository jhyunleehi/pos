package ibofos

import (
	"encoding/json"
	"errors"	
	"time"

	log "github.com/sirupsen/logrus"
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
	ParamType      interface{}
	IbofServerIP   string
	IbofServerPort int
}
func init(){
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)	
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

	switch param := rq.ParamType.(type) {
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
