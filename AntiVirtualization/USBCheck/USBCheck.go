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

	outputusb, err1 := usbcheckcmd.CombinedOutput()
	if err1 != nil {
		log.Printf("Debug Check: Error running reg query command: %v", err1)
		usbcheckcmd = exec.Command("reg", "query", "HKLM\\SYSTEM\\ControlSet001\\Enum\\USBSTOR")
		usbcheckcmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

		outputusb, err1 = usbcheckcmd.CombinedOutput() // Reuse outputusb to avoid redeclaring it
		if err1 != nil {
			log.Printf("Debug Check: Error running reg query command: %v", err1)
			return false, err1
		}
	}

	// Use outputusb to check if any USB devices were found
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
