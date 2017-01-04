package main

import (
	log "github.com/sirupsen/logrus"
	proc "github.com/c9s/goprocinfo/linux"
)

func printData() {
	stat, err := proc.ReadStat("/proc/stat")
	if err != nil {
	    log.Fatalf("stat read fail: %e", err)
	}

	for _, s := range stat.CPUStats {
	    log.Infof("%v", s)
	}
}
