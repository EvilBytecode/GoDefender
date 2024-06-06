package USBCheck

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func PluggedIn() {
    usbcheckcmd := exec.Command("reg", "query", "HKLM\\SYSTEM\\ControlSet001\\Enum\\USBSTOR")
    outputusb, err := usbcheckcmd.CombinedOutput()
    if err != nil {
        fmt.Println("Debug Check: Error:", err)
        return
    }
    usblines := strings.Split(string(outputusb), "\n")
    pluggedusb := 0
    for _, line := range usblines {
        if strings.TrimSpace(line) != "" {
            pluggedusb++
        }
    }
    if pluggedusb < 1 {
        fmt.Println("Debug Check: USB Check passed:", false)
        if pluggedusb < 0 {
            fmt.Println("Debug Check: Less than 0 mounted USB devices")
            os.Exit(-1)
        }
        return
    }
    fmt.Println("Debug Check: USB Check Passed:", true)
}
