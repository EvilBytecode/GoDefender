package ParallelsCheck

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// CheckForParallels checks for the presence of Parallels components.
// It returns true if Parallels components are detected, otherwise false.
func CheckForParallels() (bool, error) {
	parallelsDrivers := []string{"prl_sf", "prl_tg", "prl_eth"}
	sys32fold := filepath.Join(os.Getenv("SystemRoot"), "System32")

	files, err := ioutil.ReadDir(sys32fold)
	if err != nil {
		return false, err
	}

	for _, file := range files {
		for _, driver := range parallelsDrivers {
			if strings.Contains(file.Name(), driver) {
				return true, nil
			}
		}
	}

	return false, nil
}
