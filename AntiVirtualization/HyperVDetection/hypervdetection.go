package HyperVCheck

import (
	"os/exec"
	"strings"
	"syscall"
    "golang.org/x/sys/windows/registry"
	"golang.org/x/sys/windows/svc/mgr"
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

// isServiceExisting checks if a specific service exists and is running using Windows service manager.
func isServiceExisting(serviceName string) bool {
	m, err := mgr.Connect()
	if err != nil {
		return false
	}
	defer m.Disconnect()

	svc, err := m.OpenService(serviceName)
	if err != nil {
		return false
	}
	defer svc.Close()

	status, err := svc.Query()
	if err != nil {
		return false
	}

	// SERVICE_RUNNING = 4
	return status.State == 4
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
