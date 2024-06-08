package USBCheck

import (
	"log"
	"os/exec"
	"strings"
)

// PluggedIn checks if USB devices was ever plugged in and returns true if found.
func PluggedIn() (bool, error) {
	usbcheckcmd := exec.Command("reg", "query", "HKLM\\SYSTEM\\ControlSet001\\Enum\\USBSTOR")
	outputusb, err := usbcheckcmd.CombinedOutput()
	if err != nil {
		log.Printf("Debug Check: Error running reg query command: %v", err)
		return false, err
	}

	usblines := strings.Split(string(outputusb), "\n")
	pluggedusb := 0
	for _, line := range usblines {
		if strings.TrimSpace(line) != "" {
			pluggedusb++
		}
	}

	if pluggedusb < 1 {
		if pluggedusb < 0 {
			return false, nil
		}
		return false, nil
	}

	return true, nil
}
