package pcuptime

import (
    "syscall"
)

var (
    kernel32DLL  = syscall.NewLazyDLL("kernel32.dll")
    getTickCount = kernel32DLL.NewProc("GetTickCount")
)

// GetUptimeInSeconds returns the system uptime in seconds.
func GetUptimeInSeconds() (int, error) {
    uptime, _, err := getTickCount.Call()
    if err != nil && err.Error() != "The operation completed successfully." {
        return 0, err
    }
    return int(uptime / 1000), nil
}

// CheckUptime checks if the system uptime is less than a specified duration in seconds.
func CheckUptime(durationInSeconds int) (bool, error) {
    uptime, err := GetUptimeInSeconds()
    if err != nil {
        return false, err
    }

    if uptime < durationInSeconds {
        return true, nil
    } else {
        return false, nil
    }
}
