package RecentFileActivity

import (
	"log"
	"os"
	"io/ioutil"
	"path/filepath"
)

// RecentFileActivityCheck checks recent file activity.
func RecentFileActivityCheck() (bool, error) {
	recdir := filepath.Join(os.Getenv("APPDATA"), "microsoft", "windows", "recent")
	files, err := ioutil.ReadDir(recdir)
	if err != nil {
		log.Printf("Debug Check: Error reading recent file activity directory: %v", err)
		return false, err
	}

	if len(files) < 20 {
		return true, nil
	}

	return false, nil
}
