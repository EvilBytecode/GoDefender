package KVMCheck

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// CheckForKVM checks for the presence of Kernel-based Virtual Machine (KVM) components.
// It returns true if KVM components are detected, otherwise false.
// ON 17/06/2024 (Parralers and Qemu Detection Added.)
func CheckForKVM() (bool, error) {
	badDriversList := []string{"balloon.sys", "netkvm.sys", "vioinput", "viofs.sys", "vioser.sys", "qemu-ga", "qemuwmi", "prl_sf", "prl_tg", "prl_eth"}
	systemFolder := filepath.Join(os.Getenv("SystemRoot"), "System32")

	files, err := ioutil.ReadDir(systemFolder)
	if err != nil {
		log.Printf("Error accessing system folder: %v", err)
		return false, err
	}

	for _, file := range files {
		for _, badDriver := range badDriversList {
			if strings.Contains(file.Name(), badDriver) {
				return true, nil
			}
		}
	}
	for _, driver := range badDriversList {
		driverPattern := filepath.Join(systemFolder, driver)
		files, err := filepath.Glob(driverPattern)
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
