package MonitorMetrics

import (
	"syscall"
)

// IsScreenSmall checks if the screen size is considered small.
func IsScreenSmall() (bool, error) {
	getSystemMetrics := syscall.NewLazyDLL("user32.dll").NewProc("GetSystemMetrics")
	width, _, err := getSystemMetrics.Call(0)
	if err != nil {
		return false, err
	}
	height, _, err := getSystemMetrics.Call(1)
	if err != nil {
		return false, err
	}

	isSmall := width < 800 || height < 600
	if isSmall {
		return true, nil
	} else {
		return false, nil
	}
}
