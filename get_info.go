package main

import (
	"fmt"
	"os"
	"strings"
)

func get_kernel() string {
	kernel_file, err := os.Open("/proc/version")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	kernel_info := make([]byte, 1024)
	kernel_file.Read(kernel_info)
	kernel_file.Close()
	kernel_list := strings.Split(string(kernel_info), " ")
	return strings.TrimSpace(kernel_list[2])
}
