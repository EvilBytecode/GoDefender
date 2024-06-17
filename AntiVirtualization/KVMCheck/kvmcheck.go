package KVMCheck

import (
	"log"
	"os"
	"path/filepath"
)

// CheckForKVM checks for the presence of Kernel-based Virtual Machine (KVM) components.
// It returns true if KVM components are detected, otherwise false.
func CheckForKVM() (bool, error) {
	badDriversList := []string{"balloon.sys", "netkvm.sys", "vioinput", "viofs.sys", "vioser.sys"}
	for _, driver := range badDriversList {
		files, err := filepath.Glob(filepath.Join(os.Getenv("SystemRoot"), "System32", driver))
		if err != nil {
			log.Printf("Error accessing system files for %s: %v", driver, err)
			continue
		}
		if len(files) > 0 {
			return true, nil
		}
	}
	return false, nil
}
