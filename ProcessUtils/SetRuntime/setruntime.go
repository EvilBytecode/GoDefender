package RuntimeDetector

import (
	"runtime"
)

// OperatingSystem type defines the supported operating systems.
type OperatingSystem string

const (
	Unknown OperatingSystem = "unknown"
	Windows OperatingSystem = "windows"
	Linux   OperatingSystem = "linux"
	MacOS   OperatingSystem = "darwin"
)

var currentOS OperatingSystem

// DetectOS detects and returns the current operating system.
func DetectOS() OperatingSystem {
	os := runtime.GOOS
	switch os {
	case "windows":
		return Windows
	case "linux":
		return Linux
	case "darwin":
		return MacOS
	default:
		return Unknown
	}
}

// SetRuntime sets the runtime environment to the specified operating system.
func SetRuntime(os OperatingSystem) {
	currentOS = os
}

// GetCurrentOS returns the currently set operating system or detects it if not set.
func GetCurrentOS() OperatingSystem {
	if currentOS == "" {
		currentOS = DetectOS()
	}
	return currentOS
}
