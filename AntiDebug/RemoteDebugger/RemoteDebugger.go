package RemoteDebugger

import (
	"syscall"
	"unsafe"
)

var (
	mk32 = syscall.NewLazyDLL("kernel32.dll")
	crdp = mk32.NewProc("CheckRemoteDebuggerPresent")
)

// RemoteDebugger checks for the presence of a remote debugger.
func RemoteDebugger() (bool, error) {
	var isremdebpres bool
	r1, _, err := crdp.Call(^uintptr(0), uintptr(unsafe.Pointer(&isremdebpres)))
	if r1 == 0 {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	if isremdebpres {
		return true, nil
	} else {
		return false, nil
	}
}
