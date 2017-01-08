package data

import (
	"github.com/lght/systemsonde/data/cpu"
	"time"
	"github.com/lght/systemsonde/data/memory"
)

type Data struct {
	Timestamp time.Time        `json:"timestamp"`
	CPU       *cpu.CPU        `json:"cpu"`
	Memory	*memory.Memory	`json:"memory"`
}
