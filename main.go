package main

import (
	"log"

	// AntiDebug
	"github.com/EvilBytecode/GoDefender/AntiDebug/CheckBlacklistedWindowsNames"
	"github.com/EvilBytecode/GoDefender/AntiDebug/InternetCheck"
	"github.com/EvilBytecode/GoDefender/AntiDebug/IsDebuggerPresent"
	"github.com/EvilBytecode/GoDefender/AntiDebug/KillBadProcesses"
	"github.com/EvilBytecode/GoDefender/AntiDebug/ParentAntiDebug"
	"github.com/EvilBytecode/GoDefender/AntiDebug/RunningProcesses"
	"github.com/EvilBytecode/GoDefender/AntiDebug/RemoteDebugger"
	"github.com/EvilBytecode/GoDefender/AntiDebug/pcuptime"

	// AntiVirtualization
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/KVMCheck"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/MonitorMetrics"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/RecentFileActivity"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/TriageDetection"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/UsernameCheck"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/VirtualboxDetection"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/VMWareDetection"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/USBCheck"
	
	// ProcessRelatedTool
	//"github.com/EvilBytecode/GoDefender/Process/CriticalProcess"
)

func main() {
	// AntiDebug checks
	if connected, _ := InternetCheck.CheckConnection(); connected {
		log.Println("[DEBUG] Internet connection is present")
	} else {
		log.Println("[DEBUG] Internet connection isn't present")
	}

	if parentAntiDebugResult := ParentAntiDebug.ParentAntiDebug(); parentAntiDebugResult {
		log.Println("[DEBUG] ParentAntiDebug check failed")
	} else {
		log.Println("[DEBUG] ParentAntiDebug check passed")
	}

	if runningProcessesCountDetected, _ := RunningProcesses.CheckRunningProcessesCount(50); runningProcessesCountDetected {
		log.Println("[DEBUG] Running processes count detected")
	} else {
		log.Println("[DEBUG] Running processes count passed")
	}

	if pcUptimeDetected, _ := pcuptime.CheckUptime(1200); pcUptimeDetected {
		log.Println("[DEBUG] PC uptime detected")
	} else {
		log.Println("[DEBUG] PC uptime passed")
	}

	KillBadProcesses.KillProcesses()
	CheckBlacklistedWindowsNames.CheckBlacklistedWindows()
	// Other AntiDebug checks
	if isDebuggerPresentResult := IsDebuggerPresent.IsDebuggerPresent1(); isDebuggerPresentResult {
		log.Println("[DEBUG] Debugger presence detected")
	} else {
		log.Println("[DEBUG] Debugger presence passed")
	}

	if remoteDebuggerDetected, _ := RemoteDebugger.RemoteDebugger(); remoteDebuggerDetected {
		log.Println("[DEBUG] Remote debugger detected")
	} else {
		log.Println("[DEBUG] Remote debugger passed")
	}
	//////////////////////////////////////////////////////

	// AntiVirtualization checks
	if recentFileActivityDetected, _ := RecentFileActivity.RecentFileActivityCheck(); recentFileActivityDetected {
		log.Println("[DEBUG] Recent file activity detected")
	} else {
		log.Println("[DEBUG] Recent file activity passed")
	}

	if vmwareDetected, _ := VMWareDetection.GraphicsCardCheck(); vmwareDetected {
		log.Println("[DEBUG] VMWare detected")
	} else {
		log.Println("[DEBUG] VMWare passed")
	}

	if virtualboxDetected, _ := VirtualboxDetection.GraphicsCardCheck(); virtualboxDetected {
		log.Println("[DEBUG] Virtualbox detected")
	} else {
		log.Println("[DEBUG] Virtualbox passed")
	}

	if kvmDetected, _ := KVMCheck.CheckForKVM(); kvmDetected {
		log.Println("[DEBUG] KVM detected")
	} else {
		log.Println("[DEBUG] KVM passed")
	}

	if blacklistedUsernameDetected := UsernameCheck.CheckForBlacklistedNames(); blacklistedUsernameDetected {
		log.Println("[DEBUG] Blacklisted username detected")
	} else {
		log.Println("[DEBUG] Blacklisted username passed")
	}

	if triageDetected, _ := TriageDetection.TriageCheck(); triageDetected {
		log.Println("[DEBUG] Triage detected")
	} else {
		log.Println("[DEBUG] Triage passed")
	}
	if isScreenSmall, _ := MonitorMetrics.IsScreenSmall(); isScreenSmall {
		log.Println("[DEBUG] Screen size is small")
	} else {
		log.Println("[DEBUG] Screen size is not small")
	}
	// USBCheck
	if usbPluggedIn, err := USBCheck.PluggedIn(); err != nil {
			log.Println("[DEBUG] Error checking USB devices:", err)
	} else if usbPluggedIn {
			log.Println("[DEBUG] USB devices have been plugged in, check passed.")
	} else {
			log.Println("[DEBUG] No USB devices detected")
	}
	
   // updating this once v1.0.9 is out!
}
