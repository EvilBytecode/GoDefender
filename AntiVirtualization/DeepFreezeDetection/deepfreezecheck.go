package DeepFreezeDetection

import (
    "os"
    "strings"
    "golang.org/x/sys/windows/registry"
    "golang.org/x/sys/windows/svc/mgr"
)

// DetectDeepFreeze checks if Deep Freeze is installed or present.
func DetectDeepFreeze() bool {
    // Check for installation paths
    deepFreezePaths := []string{
        "C:\\Program Files\\Faronics\\Deep Freeze\\",
        "C:\\Program Files (x86)\\Faronics\\Deep Freeze\\",
    }
    for _, path := range deepFreezePaths {
        if pathExists(path) {
            return true
        }
    }

    // Check for Deep Freeze driver
    driverPath := "C:\\Persi0.sys"
    if pathExists(driverPath) {
        return true
    }

    // Check registry keys
    if checkHelpDirRegistry() {
        return true
    }

    if checkAutoRecoverMOFs() {
        return true
    }

    // Check for Deep Freeze service using Windows API
    if checkDeepFreezeService() {
        return true
    }

    return false
}

func pathExists(path string) bool {
    _, err := os.Stat(path)
    return !os.IsNotExist(err)
}

func checkHelpDirRegistry() bool {
    helpDirKey := `SOFTWARE\Classes\TypeLib\{C5D763D9-2422-4B2D-A425-02D5BD016239}\1.0\HELPDIR`
    return registryKeyExists(registry.LOCAL_MACHINE, helpDirKey)
}

func checkAutoRecoverMOFs() bool {
    cimomKey := `SOFTWARE\Microsoft\Wbem\CIMOM`
    key, err := registry.OpenKey(registry.LOCAL_MACHINE, cimomKey, registry.QUERY_VALUE)
    if err != nil {
        return false
    }
    defer key.Close()

    mofs, _, err := key.GetStringsValue("Autorecover MOFs")
    if err != nil {
        return false
    }

    for _, mof := range mofs {
        if strings.Contains(strings.ToLower(mof), "faronics") {
            return true
        }
    }
    return false
}

func checkDeepFreezeService() bool {
    return serviceExists("DFServ")
}

// Updated serviceExists uses Windows Service Manager API
func serviceExists(serviceName string) bool {
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

    const serviceRunning = 4
    return status.State == serviceRunning
}

func registryKeyExists(root registry.Key, path string) bool {
    key, err := registry.OpenKey(root, path, registry.QUERY_VALUE)
    if err != nil {
        return false
    }
    defer key.Close()
    return true
}
