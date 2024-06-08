package TriageDetection

import (
	"log"
	"os/exec"
	"strings"
)

// TriageCheckDebug checks for specific hard disk models and returns true if found.
func TriageCheck() (bool, error) {
	monki := exec.Command("wmic", "diskdrive", "get", "model")
	wowww, err := monki.Output()
	if err != nil {
		log.Printf("Error running wmic command: %v", err)
		return false, err
	}

	if strings.Contains(string(wowww), "DADY HARDDISK") || strings.Contains(string(wowww), "QEMU HARDDISK") {
		return true, nil
	}

	return false, nil
}
