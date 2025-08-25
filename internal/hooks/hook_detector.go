package hooks

import (
	"syscall"
	"unsafe"
	"GoDefender/internal/utils"
)

func DetectHooksOnCommonWinAPIFunctions(moduleName string, functions []string) bool {
	api := utils.NewWinAPI()
	libraries := []string{"kernel32.dll", "kernelbase.dll", "ntdll.dll", "user32.dll", "win32u.dll"}
	kernellibfunc := []string{"IsDebuggerPresent", "CheckRemoteDebuggerPresent", "GetThreadContext", "CloseHandle", "OutputDebugStringA", "GetTickCount", "SetHandleInformation"}
	ntdllfunc := []string{"NtQueryInformationProcess", "NtSetInformationThread", "NtClose", "NtGetContextThread", "NtQuerySystemInformation", "NtCreateFile", "NtCreateProcess", "NtCreateSection", "NtCreateThread", "NtYieldExecution", "NtCreateUserProcess"}
	user32func := []string{"FindWindowW", "FindWindowA", "FindWindowExW", "FindWindowExA", "GetForegroundWindow", "GetWindowTextLengthA", "GetWindowTextA", "BlockInput", "CreateWindowExW", "CreateWindowExA"}
	win32ufunc := []string{"NtUserBlockInput", "NtUserFindWindowEx", "NtUserQueryWindow", "NtUserGetForegroundWindow"}

	for _, library := range libraries {
		hModule := api.LowLevelGetModuleHandle(library)
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
					function := api.LowLevelGetProcAddress(hModule, winAPIFunction)
					var functionBytes [1]byte
					syscall.Syscall(uintptr(function), 1, uintptr(unsafe.Pointer(&functionBytes[0])), 0, 0)
					if functionBytes[0] == 0x90 || functionBytes[0] == 0xE9 {
						return true
					}
				}
			case "user32.dll":
				for _, winAPIFunction := range user32func {
					function := api.LowLevelGetProcAddress(hModule, winAPIFunction)
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
			hModule := api.LowLevelGetModuleHandle(moduleName)
			function := api.LowLevelGetProcAddress(hModule, winAPIFunction)
			var functionBytes [1]byte
			syscall.Syscall(uintptr(function), 1, uintptr(unsafe.Pointer(&functionBytes[0])), 0, 0)
			if functionBytes[0] == 0x90 || functionBytes[0] == 0xE9 {
				return true
			}
		}
	}
	return false
}

type HookDetector struct{}

func New() *HookDetector {
	return &HookDetector{}
}

func (h *HookDetector) AntiAntiDebug() bool {
	return DetectHooksOnCommonWinAPIFunctions("", nil)
}