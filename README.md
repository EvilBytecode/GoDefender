# Go Defender

![Go Defender](GoDefender.png)

## GoDefender

This Go package provides functionality to detect and defend against various forms of debugging tools, virtualization environments.
btw for quick setup run install.bat
### Anti-Virtualization

- **Triage Detection**: Detects if the system is running in a triage or analysis environment.
- **Monitor Metrics**: Monitors system metrics to identify abnormal behavior indicative of virtualization.
- **VirtualBox Detection**: Detects the presence of Oracle VirtualBox.
- **VMWare Detection**: Detects the presence of VMware virtualization software.
- **KVM Check**: Checks for Kernel-based Virtual Machine (KVM) hypervisor.
- **Username Check**: Verifies if the current user is a default virtualization user.

### Anti-Debug

This module includes functions to detect and prevent debugging and analysis of the running process.

- **IsDebuggerPresent**: Checks if a debugger is currently attached to the process.
- **Remote Debugger**: Detects if a remote debugger is connected to the process.
- **PC Uptime**: Monitors system uptime to detect debugging attempts based on system restarts.
- **Check Blacklisted Windows Names**: Verifies if the process name matches any blacklisted names commonly used by debuggers.
- **Running Processes**: Retrieves a list of running processes and identifies potential malicious ones.
- **Parent Anti-Debug**: Detects if the parent process is attempting to debug the current process.
- **Kill Bad Processes**: Terminates known malicious processes detected on the system.

### Process

This module focuses on critical processes that should be monitored or protected.
- **Critical Process**: Implements functionality to manage critical processes essential for system operation.
- **SetDebugPrivilege**: Grants better permissions.

### Syntax:
```go
package main

import (
	"fmt"
	"os"
	// Anti-Virtualization
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/TriageDetection"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/MonitorMetrics"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/VirtualboxDetection"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/VMWareDetection"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/KVMCheck"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/UsernameCheck"

	// Anti-Debug
	"github.com/EvilBytecode/GoDefender/AntiDebug/IsDebuggerPresent"
	"github.com/EvilBytecode/GoDefender/AntiDebug/RemoteDebugger"
	"github.com/EvilBytecode/GoDefender/AntiDebug/pcuptime"
	"github.com/EvilBytecode/GoDefender/AntiDebug/CheckBlacklistedWindowsNames"
	"github.com/EvilBytecode/GoDefender/AntiDebug/RunningProcesses"
	"github.com/EvilBytecode/GoDefender/AntiDebug/ParentAntiDebug"
	"github.com/EvilBytecode/GoDefender/AntiDebug/KillBadProcesses"

	// Process Related
	"github.com/EvilBytecode/GoDefender/Process/CriticalProcess"
)

func main() {
	/* 
	ANTIDEBUG
	-----------
	- IsDebuggerPresent
	- RemoteDebugger
	- PC Uptime Check
	- Running Proccesses Count
	- Check blacklisted windows
	- KillBlacklisted Proceseses
	- Parent AntiDebug
	*/
        IsDebuggerPresent.IsDebuggerPresent()
	remotedebuggercheck.RemoteDebugger()
	pcuptime.CheckUptime(1200)
	runningprocesses.CheckRunningProcessesCount(50)
        blacklistcheck.CheckBlacklistedWindows()
	parentantidebug.ParentAntiDebug()
	processkiller.KillProcesses()

	/* 
	AntiVirulization
	----------------
	- Triage Check
	- VMWare Check
	- Anti KVM
	- Username Check
	- 
	*/
	triagecheck.TriageCheckDebug()
	MonitorMetrics.IsScreenSmall()
	VirtualboxDetection.GraphicsCardCheck()
	fmt.Println("Debug Check: VirtualBox isnt present")
	VMWare.GraphicsCardCheck()
	fmt.Println("Debug Check: VMWare isnt present")
	if kvmcheck.CheckForKVM() {
		os.Exit(-1)
	}
	usernamecheck.CheckForBlacklistedNames()
        artifactsdetector.BadVMFilesDetection()
	fmt.Println("IF YOURE HERE YOU PASSED LOL")
	/*
	EXTRA THINGS NOW:
	*/
	//programutils.SetDebugPrivilege() this is for devs who plan on continuing
	//programutils.SetProcessCritical() // this automatically gets the SeDebugPrivillige
	fmt.Scanln()
}
```

## Credits:
- https://github.com/AdvDebug = Inspired me to start making this package, without him it wouldnt be here, check out his github.
- https://github.com/MmCopyMemory = Giving ideas and many more, check out his github.
