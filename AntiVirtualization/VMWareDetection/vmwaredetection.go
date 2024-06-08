package VMWareDetection

import (
    "log"
    "os/exec"
    "strings"
    "syscall"
)

// GraphicsCardCheck checks for virtualization software by inspecting the graphics card information.
// It returns true if VMware is detected, otherwise false.
func GraphicsCardCheck() (bool, error) {
    cmd := exec.Command("wmic", "path", "win32_VideoController", "get", "name")
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
    gpu, err := cmd.Output()
    if err != nil {
        log.Println("Error executing command:", err)
        return false, err
    }
    detected := strings.Contains(strings.ToLower(string(gpu)), "vmware")
    if detected {
        return true, nil
    }
    return false, nil
}