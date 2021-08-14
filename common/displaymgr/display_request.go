package displaymgr

import (
	"cli/cmd/globals"

	log "github.com/sirupsen/logrus"
)

func PrintRequest(reqJSON string) {
	if globals.IsJSONReq {
		log.Print(string(reqJSON))
	}
}
