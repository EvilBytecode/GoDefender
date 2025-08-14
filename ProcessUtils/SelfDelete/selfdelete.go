package SelfDelete

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

const CREATE_NO_WINDOW = 0x08000000

// SelfDelete schedules the deletion of the running executable with no visible window.
func SelfDelete() {
	exePath, err := os.Executable()
	if err != nil {
		return
	}

	exePathQuoted := `"` + exePath + `"`

	// Create a temporary batch file to delete the executable
	batFile := exePath + ".del.bat"
	batContent := fmt.Sprintf(`
@echo off
:Repeat
del %s >nul 2>&1
if exist %s goto Repeat
del "%%~f0"
`, exePathQuoted, exePathQuoted)

	_ = os.WriteFile(batFile, []byte(batContent), 0644)

	// Run the batch file hidden (no window)
	cmd := exec.Command("cmd", "/C", batFile)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: CREATE_NO_WINDOW,
	}
	_ = cmd.Start()

	// Exit so the batch file can delete us
	os.Exit(0)
}
