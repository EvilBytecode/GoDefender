package PowerShellCheck

import (
	"bytes"
	"os/exec"
	"strings"
	"syscall"
)

const CREATE_NO_WINDOW = 0x08000000

// RunPowerShellCommand runs a PowerShell command silently and returns the trimmed output or an error.
func RunPowerShellCommand(psCommand string) (string, error) {
	// Use powershell.exe with -NoProfile and -NonInteractive to reduce side effects
	cmd := exec.Command("powershell.exe", "-NoProfile", "-NonInteractive", "-Command", psCommand)

	// Hide the console window entirely
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: CREATE_NO_WINDOW,
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Return trimmed output
	return strings.TrimSpace(out.String()), nil
}

// ContainsInPowerShellOutput runs a PowerShell command and checks if the output contains the target substring (case-insensitive).
func ContainsInPowerShellOutput(psCommand, target string) (bool, error) {
	output, err := RunPowerShellCommand(psCommand)
	if err != nil {
		return false, err
	}

	return strings.Contains(strings.ToLower(output), strings.ToLower(target)), nil
}
