package main

import (
	"fmt"
	"log"

	"github.com/brandonthis-that/home-server/internal/system"
)

func main() {
	// get current kernel
	kernelVersion := system.GetKernel()
	fmt.Println("Current kernel: ", kernelVersion)

	//get all kernels
	allKernels := system.GetAllKernels()
	fmt.Println("All Available Kernels:")
	for _, k := range allKernels {
		fmt.Printf(" - %s\n", k)
	}

	// get OS Version
	osVersion, err := system.GetOsVersion()
	if err != nil {
		log.Fatalf("Error fetching OS Version: %v", err)
	}
	fmt.Println("OS Version: ", osVersion)

	uptime, err := system.GetUpTime()
	if err != nil {
		fmt.Printf("Error reading uptime: %v\n", err)
		return
	}

	days := int(uptime.Hours()) / 24
	hours := int(uptime.Hours()) % 24
	minutes := int(uptime.Minutes()) % 60
	seconds := int(uptime.Seconds()) % 60

	fmt.Printf("Server Uptime: %d days, %02d:%02d:%02d\n", days, hours, minutes, seconds)
}
