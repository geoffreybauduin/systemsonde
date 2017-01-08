package cpu

import (
	"github.com/shirou/gopsutil/cpu"
)

type CPU struct {
	Count        uint64        `json:"count"`
	LogicalCount uint64        `json:"logical_count"`
	CPUByName    map[string]*CPUStat    `json:"cpu_by_name"`
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
	Load      float64 `json:"load"`
}

var lastCallStatPerCPU map[string]*CPUStat

func init() {
	lastCallStatPerCPU = map[string]*CPUStat{}
	cpuDatas, err := getCPUData()
	if err != nil {
		panic(err)
	}
	for _, cpuData := range cpuDatas {
		lastCallStatPerCPU[cpuData.CPU] = cpuData
	}
}

// Retrieve retrieves the last metrics on the CPU
func Retrieve() (*CPU, error) {
	cpuData := &CPU{
		CPUByName: map[string]*CPUStat{},
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
	cpus, err := getCPUData()
	if err != nil {
		return nil, err
	}
	for _, c := range cpus {
		c.Load = getLoadFor(c)
		lastCallStatPerCPU[c.CPU] = c
		cpuData.CPUByName[c.CPU] = c
	}
	return cpuData, nil
}

func getCPUData() ([]*CPUStat, error) {
	cpuDatas := []*CPUStat{}
	cpuTimes, err := cpu.Times(true)
	if err != nil {
		return nil, err
	}
	for _, cpuTime := range cpuTimes {
		singleCpu := getSingleCPUData(&cpuTime)
		cpuDatas = append(cpuDatas, singleCpu)
	}
	return cpuDatas, nil
}

func getLoadFor(c *CPUStat) (float64) {
	t1All, t1Busy := getCpuTime(lastCallStatPerCPU[c.CPU])
	t2All, t2Busy := getCpuTime(c)
	if t2Busy <= t1Busy {
		return 0
	}
	if t2All <= t1All {
		return 1
	}
	return (t2Busy - t1Busy) / (t2All - t1All)
}

func getCpuTime(c *CPUStat) (float64, float64) {
	busy := c.User + c.System + c.Nice + c.IOWait + c.Irq +
	c.SoftIrq + c.Steal + c.Guest + c.GuestNice + c.Stolen
	return busy + c.Idle, busy
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