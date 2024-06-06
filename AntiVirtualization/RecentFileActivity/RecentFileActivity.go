package RecentFileActivity

import (
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
)

func RecentFileActivityCheck() {
	recdir := filepath.Join(os.Getenv("APPDATA"), "microsoft", "windows", "recent")
	files, err := ioutil.ReadDir(recdir)
	if err != nil {
		fmt.Println("Debug Check: ermm error:", err)
		return
	}
	if len(files) < 20 {
		fmt.Println("Debug Check: RECENT FILE ACTIVITY CHECK FAILED!")
	    os.Exit(-1)
	}
	fmt.Println("Debug Check: Recent file activity check passed!")
}
