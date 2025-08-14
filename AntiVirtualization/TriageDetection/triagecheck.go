package TriageDetection

import (
	"strings"
	"os/exec"
	"syscall"
)

const CREATE_NO_WINDOW = 0x08000000

// TriageCheck checks for specific hard disk models and returns true if found.
func TriageCheck() (bool, error) {
	// Modern PowerShell command to get disk drive models
	psCmd := `Get-CimInstance -ClassName Win32_DiskDrive | Select-Object -ExpandProperty Model`

	cmd := exec.Command("powershell", "-NoProfile", "-Command", psCmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: CREATE_NO_WINDOW,
	}

	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	outputStr := strings.ToUpper(string(output))
	if strings.Contains(outputStr, "DADY HARDDISK") || strings.Contains(outputStr, "QEMU HARDDISK") {
		return true, nil
	}

	return false, nil
}
