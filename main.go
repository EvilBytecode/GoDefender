package main

import (
	"fmt"
	"os"

	// Anti-Virtualization
	"github.com/EvilBytecode/GoDefender/AntiDebug/InternetCheck"
	hypervisor "github.com/EvilBytecode/GoDefender/AntiVirtualization/Hypervisor"
	kvmcheck "github.com/EvilBytecode/GoDefender/AntiVirtualization/KVMCheck"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/MonitorMetrics"
	triagecheck "github.com/EvilBytecode/GoDefender/AntiVirtualization/TriageDetection"
	usernamecheck "github.com/EvilBytecode/GoDefender/AntiVirtualization/UsernameCheck"
	vmcheck "github.com/EvilBytecode/GoDefender/AntiVirtualization/VMCheck"

	// Anti-Debug
	blacklistcheck "github.com/EvilBytecode/GoDefender/AntiDebug/CheckBlacklistedWindowsNames"
	"github.com/EvilBytecode/GoDefender/AntiDebug/IsDebuggerPresent"
	processkiller "github.com/EvilBytecode/GoDefender/AntiDebug/KillBadProcesses"
	parentantidebug "github.com/EvilBytecode/GoDefender/AntiDebug/ParentAntiDebug"
	remotedebuggercheck "github.com/EvilBytecode/GoDefender/AntiDebug/RemoteDebugger"
	runningprocesses "github.com/EvilBytecode/GoDefender/AntiDebug/RunningProcesses"
	userantiantidebug "github.com/EvilBytecode/GoDefender/AntiDebug/UserAntiAntiDebug"
	"github.com/EvilBytecode/GoDefender/AntiDebug/pcuptime"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/RecentFileActivity"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/USBCheck"
	//"github.com/EvilBytecode/GoDefender/AntiDebug/ParentAntiDebug"
	// Process Related
	//"github.com/EvilBytecode/GoDefender/Process/CriticalProcess"
)

func main() {
	/*
	   Anti-Debug
	   -----------
	   - IsDebuggerPresent
	   - RemoteDebugger
	   - PC Uptime Check
	   - Running Proccesses Count
	   - Check blacklisted windows
	   - KillBlacklisted Proceseses
	   - Parent AntiDebug
	*/
	RecentFileActivity.RecentFileActivityCheck()
	USBCheck.PluggedIn()
	userantiantidebug.AntiAntiDebug()
	IsDebuggerPresent.IsDebuggerPresent()
	remotedebuggercheck.RemoteDebugger()
	pcuptime.CheckUptime(1200)
	runningprocesses.CheckRunningProcessesCount(50)
	blacklistcheck.CheckBlacklistedWindows()
	parentantidebug.ParentAntiDebug()
	processkiller.KillProcesses()
	/*
	   Anti-Virtulization
	   ----------------
	   - Triage Check
	   - VM Gpu check
	   - Anti KVM
	   - Username Check
	   - Hypervisor CPUID bit check
	*/

	InternetCheck.CheckConnection()
	triagecheck.TriageCheckDebug()
	MonitorMetrics.IsScreenSmall()
	vmcheck.GraphicsCardCheck()
	fmt.Println("No VM GPUs present")
	if kvmcheck.CheckForKVM() {
		os.Exit(-1)
	}
	if hypervisor.CheckHypervisorBit() {
		fmt.Println("If you have the Hyper-V feature enabled on your system, please disable it.\nHow to: https://www.guidingtech.com/how-to-disable-hyper-v-windows/")
		os.Exit(-1)
	}
	usernamecheck.CheckForBlacklistedNames()
	fmt.Println("IF YOURE HERE YOU PASSED LOL")
	/*
	   EXTRA THINGS NOW:
	*/
	//programutils.SetDebugPrivilege() this is for devs who plan on continuing
	//programutils.SetProcessCritical() // this automatically gets the SeDebugPrivillige
	fmt.Scanln()
}
