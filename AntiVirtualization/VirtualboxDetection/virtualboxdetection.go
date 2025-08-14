package VirtualboxDetection

import (
	"os/exec"
	"strings"
	"syscall"
)

const CREATE_NO_WINDOW = 0x08000000

// GraphicsCardCheck checks for virtualization software by inspecting the graphics card information.
// It returns true if VirtualBox is detected, otherwise false.
func GraphicsCardCheck() (bool, error) {
	// PowerShell command to get GPU names
	psCmd := `Get-CimInstance -ClassName Win32_VideoController | Select-Object -ExpandProperty Name`

	cmd := exec.Command("powershell", "-NoProfile", "-Command", psCmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: CREATE_NO_WINDOW,
	}

	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	detected := strings.Contains(strings.ToLower(string(output)), "virtualbox")
	return detected, nil
}
