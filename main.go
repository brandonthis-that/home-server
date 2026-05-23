package main

import (
	"fmt"
)

func main() {
	kernel_version := get_kernel()
	allKernels := get_all_kernels()
	fmt.Println(kernel_version)
	for _, k := range allKernels {
		fmt.Println(k)
	}

	osVersion, err := getOsVersion()
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(osVersion)
	fmt.Println(osVersion)
	fmt.Println("All good")
}
