package SandboxieDetection

import (
    "os"
    "golang.org/x/sys/windows/registry"
    "golang.org/x/sys/windows/svc/mgr"
)

// DetectSandboxie checks if Sandboxie is installed or present.
func DetectSandboxie() bool {
    // Check for the installation paths
    sandboxiePaths := []string{
        "C:\\Program Files\\Sandboxie\\",
        "C:\\Program Files (x86)\\Sandboxie\\",
    }

    for _, path := range sandboxiePaths {
        if pathExists(path) {
            return true
        }
    }

    // Check for the existence of the Sandboxie service
    serviceName := "SbieSvc" // Sandboxie service name
    if isServiceExisting(serviceName) {
        return true
    }

    // Check for Sandboxie registry keys
    if checkSandboxieRegistry() {
        return true
    }

    return false
}

// pathExists checks if a given path exists.
func pathExists(path string) bool {
    _, err := os.Stat(path)
    return !os.IsNotExist(err)
}

// isServiceExisting checks if a service exists using Windows service manager.
func isServiceExisting(serviceName string) bool {
    m, err := mgr.Connect()
    if err != nil {
        return false // fail quietly
    }
    defer m.Disconnect()

    s, err := m.OpenService(serviceName)
    if err != nil {
        return false // service not found
    }
    s.Close()
    return true // service found
}

// checkSandboxieRegistry checks for the presence of various Sandboxie-related registry keys.
func checkSandboxieRegistry() bool {
    // Check if Sandboxie is listed in the uninstall registry key
    uninstallKey := `SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\Sandboxie`
    if registryKeyExists(registry.LOCAL_MACHINE, uninstallKey) {
        return true
    }

    // Check for the auto-run Sandman entry in HKCU
    autorunKey := `Software\Microsoft\Windows\CurrentVersion\Run`
    if registryValueExists(registry.CURRENT_USER, autorunKey, "SandboxiePlus_AutoRun") {
        return true
    }

    // Check for the shell integration for running files/folders sandboxed
    sandboxedKey := `Software\Classes\*\shell\sandbox`
    if registryKeyExists(registry.CURRENT_USER, sandboxedKey) {
        return true
    }

    folderSandboxedKey := `Software\Classes\Folder\shell\sandbox`
    if registryKeyExists(registry.CURRENT_USER, folderSandboxedKey) {
        return true
    }

    return false
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

// registryValueExists checks if a registry key has a specific value.
func registryValueExists(root registry.Key, path, valueName string) bool {
    key, err := registry.OpenKey(root, path, registry.QUERY_VALUE)
    if err != nil {
        return false
    }
    defer key.Close()

    _, _, err = key.GetStringValue(valueName)
    return err == nil
}
