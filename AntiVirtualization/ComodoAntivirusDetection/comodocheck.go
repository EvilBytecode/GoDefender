package ComodoAntivirusDetection

import (
    "fmt"
    "os"
    "golang.org/x/sys/windows/registry"
    "syscall"
    "os/exec"
    "strings"
)

// DetectComodoAntivirus checks if Comodo Antivirus is installed or present.
func DetectComodoAntivirus() bool {
    // Check for the installation paths
    comodoPaths := []string{
        "C:\\Program Files\\COMODO\\COMODO Internet Security\\",
        "C:\\Program Files (x86)\\COMODO\\COMODO Internet Security\\",
    }

    for _, path := range comodoPaths {
        if pathExists(path) {
            return true
        }
    }

    // Check for the Comodo Antivirus driver
    driverPath := "C:\\Windows\\System32\\drivers\\cmdguard.sys"
    if pathExists(driverPath) {
        return true
    }

    // Check for Comodo Antivirus registry key
    if checkComodoRegistry() {
        return true
    }

    // Check for Comodo Antivirus service via WMIC
    if checkComodoService() {
        return true
    }

    return false
}

// pathExists checks if a given path exists.
func pathExists(path string) bool {
    _, err := os.Stat(path)
    return !os.IsNotExist(err)
}

// checkComodoRegistry checks for Comodo Antivirus in the registry key.
func checkComodoRegistry() bool {
    comodoKey := `SOFTWARE\COMODO\CIS`
    return registryKeyExists(registry.LOCAL_MACHINE, comodoKey)
}

// checkComodoService checks for the Comodo Antivirus service via WMIC.
func checkComodoService() bool {
    serviceName := "cmdagent"
    return serviceExists(serviceName)
}

// serviceExists checks if a service exists using WMIC.
func serviceExists(serviceName string) bool {
    cmd := exec.Command("wmic", "service", "where", fmt.Sprintf("Name='%s'", serviceName), "get", "Name")

    // Hide the console window
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

    output, err := cmd.Output()
    return err == nil && strings.TrimSpace(string(output)) != ""
}

// registryKeyExists checks if a registry key exists.
func registryKeyExists(root registry.Key, path string) bool {
    key, err := registry.OpenKey(root, path, registry.QUERY_VALUE)
    if err != nil {
        return false
    }
    defer key.Close()
    return true
}
