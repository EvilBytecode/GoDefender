package vmcheck

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func GraphicsCardCheck() bool {
	cmd := exec.Command("wmic", "path", "win32_VideoController", "get", "name")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	gpu, err := cmd.Output()
	if err != nil {
		return false
	}
	detected := strings.Contains(strings.ToLower(string(gpu)), "vmware") || strings.Contains(strings.ToLower(string(gpu)), "virtualbox")
	if detected {
		os.Exit(-1)
	}

	return detected
}
