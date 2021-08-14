package displaymgr

import (
	"pos/api/globals"

	log "github.com/sirupsen/logrus"
)

func PrintRequest(reqJSON string) {
	if globals.IsJSONReq {
		log.Print(string(reqJSON))
	}
}
