package RunningProcesses

import (
	"log"
	"syscall"
	"unsafe"
)

var (
	kernel32DLL = syscall.NewLazyDLL("kernel32.dll")
	enumProcesses         = kernel32DLL.NewProc("K32EnumProcesses")
)

// GetRunningProcessesCount returns the number of currently running processes.
func GetRunningProcessesCount() (int, error) {
	var ids [1024]uint32
	var needed uint32
	r1, _, err := enumProcesses.Call(uintptr(unsafe.Pointer(&ids)), uintptr(len(ids)), uintptr(unsafe.Pointer(&needed)))
	if r1 == 0 {
		log.Printf("K32EnumProcesses failed: %v", err)
		return 0, nil
	}
	return int(needed / 4), nil
}

// CheckRunningProcessesCount checks if the number of currently running processes is less than a specified count.
func CheckRunningProcessesCount(count int) (bool, error) {
	processesCount, err := GetRunningProcessesCount()
	if err != nil {
		return false, err
	}

	if processesCount < count {
		return true, nil
	}
	return false, nil
}
