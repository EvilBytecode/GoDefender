package IsDebuggerPresent

import (
	"syscall"
)

var (
	kernel32DLL = syscall.NewLazyDLL("kernel32.dll")
	isDebugger  = kernel32DLL.NewProc("IsDebuggerPresent")
)

// IsDebuggerPresent1 checks if a debugger is present.
func IsDebuggerPresent1() bool {
	flag, _, _ := isDebugger.Call()
	return flag != 0
}

// IsDebuggerPresent checks if a debugger is present and logs the result.
func IsDebuggerPresent() bool {
	if IsDebuggerPresent1() {
		return true
	} else {
		return false
	}
}
