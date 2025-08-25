package antidll

import (
	"GoDefender/internal/utils"
	"unsafe"
)

type DLLProtector struct {
	winapi *utils.WinAPI
}

type PROCESS_MITIGATION_BINARY_SIGNATURE_POLICY struct {
	MicrosoftSignedOnly uint32
}

const (
	ProcessSignaturePolicyMitigation = 8
)

func New() *DLLProtector {
	return &DLLProtector{
		winapi: utils.NewWinAPI(),
	}
}

func (d *DLLProtector) PreventDLLInjection() error {
	var onlyMicrosoftBinaries PROCESS_MITIGATION_BINARY_SIGNATURE_POLICY
	onlyMicrosoftBinaries.MicrosoftSignedOnly = 1

	kernelbase := d.winapi.GetModuleHandle("kernelbase.dll")
	if kernelbase == 0 {
		return d.winapi.LastError()
	}

	procSetProcessMitigationPolicy := d.winapi.GetProcAddress(kernelbase, "SetProcessMitigationPolicy")
	if procSetProcessMitigationPolicy == 0 {
		return d.winapi.LastError()
	}

	ret, _, err := d.winapi.CallProc(
		procSetProcessMitigationPolicy,
		uintptr(ProcessSignaturePolicyMitigation),
		uintptr(unsafe.Pointer(&onlyMicrosoftBinaries)),
		uintptr(unsafe.Sizeof(onlyMicrosoftBinaries)),
	)

	if ret == 0 {
		return err
	}

	return nil
}

func (d *DLLProtector) PatchAllLoadLibrary() error {
	kernelbase := d.winapi.GetModuleHandle("kernelbase.dll")
	ntdll := d.winapi.GetModuleHandle("ntdll.dll")
	
	if kernelbase == 0 {
		return d.winapi.LastError()
	}
    // remember go uses loadlib, so we patch it but issue is the program will crash (access violation, so thats why i put it as last lol)
	kernelbaseFunctions := []string{"LoadLibraryA", "LoadLibraryW", "LoadLibraryExA", "LoadLibraryExW"}
	ntdllFunctions := []string{"LdrLoadDll"}
	/*
	you might be confused why its not 0xc3 (ret)
	ret imm16 (C2 xx xx) pops the return address, then additionally
	adds imm16 bytes to ESP. This is commonly used in stdcall functions
	to clean up arguments. Here, 'C2 04 00' means 'ret 4', which pops
	the return address AND cleans up 4 bytes (one argument).
	It's not just "ret + N"; it's specifically callee stack cleanup.
	*/
	// i saw this back then in a advdebugs code so i just recoded it, but he removed it i think...
	// but anyways i opened x64, ctrl+g (Function we want to patch) and yes if you check start bytes you can check that its been patched.
	hookedCode := []byte{0xC2, 0x04, 0x00}
	for _, funcName := range kernelbaseFunctions {
		funcAddr := d.winapi.GetProcAddress(kernelbase, funcName)
		if funcAddr == 0 {
			continue 
		}
		success := d.winapi.WriteProcessMemory(funcAddr, hookedCode)
		if !success {
			//utils.Print("Failed to patch %s", funcName)
		}
	}

	if ntdll != 0 {
		for _, funcName := range ntdllFunctions {
			funcAddr := d.winapi.GetProcAddress(ntdll, funcName)
			if funcAddr == 0 {
				//utils.Print("Function %s not found in ntdll.dll", funcName)
				continue 
			}

			//utils.Print("Patching %s at address 0x%x", funcName, funcAddr)
			success := d.winapi.WriteProcessMemory(funcAddr, hookedCode)
			if !success {
				//utils.Print("Failed to patch %s", funcName)
			} else {
				//utils.Print("Successfully patched %s", funcName)
			}
		}
	}

	return nil
}