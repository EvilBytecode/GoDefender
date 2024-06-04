package remotedebuggercheck

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

var (
	mk32  = syscall.NewLazyDLL("kernel32.dll")
	crdp  = mk32.NewProc("CheckRemoteDebuggerPresent")
)

// RemoteDebugger checks for the presence of a remote debugger.
func RemoteDebugger() {
	var isremdebpres bool
	crdp.Call(^uintptr(0), uintptr(unsafe.Pointer(&isremdebpres)))
	if isremdebpres {
		fmt.Println("Debug check: Remote debugger detected.")
		os.Exit(-1)
	} else {
		fmt.Println("Debug check: Remote debugger is not present.")
	}
}
