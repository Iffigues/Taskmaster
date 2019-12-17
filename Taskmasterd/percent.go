package main

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
)

func percent(i int32) (str string) {
	str = "start"
	if p, err := process.NewProcess(i); err == nil {
		if cpu, err := p.CPUPercent(); err == nil {
			str = str + " CPU: " + fmt.Sprintf("%f ", cpu)
		}
		if memory, err := p.MemoryPercent(); err == nil {
			str = str + " RAM: " + fmt.Sprintf("%f ", memory)
		}
	}
	return
}
