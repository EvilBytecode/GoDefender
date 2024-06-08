package InternetCheck

import (
	"log"
	"net"
	"errors"
)

func CheckConnection() (bool, error) {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		err = errors.New("error checking internet connection: " + err.Error())
		log.Printf("[DEBUG] Error checking internet connection: %v", err)
		return false, err
	}
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			log.Printf("[DEBUG] Error closing connection: %v", cerr)
		}
	}()

	return true, nil
}
