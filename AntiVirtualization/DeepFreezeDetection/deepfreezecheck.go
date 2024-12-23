package DeepFreezeDetection

import (
    "fmt"
    "os"
    "golang.org/x/sys/windows/registry"
    "syscall"
    "os/exec"
    "strings"
)

// DetectDeepFreeze checks if Deep Freeze is installed or present.
func DetectDeepFreeze() bool {
    // Check for the installation paths
    deepFreezePaths := []string{
        "C:\\Program Files\\Faronics\\Deep Freeze\\",
        "C:\\Program Files (x86)\\Faronics\\Deep Freeze\\",
    }

    for _, path := range deepFreezePaths {
        if pathExists(path) {
            return true
        }
    }

    // Check for the Deep Freeze driver
    driverPath := "C:\\Persi0.sys"
    if pathExists(driverPath) {
        return true
    }

    // Check for Deep Freeze registry key at HKEY_LOCAL_MACHINE\SOFTWARE\Classes\TypeLib
    if checkHelpDirRegistry() {
        return true
    }

    // Check for Autorecover MOFs containing "Faronics" under HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Wbem\CIMOM
    if checkAutoRecoverMOFs() {
        return true
    }

    // Check for Deep Freeze service via WMIC
    if checkDeepFreezeService() {
        return true
    }

    return false
}

// pathExists checks if a given path exists.
func pathExists(path string) bool {
    _, err := os.Stat(path)
    return !os.IsNotExist(err)
}

// checkHelpDirRegistry checks for Deep Freeze in the HELPDIR registry key.
func checkHelpDirRegistry() bool {
    helpDirKey := `SOFTWARE\Classes\TypeLib\{C5D763D9-2422-4B2D-A425-02D5BD016239}\1.0\HELPDIR`
    return registryKeyExists(registry.LOCAL_MACHINE, helpDirKey)
}

// checkAutoRecoverMOFs checks for Autorecover MOFs containing "Faronics".
func checkAutoRecoverMOFs() bool {
    cimomKey := `SOFTWARE\Microsoft\Wbem\CIMOM`
    key, err := registry.OpenKey(registry.LOCAL_MACHINE, cimomKey, registry.QUERY_VALUE)
    if err != nil {
        return false
    }
    defer key.Close()

    // Get the "Autorecover MOFs" value
    mofs, _, err := key.GetStringsValue("Autorecover MOFs")
    if err != nil {
        return false
    }

    // Check if any of the MOF paths contain the "Faronics" keyword
    for _, mof := range mofs {
        if strings.Contains(strings.ToLower(mof), "faronics") {
            return true
        }
    }
    return false
}

// checkDeepFreezeService checks for the Deep Freeze service via WMIC.
func checkDeepFreezeService() bool {
    serviceName := "DFServ"
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
