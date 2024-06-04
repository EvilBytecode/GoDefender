package VMWare

import (
    "os/exec"
	"os"
	"syscall"
    "strings"
)

func GraphicsCardCheck() bool {
    cmd := exec.Command("wmic", "path", "win32_VideoController", "get", "name")
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
    gpu, err := cmd.Output()
    if err != nil {
        return false
    }
    detected := strings.Contains(strings.ToLower(string(gpu)), "vmware")
    if detected {
        os.Exit(-1)
    }
	
    return detected
}