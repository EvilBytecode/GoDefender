package VMArtifacts
import (
	"fmt"
	"strings"
	"os"
	"path/filepath"
)
// VMArtifactsDetect checks for the presence of files and directories related to VirtualBox and VMware.
// It returns true if any of the bad files or directories are detected, otherwise false.
func VMArtifactsDetect() bool {
	badFileNames := []string{"VBoxMouse.sys", "VBoxGuest.sys", "VBoxSF.sys", "VBoxVideo.sys", "vmmouse.sys", "vboxogl.dll"}
	badDirs := []string{`C:\Program Files\VMware`, `C:\Program Files\oracle\virtualbox guest additions`}

	system32Folder := os.Getenv("SystemRoot") + `\System32`
	files, err := filepath.Glob(filepath.Join(system32Folder, "*"))
	if err != nil {
		fmt.Printf("Error accessing System32 folder: %v\n", err)
		return false
	}

	for _, file := range files {
		fileName := strings.ToLower(filepath.Base(file))
		for _, badFileName := range badFileNames {
			if fileName == strings.ToLower(badFileName) {
				return true
			}
		}
	}

	for _, badDir := range badDirs {
		if _, err := os.Stat(badDir); err == nil {
			return true
		}
	}

	return false
}
