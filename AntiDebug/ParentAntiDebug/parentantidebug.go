package ParentAntiDebug

import (
	"log"
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
func NtQueryProc(handle syscall.Handle, class uint32, info *ProcessInfo, length uint32) error {
	r1, _, err := syscall.Syscall6(ntquery.Addr(), 5, uintptr(handle), uintptr(class), uintptr(unsafe.Pointer(info)), uintptr(length), 0, 0)
	if err != 0 {
		log.Printf("NtQueryInformationProcess failed: %v", err)
		return err
	}
	if r1 != 0 {
		log.Printf("NtQueryInformationProcess failed: unexpected return value: %v", r1)
		return err
	}
	return nil
}

// QueryImageName retrieves the full image name of the process
func QueryImageName(handle syscall.Handle, flags uint32, nameBuffer []uint16, size *uint32) error {
	err := windows.QueryFullProcessImageName(windows.Handle(handle), flags, &nameBuffer[0], size)
	if err != nil {
		log.Printf("QueryFullProcessImageName failed: %v", err)
		return err
	}
	return nil
}

// CurrentProcName returns the name of the current executable
func CurrentProcName() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		log.Printf("os.Executable failed: %v", err)
		return "", err
	}
	return filepath.Base(exePath), nil
}

// ParentAntiDebug checks the parent process if it's explorer.exe or cmd.exe
func ParentAntiDebug() bool {
	const ProcInfo = 0
	var p ProcessInfo
	if err := NtQueryProc(syscall.Handle(windows.CurrentProcess()), ProcInfo, &p, uint32(unsafe.Sizeof(p))); err != nil {
		log.Printf("Error querying process information: %v", err)
		return false
	}
	par := int32(p.InheritedFromPID)
	if par == 0 {
		return false
	}
	handle, err := syscall.OpenProcess(syscall.PROCESS_QUERY_INFORMATION, false, uint32(par))
	if err != nil {
		log.Printf("Error opening process handle: %v", err)
		return false
	}
	defer syscall.CloseHandle(handle)

	buff13 := make([]uint16, windows.MAX_PATH)
	size := uint32(len(buff13))
	if err := QueryImageName(handle, 0, buff13, &size); err != nil {
		log.Printf("Error querying image name: %v", err)
		return false
	}
	parname := filepath.Base(syscall.UTF16ToString(buff13[:size]))

	if parname != "explorer.exe" && parname != "cmd.exe" {
		return true
	} else {
		return false
	}
}
