package cpu

import (
	"github.com/shirou/gopsutil/cpu"
)

type CPU struct {
	Count        uint64        `json:"count"`
	LogicalCount uint64        `json:"logical_count"`
	Global       *CPUStat      `json:"global"`
	CPUS         []*CPUStat    `json:"cpus"`
}

type CPUStat struct {
	CPU       string  `json:"cpu"`
	User      float64 `json:"user"`
	System    float64 `json:"system"`
	Idle      float64 `json:"idle"`
	Nice      float64 `json:"nice"`
	IOWait    float64 `json:"io_wait"`
	Irq       float64 `json:"irq"`
	SoftIrq   float64 `json:"soft_irq"`
	Steal     float64 `json:"steal"`
	Guest     float64 `json:"guest"`
	GuestNice float64 `json:"guest_nice"`
	Stolen    float64 `json:"stolen"`
	Load	  float64 `json:"load"`
}

func Retrieve() (*CPU, error) {
	cpuData := &CPU{
		CPUS: []*CPUStat{},
	}
	cpuCount, err := cpu.Counts(true)
	if err != nil {
		return nil, err
	}
	cpuData.LogicalCount = uint64(cpuCount)
	cpuCount, err = cpu.Counts(false)
	if err != nil {
		return nil, err
	}
	cpuData.Count = uint64(cpuCount)
	cpuTimes, err := cpu.Times(true)
	if err != nil {
		return nil, err
	}
	cpuPercent, err := cpu.Percent(0, true)
	if err != nil {
		return nil, err
	}
	for i, cpuTime := range cpuTimes {
		singleCpu := getSingleCPUData(&cpuTime)
		singleCpu.Load = cpuPercent[i] / 100.0
		cpuData.CPUS = append(cpuData.CPUS, singleCpu)
	}
	return cpuData, nil
}

func getSingleCPUData(cpuTime *cpu.TimesStat) *CPUStat {
	return &CPUStat{
		CPU: cpuTime.CPU,
		User: cpuTime.User,
		System: cpuTime.System,
		Idle: cpuTime.Idle,
		Nice: cpuTime.Nice,
		IOWait: cpuTime.Iowait,
		Irq: cpuTime.Irq,
		SoftIrq: cpuTime.Softirq,
		Steal: cpuTime.Steal,
		Guest: cpuTime.Guest,
		GuestNice: cpuTime.GuestNice,
		Stolen: cpuTime.Stolen,
	}
}