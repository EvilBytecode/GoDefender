package QEMUCheck

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// CheckForQEMU checks for the presence of QEMU components.
// It returns true if QEMU components are detected, otherwise false.
func CheckForQEMU() (bool, error) {
	qemuDrivers := []string{"qemu-ga", "qemuwmi"}
	sys32 := filepath.Join(os.Getenv("SystemRoot"), "System32")

	files, err := ioutil.ReadDir(sys32)
	if err != nil {
		return false, err
	}

	for _, file := range files {
		for _, driver := range qemuDrivers {
			if strings.Contains(file.Name(), driver) {
				return true, nil
			}
		}
	}

	return false, nil
}
