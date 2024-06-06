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
	"github.com/EvilBytecode/GoDefender/AntiDebug/UserAntiAntiDebug"

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
