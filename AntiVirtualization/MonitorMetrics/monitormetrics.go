package MonitorMetrics

import (
    "fmt"
    "syscall"
	"os"
)

func IsScreenSmall() bool {
    getSystemMetrics := syscall.NewLazyDLL("user32.dll").NewProc("GetSystemMetrics")
    width, _, _ := getSystemMetrics.Call(0)
    height, _, _ := getSystemMetrics.Call(1)

    isSmall := width < 800 || height < 600
    if isSmall {
        fmt.Println("Debug Check: Screen size is considered small.")
		os.Exit(-1)
    } else {
        fmt.Println("Debug Check: Screen size is not considered small.")
    }
    return isSmall
}
