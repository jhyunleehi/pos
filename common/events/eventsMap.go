package events

import (
	"errors"
	"io/ioutil"
	"os"
	"pos/model"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var eventsmap PosEvents

type info2 struct {
	Code     int    `yaml:"code"`
	Level    string `yaml:"level"`
	Message  string `yaml:"message"`
	Problem  string `yaml:"problem,omitempty"`
	Solution string `yaml:"solution,omitempty"`
}

type module struct {
	Name    string  `yaml:"name"`
	Count   int     `yaml:"count"`
	Idstart int     `yaml:"idStart"`
	Idend   int     `yaml:"idEnd"`
	Info    []info2 `yaml:"info"`
}
type PosEvents struct {
	Modules []module `yaml:"modules"`
}

func init() {
	eventsyaml := "./" + "common/events/events.yaml"
	if _, err := os.Stat(eventsyaml); os.IsNotExist(err) {
		eventsyaml = "../../" + "common/events/events.yaml"
	}

	yfile, err := ioutil.ReadFile(eventsyaml)
	if err != nil {
		log.Fatal(err)
	}
	err2 := yaml.Unmarshal(yfile, &eventsmap)
	if err2 != nil {
		log.Fatal(err2)
	}
}

func Setup() {
}

func GetStatusInfo(code int) (model.Status, error) {
	var status model.Status
	status.CODE = code
	totMods := len(eventsmap.Modules)

	for i := 0; i < totMods; i++ {
		if code >= eventsmap.Modules[i].Idstart && code <= eventsmap.Modules[i].Idend {
			totInfo := len(eventsmap.Modules[i].Info)
			for j := 0; j < totInfo; j++ {
				if eventsmap.Modules[i].Info[j].Code == code {
					status.MODULE = eventsmap.Modules[i].Name
					status.DESCRIPTION = eventsmap.Modules[i].Info[j].Message
					status.PROBLEM = eventsmap.Modules[i].Info[j].Problem
					status.SOLUTION = eventsmap.Modules[i].Info[j].Solution
					status.LEVEL = eventsmap.Modules[i].Info[j].Level

					return status, nil
				}
			}
		}
	}

	err := errors.New("there is no event info")

	return status, err
}