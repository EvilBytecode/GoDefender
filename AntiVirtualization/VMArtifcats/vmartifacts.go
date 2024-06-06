package artifactsdetector

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func BadVMFilesDetection() {
	badFileNames := []string{"VBoxMouse.sys", "VBoxGuest.sys", "VBoxSF.sys", "VBoxVideo.sys", "vmmouse.sys", "vboxogl.dll"}
	badDirs := []string{`C:\Program Files\VMware`, `C:\Program Files\oracle\virtualbox guest additions`}
	sys32, err := filepath.Glob(filepath.Join(os.Getenv("SystemRoot"), "System32", "*"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, sys32files := range sys32 {
		fileName := filepath.Base(sys32files)
		for _, badFileName := range badFileNames {
			if strings.EqualFold(fileName, badFileName) {
				fmt.Printf("Debug Check: VM Artifact file : %s\n", fileName)
				os.Exit(-1)
			}
		}
	}
	for _, badDir := range badDirs {
		if _, err := os.Stat(badDir); err == nil {
			fmt.Printf("Debug Check: Bad VM Artifact directory: %s\n", badDir)
			os.Exit(-1)
		}
	}
	fmt.Println("Debug Check: VM Artifcats havent been found.")
}
