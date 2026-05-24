package main

import (
	"fmt"
	"log"
)

func main() {
	// get current kernel
	kernelVersion := getKernel()
	fmt.Println("Current kernel: ", kernelVersion)

	//get all kernels
	allKernels := getAllKernels()
	fmt.Println("All Available Kernels:")
	for _, k := range allKernels {
		fmt.Printf(" - %s\n", k)
	}

	// get OS Version
	osVersion, err := getOsVersion()
	if err != nil {
		log.Fatalf("Error fetching OS Version: %v", err)
	}
	fmt.Println("OS Version: ", osVersion)
}
