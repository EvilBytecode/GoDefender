package VirtualboxDetection

import (
    "log"
    "os/exec"
	"syscall"
    "strings"
)

// GraphicsCardCheck checks for virtualization software by inspecting the graphics card information.
// It returns true if VirtualBox is detected, otherwise false.
func GraphicsCardCheck() (bool, error) {
    cmd := exec.Command("wmic", "path", "win32_VideoController", "get", "name")
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
    gpu, err := cmd.Output()
    if err != nil {
        log.Println("Error executing command:", err)
        return false, err
    }
    detected := strings.Contains(strings.ToLower(string(gpu)), "virtualbox")
    if detected {
        return true, nil
    }
    return false, nil
}