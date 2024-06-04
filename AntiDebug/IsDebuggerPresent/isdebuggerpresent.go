package IsDebuggerPresent

import (
    "os"
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

// CheckAndPrint checks if a debugger is present and prints a message.
func IsDebuggerPresent() {
    if IsDebuggerPresent1() {
        println("Debug check: IsDebuggerPresent is present.")
		os.Exit(-1)
    } else {
        println("Debug check: IsDebuggerPresent is not present.")
    }
}
