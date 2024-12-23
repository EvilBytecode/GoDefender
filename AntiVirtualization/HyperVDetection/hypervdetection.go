package HyperVCheck

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"golang.org/x/sys/windows/registry"
)

// DetectHyperV checks if Hyper-V is installed using registry, service, and process checks.
func DetectHyperV() (bool, error) {
	// 1. Check Registry for Hyper-V
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\Virtualization`, registry.QUERY_VALUE)
	if err == nil {
		defer key.Close()
		enabled, _, err := key.GetIntegerValue("Enabled")
		if err == nil && enabled == 1 {
			return true, nil
		}
	}

	// 2. Check if Hyper-V services are running (vmms and vmbus)
	if isServiceExisting("vmms") || isServiceExisting("vmbus") {
		return true, nil
	}

	// 3. Check if Hyper-V processes are running (vmms.exe and vmbus.sys)
	if isProcessRunning("vmms.exe") || isProcessRunning("vmbus.sys") {
		return true, nil
	}

	return false, nil
}

// isServiceExisting checks if a specific service exists and is running
func isServiceExisting(serviceName string) bool {
	cmd := exec.Command("wmic", "service", "where", fmt.Sprintf("name='%s'", serviceName), "get", "name")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	output, err := cmd.Output()
	if err != nil {
		return false
	}

	return strings.Contains(strings.ToLower(string(output)), strings.ToLower(serviceName))
}

// isProcessRunning checks if a specific process is running
func isProcessRunning(processName string) bool {
	cmd := exec.Command("tasklist")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	output, err := cmd.Output()
	if err != nil {
		return false
	}

	return strings.Contains(strings.ToLower(string(output)), strings.ToLower(processName))
}
