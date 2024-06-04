package parentantidebug

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"
	"golang.org/x/sys/windows"
)

var (
	ntdll   = syscall.NewLazyDLL("ntdll.dll")
	ntquery = ntdll.NewProc("NtQueryInformationProcess")
)

type ProcessInfo struct {
	Res1             uintptr
	PebAddr          uintptr
	Res2             [2]uintptr
	PID              uintptr
	InheritedFromPID uintptr
}

// NtQueryProc queries process information
func NtQueryProc(handle syscall.Handle, class uint32, info *ProcessInfo, length uint32) {
	syscall.Syscall6(ntquery.Addr(), 5, uintptr(handle), uintptr(class), uintptr(unsafe.Pointer(info)), uintptr(length), 0, 0)
}

// QueryImageName retrieves the full image name of the process
func QueryImageName(handle syscall.Handle, flags uint32, nameBuffer []uint16, size *uint32) {
	windows.QueryFullProcessImageName(windows.Handle(handle), flags, &nameBuffer[0], size)
}

// CurrentProcName returns the name of the current executable
func CurrentProcName() string {
	exePath, _ := os.Executable()
	return filepath.Base(exePath)
}

// ParentAntiDebug checks the parent process and exits if it's not explorer.exe or cmd.exe
func ParentAntiDebug() {
	const ProcInfo = 0
	var p ProcessInfo
	NtQueryProc(syscall.Handle(windows.CurrentProcess()), ProcInfo, &p, uint32(unsafe.Sizeof(p)))
	par := int32(p.InheritedFromPID)
	if par == 0 {
		return
	}
	handle, _ := syscall.OpenProcess(syscall.PROCESS_QUERY_INFORMATION, false, uint32(par))
	defer syscall.CloseHandle(handle)
	buff13 := make([]uint16, windows.MAX_PATH)
	size := uint32(len(buff13))
	QueryImageName(handle, 0, buff13, &size)
	pa1231 := syscall.UTF16ToString(buff13[:size])
	parname := filepath.Base(pa1231)

	if parname != "explorer.exe" && parname != "cmd.exe" {
		fmt.Printf("Debug Check: Parent process (%s) is not in the whitelist\n", parname)
		os.Exit(-1)
	} else {
		fmt.Printf("Debug Check: Parent process (%s) is in the whitelist\n", parname)
	}
}
