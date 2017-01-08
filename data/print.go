package data

import (
	log "github.com/sirupsen/logrus"
	"github.com/lght/systemsonde/data/cpu"
	"github.com/lght/systemsonde/data/memory"
	"time"
)

// Get returns the data collected for further use
func Get() (*Data, error) {
	data := &Data{}
	data.Timestamp = time.Now()
	cpuData, err := cpu.Retrieve()
	if err != nil {
		return nil, err
	}
	data.CPU = cpuData
	memoryData, err := memory.Retrieve()
	if err != nil {
		return nil, err
	}
	data.Memory = memoryData
	return data, nil
}

// Print prints on the stdout the data collected
func Print() error {
	data, err := Get()
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"data": data,
	}).Info("Data")
	return nil
}
