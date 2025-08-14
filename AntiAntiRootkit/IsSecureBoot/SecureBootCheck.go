package SecureBootCheck

import (
	"os/exec"
	"strings"
	"syscall"
)

const CREATE_NO_WINDOW = 0x08000000

// IsSecureBootEnabled runs Confirm-SecureBootUEFI silently (no console window) and returns true if Secure Boot is enabled.
func IsSecureBootEnabled() bool {
	cmd := exec.Command(
		"powershell.exe",
		"-NoProfile",
		"-NonInteractive",
		"-Command",
		"Confirm-SecureBootUEFI",
	)
	// Hide the console window entirely
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: CREATE_NO_WINDOW,
	}

	output, err := cmd.Output()
	if err != nil {
		return false
	}

	outStr := strings.TrimSpace(strings.ToLower(string(output)))
	return outStr == "true"
}
