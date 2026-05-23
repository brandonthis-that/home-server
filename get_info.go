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
		os.Exit(0)
	}
	kernel_info := make([]byte, 1024)
	kernel_file.Read(kernel_info)
	kernel_file.Close()
	kernel_list := strings.Split(string(kernel_info), " ")
	return strings.TrimSpace(kernel_list[2])
}

func get_all_kernels() []string {
	entries, err := os.ReadDir("/boot")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(0)
	}
	var kernels []string
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "vmlinuz") {
			kernels = append(kernels, entry.Name())
		}
	}
	return kernels
}

func getOsVersion() (string, error) {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "Unknown", fmt.Errorf("reading os-release: %w", err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		if val, ok := strings.CutPrefix(line, "PRETTY_NAME="); ok {
			return strings.Trim(val, `'"`), nil
		}
	}
	return "Unknown", nil
}
