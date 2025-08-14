package VMPlatformCheck

import (
    "strings"
    "golang.org/x/sys/windows/registry"
)

func DetectVMPlatform() (bool, error) {
    serial, err := getBiosSerialFromRegistry()
    if err != nil {
        return false, err
    }
    model, err := getSystemProductName()
    if err != nil {
        return false, err
    }
    manufacturer, err := getSystemManufacturer()
    if err != nil {
        return false, err
    }

    serialLower := strings.ToLower(serial)
    modelLower := strings.ToLower(model)
    manufacturerLower := strings.ToLower(manufacturer)

    if serialLower == "0" {
        return true, nil
    }

    vmIndicators := []string{
        "vmware",
        "virtual",
        "microsoft",
        "innotek",
        "virtualbox",
    }

    for _, indicator := range vmIndicators {
        if strings.Contains(modelLower, indicator) ||
            strings.Contains(manufacturerLower, indicator) ||
            strings.Contains(serialLower, indicator) {
            return true, nil
        }
    }
    return false, nil
}

func getBiosSerialFromRegistry() (string, error) {
    key, err := registry.OpenKey(registry.LOCAL_MACHINE, `HARDWARE\DESCRIPTION\System\BIOS`, registry.QUERY_VALUE)
    if err != nil {
        return "", err
    }
    defer key.Close()
    serial, _, err := key.GetStringValue("SerialNumber")
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(serial), nil
}

func getSystemProductName() (string, error) {
    key, err := registry.OpenKey(registry.LOCAL_MACHINE, `HARDWARE\DESCRIPTION\System`, registry.QUERY_VALUE)
    if err != nil {
        return "", err
    }
    defer key.Close()
    productName, _, err := key.GetStringValue("SystemProductName")
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(productName), nil
}

func getSystemManufacturer() (string, error) {
    key, err := registry.OpenKey(registry.LOCAL_MACHINE, `HARDWARE\DESCRIPTION\System`, registry.QUERY_VALUE)
    if err != nil {
        return "", err
    }
    defer key.Close()
    manufacturer, _, err := key.GetStringValue("SystemManufacturer")
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(manufacturer), nil
}
