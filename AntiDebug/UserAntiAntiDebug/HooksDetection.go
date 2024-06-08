package HooksDetection

import (
	"log"
	"syscall"
	"unsafe"
)

type UNICODE_STRING struct {
	Length        uint16
	MaximumLength uint16
	Buffer        uintptr
}

type ANSI_STRING struct {
	Length        int16
	MaximumLength int16
	Buffer        *byte
}
// creds to advdebug ported from cs to go
var (
	modNtdll      = syscall.NewLazyDLL("ntdll.dll")
	modKernelbase = syscall.NewLazyDLL("kernelbase.dll")
	modKernel32   = syscall.NewLazyDLL("kernel32.dll")
	modUser32     = syscall.NewLazyDLL("user32.dll")
	modWin32u     = syscall.NewLazyDLL("win32u.dll")

	procRtlInitUnicodeString          = modNtdll.NewProc("RtlInitUnicodeString")
	procRtlUnicodeStringToAnsiString  = modNtdll.NewProc("RtlUnicodeStringToAnsiString")
	procLdrGetDllHandleEx             = modNtdll.NewProc("LdrGetDllHandleEx")
	procGetModuleHandleA              = modKernelbase.NewProc("GetModuleHandleA")
	procGetProcAddress                = modKernelbase.NewProc("GetProcAddress")
	procLdrGetProcedureAddressForCall = modNtdll.NewProc("LdrGetProcedureAddressForCaller")
)

func RtlInitUnicodeString(destinationString *UNICODE_STRING, sourceString string) {
	sourcePtr, _ := syscall.UTF16PtrFromString(sourceString)
	syscall.Syscall(procRtlInitUnicodeString.Addr(), 2, uintptr(unsafe.Pointer(destinationString)), uintptr(unsafe.Pointer(sourcePtr)), 0)
}

func RtlUnicodeStringToAnsiString(destinationString *ANSI_STRING, unicodeString *UNICODE_STRING, allocateDestinationString bool) {
	syscall.Syscall(procRtlUnicodeStringToAnsiString.Addr(), 3, uintptr(unsafe.Pointer(destinationString)), uintptr(unsafe.Pointer(unicodeString)), uintptr(boolToInt(allocateDestinationString)))
}

func GetModuleHandleA(library string) uintptr {
	ret, _, _ := procGetModuleHandleA.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(library))))
	return ret
}

func GetProcAddress(hModule uintptr, function string) uintptr {
	ret, _, _ := procGetProcAddress.Call(hModule, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(function))))
	return ret
}

func LdrGetProcedureAddressForCaller(moduleHandle uintptr, procedureName *ANSI_STRING, procedureNumber uint16, functionHandle *uintptr, flags uint64, callback uintptr) uint32 {
	ret, _, _ := procLdrGetProcedureAddressForCall.Call(moduleHandle, uintptr(unsafe.Pointer(procedureName)), uintptr(procedureNumber), uintptr(unsafe.Pointer(functionHandle)), uintptr(flags), callback)
	return uint32(ret)
}

func LdrGetDllHandleEx(flags uint64, dllPath string, dllCharacteristics string, libraryName *UNICODE_STRING, dllHandle *uintptr) uint32 {
	ret, _, _ := procLdrGetDllHandleEx.Call(uintptr(flags), 0, 0, uintptr(unsafe.Pointer(libraryName)), uintptr(unsafe.Pointer(dllHandle)))
	return uint32(ret)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func LowLevelGetModuleHandle(library string) uintptr {
	var hModule uintptr
	var unicodeString UNICODE_STRING
	RtlInitUnicodeString(&unicodeString, library)
	LdrGetDllHandleEx(0, "", "", &unicodeString, &hModule)
	return hModule
}

func LowLevelGetProcAddress(hModule uintptr, function string) uintptr {
	var functionHandle uintptr
	var unicodeString UNICODE_STRING
	var ansiString ANSI_STRING
	RtlInitUnicodeString(&unicodeString, function)
	RtlUnicodeStringToAnsiString(&ansiString, &unicodeString, true)
	LdrGetProcedureAddressForCaller(hModule, &ansiString, 0, &functionHandle, 0, 0)
	return functionHandle
}

func DetectHooksOnCommonWinAPIFunctions(moduleName string, functions []string) bool {
	libraries := []string{"kernel32.dll", "kernelbase.dll", "ntdll.dll", "user32.dll", "win32u.dll"}
	kernellibfunc := []string{"IsDebuggerPresent", "CheckRemoteDebuggerPresent", "GetThreadContext", "CloseHandle", "OutputDebugStringA", "GetTickCount", "SetHandleInformation"}
	ntdllfunc := []string{"NtQueryInformationProcess", "NtSetInformationThread", "NtClose", "NtGetContextThread", "NtQuerySystemInformation", "NtCreateFile", "NtCreateProcess", "NtCreateSection", "NtCreateThread", "NtYieldExecution", "NtCreateUserProcess"}
	user32func := []string{"FindWindowW", "FindWindowA", "FindWindowExW", "FindWindowExA", "GetForegroundWindow", "GetWindowTextLengthA", "GetWindowTextA", "BlockInput", "CreateWindowExW", "CreateWindowExA"}
	win32ufunc := []string{"NtUserBlockInput", "NtUserFindWindowEx", "NtUserQueryWindow", "NtUserGetForegroundWindow"}

	for _, library := range libraries {
		hModule := LowLevelGetModuleHandle(library)
		if hModule != 0 {
			switch library {
			case "kernel32.dll", "kernelbase.dll", "ntdll.dll", "win32u.dll":
				var commonFunctions []string
				switch library {
				case "kernel32.dll", "kernelbase.dll":
					commonFunctions = kernellibfunc
				case "ntdll.dll":
					commonFunctions = ntdllfunc
				case "win32u.dll":
					commonFunctions = win32ufunc
				}
				for _, winAPIFunction := range commonFunctions {
					function := LowLevelGetProcAddress(hModule, winAPIFunction)
					var functionBytes [1]byte
					syscall.Syscall(uintptr(function), 1, uintptr(unsafe.Pointer(&functionBytes[0])), 0, 0)
					if functionBytes[0] == 0x90 || functionBytes[0] == 0xE9 {
						return true
					}
				}
			case "user32.dll":
				for _, winAPIFunction := range user32func {
					function := LowLevelGetProcAddress(hModule, winAPIFunction)
					var functionBytes [1]byte
					syscall.Syscall(uintptr(function), 1, uintptr(unsafe.Pointer(&functionBytes[0])), 0, 0)
					if functionBytes[0] == 0x90 || functionBytes[0] == 0xE9 {
						return true
					}
				}
			}
		}
	}

	if moduleName != "" && functions != nil {
		for _, winAPIFunction := range functions {
			hModule := LowLevelGetModuleHandle(moduleName)
			function := LowLevelGetProcAddress(hModule, winAPIFunction)
			var functionBytes [1]byte
			syscall.Syscall(uintptr(function), 1, uintptr(unsafe.Pointer(&functionBytes[0])), 0, 0)
			if functionBytes[0] == 0x90 || functionBytes[0] == 0xE9 {
				return true
			}
		}
	}
	return false
}

func AntiAntiDebug() {
	log.Println("Detecting Hooks on Common WinAPI Functions by checking for Bad Instructions on Functions Addresses (Most Effective on x64): ", DetectHooksOnCommonWinAPIFunctions("", nil))
}
