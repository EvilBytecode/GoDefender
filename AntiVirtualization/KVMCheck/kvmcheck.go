package kvmcheck

import (
	"os"
	"fmt"
	"path/filepath"
)

// CheckForKVM checks for the presence of Kernel-based Virtual Machine (KVM) components.
// It returns true if KVM components are detected, otherwise false.
func CheckForKVM() bool {
	badDriversList := []string{"balloon.sys", "netkvm.sys", "vioinput", "viofs.sys", "vioser.sys"}
	for _, driver := range badDriversList {
		files, err := filepath.Glob(filepath.Join(os.Getenv("SystemRoot"), "System32", driver))
		if err != nil {
			fmt.Println("Debug Check: Error accessing system files:", err)
			continue
		}
		if len(files) > 0 {
			fmt.Println("Debug Check: Kernel-based Virtual Machine (KVM) components detected:", driver)
			return true
		}
	}
	fmt.Println("Debug Check: No Kernel-based Virtual Machine (KVM) components detected.")
	return false
}