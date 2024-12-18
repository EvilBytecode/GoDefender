package AnyRunDetection

import (
	"log"
	"net"
	"strings"
)

// AnyRunDetection checks for the "52-54-00" MAC address prefix and returns true if found.
func AnyRunDetection() (bool, error) {
	// Retrieve all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Printf("Error getting network interfaces: %v", err)
		return false, err
	}

	// Loop through each interface and check the MAC address
	for _, iface := range interfaces {
		macAddress := iface.HardwareAddr.String() // Fetch the MAC address
		if strings.HasPrefix(strings.ToUpper(strings.ReplaceAll(macAddress, "-", ":")), "52:54:00") {
			return true, nil
		}
	}

	return false, nil
}
