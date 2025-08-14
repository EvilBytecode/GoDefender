package RunningProcesses

import (
	"syscall"
	"unsafe"
)

var (
	kernel32DLL     = syscall.NewLazyDLL("kernel32.dll")
	enumProcesses   = kernel32DLL.NewProc("K32EnumProcesses")
)

// GetRunningProcessesCount returns the number of currently running processes.
func GetRunningProcessesCount() (int, error) {
	var ids [1024]uint32
	var needed uint32

	r1, _, _ := enumProcesses.Call(
		uintptr(unsafe.Pointer(&ids[0])),
		uintptr(len(ids)*4), // size in bytes
		uintptr(unsafe.Pointer(&needed)),
	)

	if r1 == 0 {
		return 0, syscall.GetLastError()
	}

	return int(needed / 4), nil // each process ID is 4 bytes
}

// CheckRunningProcessesCount checks if the number of currently running processes is less than a specified count.
func CheckRunningProcessesCount(threshold int) (bool, error) {
	count, err := GetRunningProcessesCount()
	if err != nil {
		return false, err
	}
	return count < threshold, nil
}
