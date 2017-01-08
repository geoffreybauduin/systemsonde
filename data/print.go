package data

import (
	log "github.com/sirupsen/logrus"
	"encoding/json"
	"github.com/lght/systemsonde/data/cpu"
)

func Print() {
	cpuData, err := cpu.Retrieve()
	if err != nil {
		log.Fatalf("Cannot retrieve cpu stats: %e", err)
	}
	data, err := json.MarshalIndent(cpuData, "", "    ")
	if err != nil {
		log.Fatalf("stat read fail: %e", err)
	}
	log.Infof("%s", data)
}
