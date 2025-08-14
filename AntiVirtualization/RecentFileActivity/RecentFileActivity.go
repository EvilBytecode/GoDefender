package RecentFileActivity

import (
	
	"os"
	"io/ioutil"
	"path/filepath"
)

// RecentFileActivityCheck checks recent file activity.
func RecentFileActivityCheck() (bool, error) {
	recdir := filepath.Join(os.Getenv("APPDATA"), "microsoft", "windows", "recent")
	files, err := ioutil.ReadDir(recdir)
	if err != nil {
		return false, err
	}

	if len(files) < 20 {
		return true, nil
	}

	return false, nil
}
