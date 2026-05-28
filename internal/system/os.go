package system

import (
	"fmt"
	"os"
	"strings"
)

func GetOsVersion() (string, error) {
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
