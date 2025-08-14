package CyberCapture

import (
	"os"
	"path/filepath"
)

func CreateDirectory() bool {
	programFiles := os.Getenv("ProgramFiles")
	if programFiles == "" {
		programFiles = `C:\Program Files`
	}

	dir := filepath.Join(programFiles, "antvirusdefender2025")

	err := os.MkdirAll(dir, 0755)
	return err == nil
}
