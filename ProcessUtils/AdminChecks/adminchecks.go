package AdminCheck

import (
	"os"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

// IsAdmin checks if the current process is running as an administrator.
func IsAdmin() bool {
	ret, _, _ := syscall.NewLazyDLL("shell32.dll").NewProc("IsUserAnAdmin").Call()
	return ret != 0
}

// ElevateProcess attempts to elevate the process to run with administrative privileges.
// Reference: https://stackoverflow.com/questions/31558066/how-to-ask-for-administer-privileges-on-windows-with-go
func ElevateProcess() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 // SW_NORMAL

	// Call ShellExecute and ignore any error
	_ = windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
}
