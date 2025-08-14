package USBCheck

import (
	"golang.org/x/sys/windows/registry"
)

// PluggedIn checks if USB devices were ever plugged in by querying the USBSTOR registry keys.
// It returns true if any USB storage device entries are found.
func PluggedIn() (bool, error) {
	// First try checking the Services\USBSTOR key
	key1, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\ControlSet001\Services\USBSTOR`, registry.READ)
	if err == nil {
		defer key1.Close()
		return true, nil
	}

	// If not present, try Enum\USBSTOR to check history of connected USB devices
	key2, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\ControlSet001\Enum\USBSTOR`, registry.READ)
	if err != nil {
		return false, err
	}
	defer key2.Close()

	// Read subkey names under USBSTOR (each one typically represents a connected device)
	subkeys, err := key2.ReadSubKeyNames(-1)
	if err != nil {
		return false, err
	}

	// Return true if at least one USB device entry is present
	return len(subkeys) > 0, nil
}
