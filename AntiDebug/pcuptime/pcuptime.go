// pcuptime.go
package pcuptime

import (
    "fmt"
    "os"
    "syscall"
)

var (
    kernel32DLL  = syscall.NewLazyDLL("kernel32.dll")
    getTickCount = kernel32DLL.NewProc("GetTickCount")
)

// GetUptimeInSeconds returns the system uptime in seconds, predefined one is 1200 which is 20mins.
func GetUptimeInSeconds() int {
    uptime, _, _ := getTickCount.Call()
    return int(uptime / 1000)
}

// CheckUptime checks if the system uptime is less than a specified duration in seconds and prints a message.
func CheckUptime(durationInSeconds int) {
    uptime := GetUptimeInSeconds()
    if uptime < durationInSeconds {
        fmt.Println("Debug Check: System uptime is less than the specified duration.")
        os.Exit(-1)
    } else {
        fmt.Println("Debug Check: System uptime is greater than or equal to the specified duration.")
    }
}
