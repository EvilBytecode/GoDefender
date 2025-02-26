package USBCheck

import (
	"log"
	"os/exec"
	"strings"
	"syscall"
)

// yes this detects https://tria.ge lol
// PluggedIn checks if USB devices were ever plugged in and returns true if found.
func PluggedIn() (bool, error) {
	usbcheckcmd := exec.Command("reg", "query", "HKLM\\SYSTEM\\ControlSet001\\Services\\USBSTOR")
	usbcheckcmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	outputusb, err := usbcheckcmd.CombinedOutput()
	if err != nil {
		log.Printf("Debug Check: Error running reg query command: %v", err)
		usbcheckcmd := exec.Command("reg", "query", "HKLM\\SYSTEM\\ControlSet001\\Enum\\USBSTOR")
		usbcheckcmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

		outputusb, err = usbcheckcmd.CombinedOutput()
		if err != nil {
			log.Printf("Debug Check: Error running reg query command: %v", err)
			return false, err
		}
	}

	usblines := strings.Split(string(outputusb), "\n")
	pluggedusb := 0
	for _, line := range usblines {
		if strings.TrimSpace(line) != "" {
			pluggedusb++
		}
	}

	if pluggedusb < 1 {
		return false, nil
	}

	return true, nil
}
