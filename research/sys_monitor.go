package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Percent(0, false)
	d, _ := disk.Usage("/")

	fmt.Printf("totoal :%v free %v per %f \n", v.Total, v.Free, v.UsedPercent)
	fmt.Printf("per %f \n", c)
	fmt.Printf("disk %f \n", d.UsedPercent)

}