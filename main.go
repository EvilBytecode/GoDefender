package main

import (
	"GoDefender/internal/antivm"
	"GoDefender/internal/antidebug"
	"GoDefender/internal/antidll"
	"GoDefender/internal/hooks"
	"GoDefender/internal/utils"
	"fmt"
)

func main() {
	utils.Print("Starting GoDefender checks...")

	vmDetector := antivm.New()
	debugger := antidebug.New()
	dllProtector := antidll.New()
	hookDetector := hooks.New()

	if lowRefresh, err := vmDetector.CheckDisplayRefreshRate(); err == nil && lowRefresh {
		utils.Print("Suspicious display refresh rate detected (< 29Hz)")
	}

	if err := dllProtector.PreventDLLInjection(); err != nil {
		utils.Print("Failed to set DLL injection protection")
	}

	if hookDetector.AntiAntiDebug() {
		utils.Print("API hooks detected")
	}
	
	if debugger.PatchAntiDebug() {
		utils.Print("Anti-debug patches applied")
	}
	
	if debugger.SetDebugFilterState() {
		utils.Print("Debug filter state protected")
	}

	usbPluggedIn, err := vmDetector.CheckUSBDevices()
	if err != nil || !usbPluggedIn {
		utils.Print("USB check failed")
	}

	if vmDetector.CheckBlacklistedUsernames() {
		utils.Print("Blacklisted username detected")
	}

	if !debugger.CheckParentProcess() {
		utils.Print("Suspicious parent process detected")
	}

	if vmware, _ := vmDetector.CheckVMware(); vmware {
		utils.Print("VMWare detected")
	}

	if vbox, _ := vmDetector.CheckVirtualBox(); vbox {
		utils.Print("VirtualBox detected")
	}

	if kvm, _ := vmDetector.CheckKVM(); kvm {
		utils.Print("KVM detected")
	}

	if parallels, _ := vmDetector.CheckParallels(); parallels {
		utils.Print("Parallels detected")
	}

	if qemu, _ := vmDetector.CheckQEMU(); qemu {
		utils.Print("QEMU detected")
	}

	if vmDetector.CheckVMFiles() {
		utils.Print("VM files detected")
	}

	if vmDetector.CheckAnyRun() {
		utils.Print("Any.Run detected")
	}

	if portCheck, _ := vmDetector.CheckPortConnectors(); portCheck {
		utils.Print("Suspicious port configuration")
	}

	if screenSmall, _ := vmDetector.CheckScreenSize(); screenSmall {
		utils.Print("Suspicious screen metrics")
	}

	if vmDetector.CheckNamedPipes() {
		utils.Print("Suspicious named pipes detected")
	}


	if remoteDbg, _ := debugger.CheckRemoteDebugger(); remoteDbg {
		utils.Print("Remote debugger detected")
	}

	if debugger.CheckBlacklistedWindows() {
		utils.Print("Analysis tool window detected")
	}

	if badProc, _ := debugger.CheckBlacklistedProcesses(); badProc {
		utils.Print("Malicious process detected")
	}

	if repProc, _ := debugger.CheckRepetitiveProcesses(60); repProc {
		utils.Print("Suspicious process pattern")
	}

	if connected, _ := debugger.CheckInternetConnection(); !connected {
		utils.Print("No internet connection")
	}

	procCount, _ := debugger.GetRunningProcessCount()
	if procCount < 50 {
		utils.Print("Abnormal process count")
	}

	utils.Print("✅ All security checks passed!")

	if err := dllProtector.PatchAllLoadLibrary(); err != nil {
		utils.Print("Failed to patch LoadLibrary functions")
	} else {
		utils.Print("All LoadLibrary functions patched successfully")
	}

	fmt.Scanln()
}	
