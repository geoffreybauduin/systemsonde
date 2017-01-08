package main

import (
	"github.com/lght/systemsonde/data"
	"time"
	log "github.com/sirupsen/logrus"
	"encoding/json"
	"fmt"
)

func main() {
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(new(MyJSONFormatter))
	timer := time.Tick(5 * time.Second)
	for _ = range timer {
		data.Print()
	}
}

type MyJSONFormatter struct {

}

func (f *MyJSONFormatter) Format(entry *log.Entry) ([]byte, error) {
	var serialized []byte
	var err error
	if entry.Data["data"] != nil {
		serialized, err = json.Marshal(entry.Data["data"])
	} else {
		serialized, err = json.Marshal(entry.Data)
	}
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}