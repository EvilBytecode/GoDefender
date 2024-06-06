package InternetCheck

import (
    "fmt"
    "net"
    "os"
)

func CheckConnection() {
    _, err := net.Dial("tcp", "google.com:80")
    if err == nil {
        fmt.Println("Debug Check: [!] Internet connection is active.")
    } else {
        fmt.Println("Debug Check: INTERNET CONNECTION CHECK FAILED!")
        os.Exit(-1)
    }
}