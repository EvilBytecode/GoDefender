package InternetCheck

import (
    "errors"
    "net"
)

func CheckConnection() (bool, error) {
    conn, err := net.Dial("tcp", "google.com:80")
    if err != nil {
        return false, errors.New("error checking internet connection: " + err.Error())
    }
    defer conn.Close()
    return true, nil
}
