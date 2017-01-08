package memory

import "github.com/shirou/gopsutil/mem"

type Memory struct {
	Virtual *MemoryStat        `json:"virtual"`
	Swap    *MemoryStat        `json:"swap"`
}

type MemoryStat struct {
	Total        uint64  `json:"total"`
	Used         uint64  `json:"used"`
	Free         uint64  `json:"free"`
	UsedPercent  float64 `json:"used_percent"`
	Available    uint64 `json:"available"`
	Active       uint64 `json:"active"`
	Inactive     uint64 `json:"inactive"`
	Wired        uint64 `json:"wired"`
	Buffers      uint64 `json:"buffers"`
	Cached       uint64 `json:"cached"`
	Writeback    uint64 `json:"writeback"`
	Dirty        uint64 `json:"dirty"`
	WritebackTmp uint64 `json:"writeback_tmp"`
	Shared       uint64 `json:"shared"`
	Slab         uint64 `json:"slab"`
	PageTables   uint64 `json:"page_tables"`
	SwapCached   uint64 `json:"swap_cached"`
	Sin          uint64  `json:"sin"`
	Sout         uint64  `json:"sout"`
}

func Retrieve() (*Memory, error) {
	memory := &Memory{}
	virtual, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	memory.Virtual = convertVirtualMemory(virtual)
	swap, err := mem.SwapMemory()
	if err != nil {
		return nil, err
	}
	memory.Swap = convertSwapMemory(swap)
	return memory, nil
}

func convertVirtualMemory(virtual *mem.VirtualMemoryStat) *MemoryStat {
	data := &MemoryStat{
		Total: virtual.Total,
		Used: virtual.Used,
		Free: virtual.Free,
		UsedPercent: virtual.UsedPercent,
		Available: virtual.Available,
		Active: virtual.Active,
		Inactive: virtual.Inactive,
		Wired: virtual.Wired,
		Buffers: virtual.Buffers,
		Cached: virtual.Cached,
		Writeback: virtual.Writeback,
		Dirty: virtual.Dirty,
		WritebackTmp: virtual.WritebackTmp,
		Shared: virtual.Shared,
		Slab: virtual.Slab,
		PageTables: virtual.PageTables,
		SwapCached: virtual.SwapCached,
	}
	return data
}

func convertSwapMemory(swap *mem.SwapMemoryStat) *MemoryStat {
	data := &MemoryStat{
		Total: swap.Total,
		Used: swap.Used,
		Free: swap.Free,
		UsedPercent: swap.UsedPercent,
		Sin: swap.Sin,
		Sout: swap.Sout,
	}
	return data
}