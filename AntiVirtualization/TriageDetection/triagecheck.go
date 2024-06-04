package triagecheck

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

// coded by codepulze
// TriageCheckDebug checks for specific hard disk models and prints a message if found.
func TriageCheckDebug() {
    monki := exec.Command("wmic", "diskdrive", "get", "model")
    wowww, _ := monki.Output()
    if strings.Contains(string(wowww), "DADY HARDDISK") || strings.Contains(string(wowww), "QEMU HARDDISK") {
        fmt.Println("Debug check: Specified hard disk model is present, this is a triage.")
		os.Exit(-1)
    } else {
        fmt.Println("Debug check: Specified hard disk model is not present, this isnt a triage.")
    }
}
