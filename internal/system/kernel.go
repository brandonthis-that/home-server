package system

import (
	"fmt"
	"os"
	"strings"
)

func GetKernel() string {
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

func GetAllKernels() []string {
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
