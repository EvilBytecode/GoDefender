package ShadowDefenderDetection

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "syscall"
    "golang.org/x/sys/windows/registry"
)

// DetectShadowDefender checks if Shadow Defender is installed or present.
func DetectShadowDefender() bool {
    // Check for the installation paths
    shadowDefenderPaths := []string{
        "C:\\Program Files\\Shadow Defender\\",
        "C:\\Program Files (x86)\\Shadow Defender\\",
    }

    for _, path := range shadowDefenderPaths {
        if pathExists(path) {
            return true
        }
    }

    // Check for Shadow Defender registry keys
    if checkShadowDefenderRegistry() || checkCLSIDRegistry() || checkTypeLibRegistry() ||
       checkUninstallRegistry() || checkSoftwareRegistry() || checkServicesRegistry() ||
       checkDiskptRegistry() || checkUserDatFiles() || checkShadowDefenderService() {
        return true
    }

    return false
}

// pathExists checks if a given path exists.
func pathExists(path string) bool {
    _, err := os.Stat(path)
    return !os.IsNotExist(err)
}

// checkShadowDefenderRegistry checks for the presence of the Shadow Defender-related registry key.
func checkShadowDefenderRegistry() bool {
    runKey := `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`
    return registryValueExists(registry.LOCAL_MACHINE, runKey, "Shadow Defender")
}

// checkCLSIDRegistry checks for the CLSID key associated with Shadow Defender.
func checkCLSIDRegistry() bool {
    clsidKey := `CLSID\{78C3F4BC-C7BC-48E4-AD72-2DD16F6704A9}`
    return registryKeyExists(registry.CLASSES_ROOT, clsidKey)
}

// checkTypeLibRegistry checks for the specific TypeLib entry associated with Shadow Defender.
func checkTypeLibRegistry() bool {
    typeLibKey := `TypeLib\{3A5C2EFF-619A-481D-8D5D-A6968DB02AF1}\1.0\0\win64`
    return registryKeyExists(registry.CLASSES_ROOT, typeLibKey)
}

// checkUninstallRegistry checks for the specific Uninstall entry associated with Shadow Defender.
func checkUninstallRegistry() bool {
    uninstallKey := `SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\{93A07A0D-454E-43d1-86A9-5DE9C5F4411A}`
    return registryKeyExists(registry.LOCAL_MACHINE, uninstallKey)
}

// checkSoftwareRegistry checks for the presence of the Shadow Defender software registry key.
func checkSoftwareRegistry() bool {
    softwareKey := `SOFTWARE\Shadow Defender`
    return registryKeyExists(registry.LOCAL_MACHINE, softwareKey)
}

// checkServicesRegistry checks for the specific services key associated with Shadow Defender.
func checkServicesRegistry() bool {
    servicesKey := `SYSTEM\ControlSet001\Services\{0CBD4F48-3751-475D-BE88-4F271385B672}`
    return registryKeyExists(registry.LOCAL_MACHINE, servicesKey)
}

// checkDiskptRegistry checks for the specific diskpt service key associated with Shadow Defender.
func checkDiskptRegistry() bool {
    diskptKey := `SYSTEM\ControlSet001\Services\diskpt`
    return registryKeyExists(registry.LOCAL_MACHINE, diskptKey)
}

// checkUserDatFiles checks if the user.dat files contain the keyword "Shadow Defender".
func checkUserDatFiles() bool {
    userDir := os.Getenv("USERPROFILE")
    shadowDefenderUserDataPaths := []string{
        filepath.Join(userDir, "Shadow Defender", "user.dat"),
        filepath.Join("C:\\Users\\*", "Shadow Defender", "user.dat"),
    }

    for _, userDataPath := range shadowDefenderUserDataPaths {
        if fileContainsKeyword(userDataPath, "Shadow Defender") {
            return true
        }
    }
    return false
}

// checkShadowDefenderService checks for the existence of the Shadow Defender Service using WMIC.
func checkShadowDefenderService() bool {
    serviceDisplayName := "Shadow Defender Service"
    return serviceExists(serviceDisplayName)
}

// serviceExists checks if a service exists using WMIC.
func serviceExists(displayName string) bool {
    cmd := exec.Command("wmic", "service", "where", fmt.Sprintf("DisplayName='%s'", displayName), "get", "Name")

    // Hide the console window
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
    
    output, err := cmd.Output()
    return err == nil && strings.TrimSpace(string(output)) != ""
}

// fileContainsKeyword checks if a file contains a specific keyword.
func fileContainsKeyword(filePath, keyword string) bool {
    file, err := os.Open(filePath)
    if err != nil {
        return false
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        if strings.Contains(scanner.Text(), keyword) {
            return true
        }
    }
    return false
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

// registryKeyExists checks if a registry key exists.
func registryKeyExists(root registry.Key, path string) bool {
    key, err := registry.OpenKey(root, path, registry.QUERY_VALUE)
    if err != nil {
        return false
    }
    defer key.Close()
    return true
}
