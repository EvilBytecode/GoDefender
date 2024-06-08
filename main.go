package main

import (
	"log"

	// AntiDebug
	"GoDefenderREWRITE/AntiDebug/CheckBlacklistedWindowsNames"
	"GoDefenderREWRITE/AntiDebug/InternetCheck"
	"GoDefenderREWRITE/AntiDebug/IsDebuggerPresent"
	"GoDefenderREWRITE/AntiDebug/KillBadProcesses"
	"GoDefenderREWRITE/AntiDebug/ParentAntiDebug"
	"GoDefenderREWRITE/AntiDebug/RunningProcesses"
	"GoDefenderREWRITE/AntiDebug/RemoteDebugger"
	"GoDefenderREWRITE/AntiDebug/pcuptime"

	// AntiVirtualization
	"GoDefenderREWRITE/AntiVirtualization/KVMCheck"
	"GoDefenderREWRITE/AntiVirtualization/MonitorMetrics"
	"GoDefenderREWRITE/AntiVirtualization/RecentFileActivity"
	"GoDefenderREWRITE/AntiVirtualization/TriageDetection"
	"GoDefenderREWRITE/AntiVirtualization/UsernameCheck"
	"GoDefenderREWRITE/AntiVirtualization/VirtualboxDetection"
	"GoDefenderREWRITE/AntiVirtualization/VMWareDetection"
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

	// Continue with other checks... (you can add ones related to critical process or sedebugprivvilege)
}
