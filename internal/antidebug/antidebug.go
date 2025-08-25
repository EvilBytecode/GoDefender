package antidebug

import (
	"GoDefender/internal/utils"
	"net"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
	"golang.org/x/sys/windows"
)

type Debugger struct {
	Winapi               *utils.WinAPI
	blacklistedProcesses []string
	blacklistedWindows   []string
}

type ProcessInfo struct {
	Res1             uintptr
	PebAddr          uintptr
	Res2             [2]uintptr
	PID              uintptr
	InheritedFromPID uintptr
}

func New() *Debugger {
	return &Debugger{
		Winapi: utils.NewWinAPI(),
		blacklistedProcesses: []string{
			"taskmgr.exe", "process.exe", "processhacker.exe", "ksdumper.exe", "fiddler.exe",
			"httpdebuggerui.exe", "wireshark.exe", "httpanalyzerv7.exe", "decoder.exe",
			"regedit.exe", "procexp.exe", "dnspy.exe", "vboxservice.exe", "burpsuite.exe",
			"DbgX.Shell.exe", "ILSpy.exe", "ollydbg.exe", "x32dbg.exe", "x64dbg.exe", "gdb.exe",
			"idaq.exe", "idag.exe", "idaw.exe", "ida64.exe", "idag64.exe", "idaw64.exe",
			"idaq64.exe", "windbg.exe", "immunitydebugger.exe", "windasm.exe",
		},
		blacklistedWindows: []string{
			"proxifier", "graywolf", "extremedumper", "zed", "exeinfope", "dnspy",
			"titanHide", "ilspy", "titanhide", "x32dbg", "codecracker", "simpleassembly",
			"process hacker 2", "pc-ret", "http debugger", "debug", "ILSpy",
			"simpleassemblyexplorer", "process", "de4dotmodded", "pizza", "crack",
			"strongod", "ida -", "brute", "dump", "StringDecryptor", "wireshark",
			"debugger", "httpdebugger", "gdb", "windbg", "x64_dbg", "x64netdumper",
			"ollydbg", "immunitydebugger",
		},
	}
}

func (d *Debugger) PatchAntiDebug() bool {
	ntdllModule := d.Winapi.GetModuleHandle("ntdll.dll")
	if ntdllModule == 0 {
		return false
	}

	dbgUiRemoteBreakinAddr := d.Winapi.GetProcAddress(ntdllModule, "DbgUiRemoteBreakin")
	dbgBreakPointAddr := d.Winapi.GetProcAddress(ntdllModule, "DbgBreakPoint")

	if dbgUiRemoteBreakinAddr == 0 || dbgBreakPointAddr == 0 {
		return false
	}

	int3InvalidCode := []byte{0xCC}
	retCode := []byte{0xC3}

	status1 := d.Winapi.WriteProcessMemory(dbgUiRemoteBreakinAddr, int3InvalidCode)
	status2 := d.Winapi.WriteProcessMemory(dbgBreakPointAddr, retCode)

	return status1 && status2
}

func (d *Debugger) SetDebugFilterState() bool {
	return d.Winapi.SetDebugFilterState(0, 0, true)
}

func (d *Debugger) CheckRemoteDebugger() (bool, error) {
	return d.Winapi.CheckRemoteDebugger()
}

func (d *Debugger) GetRunningProcessCount() (int, error) {
	return d.Winapi.GetRunningProcessCount()
}

func (d *Debugger) CheckInternetConnection() (bool, error) {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		return false, err
	}
	defer conn.Close()
	return true, nil
}

func (d *Debugger) CheckBlacklistedProcesses() (bool, error) {
	processNames, err := d.Winapi.GetRunningProcessNames()
	if err != nil {
		return false, err
	}

	for _, processName := range processNames {
		processNameLower := strings.ToLower(processName)
		for _, badProcess := range d.blacklistedProcesses {
			if processNameLower == strings.ToLower(badProcess) {
				return true, nil
			}
		}
	}

	return false, nil
}

func (d *Debugger) CheckRepetitiveProcesses(threshold int) (bool, error) {
	processNames, err := d.Winapi.GetRunningProcessNames()
	if err != nil {
		return false, err
	}

	processCounts := make(map[string]int)
	for _, processName := range processNames {
		processName = strings.ToLower(processName)
		if processName != "svchost.exe" {
			processCounts[processName]++
		}
	}

	for _, count := range processCounts {
		if count > threshold {
			return true, nil
		}
	}

	return false, nil
}

func (d *Debugger) CheckParentProcess() bool {
	const ProcInfo = 0
	var p ProcessInfo
	
	handle := syscall.Handle(windows.CurrentProcess())
	
	r1, _, err := d.Winapi.QueryInformationProcess(
		handle,
		ProcInfo,
		uintptr(unsafe.Pointer(&p)),
		uint32(unsafe.Sizeof(p)),
	)
	
	if r1 != 0 || err != nil && err != syscall.Errno(0) {
		return false
	}
	
	parentPID := int32(p.InheritedFromPID)
	if parentPID == 0 {
		return false
	}
	
	parentHandle, err := syscall.OpenProcess(syscall.PROCESS_QUERY_INFORMATION, false, uint32(parentPID))
	if err != nil {
		return false
	}
	defer syscall.CloseHandle(parentHandle)
	
	var nameBuffer [windows.MAX_PATH]uint16
	size := uint32(len(nameBuffer))
	err = windows.QueryFullProcessImageName(windows.Handle(parentHandle), 0, &nameBuffer[0], &size)
	if err != nil {
		return false
	}
	
	parentName := filepath.Base(syscall.UTF16ToString(nameBuffer[:size]))
	return parentName == "explorer.exe" || parentName == "cmd.exe"
}

func (d *Debugger) CheckBlacklistedWindows() bool {
	user32 := windows.NewLazySystemDLL("user32.dll")
	procGetWindowText := user32.NewProc("GetWindowTextW")
	procEnumWindows := user32.NewProc("EnumWindows")
    found = false
	var enumWindowsProc = func(hwnd windows.HWND, lparam uintptr) uintptr {
		var title [256]uint16
		procGetWindowText.Call(
			uintptr(hwnd),
			uintptr(unsafe.Pointer(&title[0])),
			uintptr(len(title)),
		)
		windowTitle := syscall.UTF16ToString(title[:])
		
		for _, blacklisted := range d.blacklistedWindows {
			if strings.Contains(strings.ToLower(windowTitle), strings.ToLower(blacklisted)) {
				found = true
				return 0
			}
		}
		return 1
	}

	procEnumWindows.Call(
		windows.NewCallback(enumWindowsProc),
		0,
	)
	return found
}
