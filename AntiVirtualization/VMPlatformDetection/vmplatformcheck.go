package VMPlatformCheck

import (
	"os/exec"
	"strings"
	"syscall"
)

// DetectVMPlatform checks if the system is running on a VMware, Hyper-V, or other virtualization platform.
func DetectVMPlatform() (bool, error) {
	// Check the hardware serial number
	serialCmd := exec.Command("wmic", "bios", "get", "serialnumber")
	serialCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	serialOutput, err := serialCmd.Output()
	if err != nil {
		return false, err
	}

	// Check the model
	modelCmd := exec.Command("wmic", "computersystem", "get", "model")
	modelCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	modelOutput, err := modelCmd.Output()
	if err != nil {
		return false, err
	}

	// Check the manufacturer
	manufacturerCmd := exec.Command("wmic", "computersystem", "get", "manufacturer")
	manufacturerCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	manufacturerOutput, err := manufacturerCmd.Output()
	if err != nil {
		return false, err
	}

	// Convert outputs to lowercase for easy comparison
	serialLower := strings.ToLower(strings.TrimSpace(string(serialOutput)))
	modelLower := strings.ToLower(strings.TrimSpace(string(modelOutput)))
	manufacturerLower := strings.ToLower(strings.TrimSpace(string(manufacturerOutput)))

	// Check if the serial number is "0"
	if serialLower == "0" {
		return true, nil
	}

	// Detect if any of the values indicate a virtual machine (e.g., VMware, VirtualBox, Hyper-V, etc.)
	if strings.Contains(modelLower, "vmware") || strings.Contains(modelLower, "virtual") ||
		strings.Contains(manufacturerLower, "vmware") || strings.Contains(manufacturerLower, "microsoft") ||
		strings.Contains(serialLower, "vmware") || strings.Contains(manufacturerLower, "innotek") ||
		strings.Contains(modelLower, "virtualbox") {
		return true, nil
	}

	return false, nil
}
