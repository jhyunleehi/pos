package displaymgr

import (	

	log "github.com/sirupsen/logrus"
)

func PrintRequest(reqJSON string) {
	if IsJSONReq {
		log.Print(string(reqJSON))
	}
}
