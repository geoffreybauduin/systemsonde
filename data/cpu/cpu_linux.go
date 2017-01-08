package cpu

import (
	proc "github.com/c9s/goprocinfo/linux"
)

func Retrieve() (*CPU, error) {
	stat, err := proc.ReadStat("/proc/stat")
	if err != nil {
		return nil, err
	}
	return convert(stat)
}

func convert(stat *proc.Stat) *CPU {
	cpuData := &CPU{

	}
	return cpuData
}