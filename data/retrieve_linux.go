package data

import (
	log "github.com/sirupsen/logrus"
	proc "github.com/c9s/goprocinfo/linux"
	"encoding/json"
)

func Print() {
	stat, err := proc.ReadStat("/proc/stat")
	if err != nil {
	    log.Fatalf("stat read fail: %e", err)
	}
	data, err := json.MarshalIndent(stat, "", "    ")
	if err != nil {
		log.Fatalf("stat read fail: %e", err)
	}
	log.Infof("%s", data) 
}
