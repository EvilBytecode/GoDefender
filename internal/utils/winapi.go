package utils

import (
	"syscall"
	"unsafe"
	"golang.org/x/sys/windows"
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

type DEVMODE struct {
	DmDeviceName       [32]uint16
	DmSpecVersion      uint16
	DmDriverVersion    uint16
	DmSize             uint16
	DmDriverExtra      uint16
	DmFields           uint32
	DmOrientation      int16
	DmPaperSize        int16
	DmPaperLength      int16
	DmPaperWidth       int16
	DmScale            int16
	DmCopies           int16
	DmDefaultSource    int16
	DmPrintQuality     int16
	DmColor           int16
	DmDuplex          int16
	DmYResolution     int16
	DmTTOption        int16
	DmCollate         int16
	DmFormName        [32]uint16
	DmLogPixels       uint16
	DmBitsPerPel      uint32
	DmPelsWidth       uint32
	DmPelsHeight      uint32
	DmDisplayFlags    uint32
	DmDisplayFrequency uint32
	DmICMMethod       uint32
	DmICMIntent       uint32
	DmMediaType       uint32
	DmDitherType      uint32
	DmReserved1       uint32
	DmReserved2       uint32
	DmPanningWidth    uint32
	DmPanningHeight   uint32
}

type WinAPI struct {
	Kernel32              *syscall.LazyDLL
	Ntdll                 *syscall.LazyDLL
	Psapi                 *syscall.LazyDLL
	User32                *syscall.LazyDLL
	Kernelbase            *syscall.LazyDLL
	Win32u                *syscall.LazyDLL

	ProcGetModuleHandle   *syscall.LazyProc
	ProcGetProcAddress    *syscall.LazyProc
	ProcEnumProcesses     *syscall.LazyProc
	ProcOpenProcess       *syscall.LazyProc
	ProcCreateFile        *syscall.LazyProc
	ProcCloseHandle       *syscall.LazyProc
	ProcGetModuleBaseName *syscall.LazyProc
	ProcWriteProcessMemory *syscall.LazyProc
	ProcGetCurrentProcess  *syscall.LazyProc

	ProcIsDebuggerPresent  *syscall.LazyProc
	ProcCheckRemoteDebugger *syscall.LazyProc
	ProcNtSetDebugFilterState *syscall.LazyProc
	ProcNtQueryInformationProcess *syscall.LazyProc

	ProcRtlInitUnicodeString         *syscall.LazyProc
	ProcRtlUnicodeStringToAnsiString *syscall.LazyProc
	ProcLdrGetDllHandleEx            *syscall.LazyProc
	ProcLdrGetProcedureAddressForCall *syscall.LazyProc
	ProcEnumDisplaySettingsW          *syscall.LazyProc
	ProcSetLastError                  *syscall.LazyProc
	ProcGetLastError                  *syscall.LazyProc
	ProcOutputDebugString            *syscall.LazyProc
	ProcFopen                       *syscall.LazyProc
	ProcFclose                      *syscall.LazyProc
	ProcVirtualProtect              *syscall.LazyProc
	ProcVirtualFree                 *syscall.LazyProc
}

func NewWinAPI() *WinAPI {
	w := &WinAPI{
		Kernel32:    syscall.NewLazyDLL("kernel32.dll"),
		Ntdll:       syscall.NewLazyDLL("ntdll.dll"),
		Psapi:       syscall.NewLazyDLL("psapi.dll"),
		User32:      syscall.NewLazyDLL("user32.dll"),
		Kernelbase:  syscall.NewLazyDLL("kernelbase.dll"),
		Win32u:      syscall.NewLazyDLL("win32u.dll"),
	}

	w.ProcGetModuleHandle = w.Kernel32.NewProc("GetModuleHandleW")
	w.ProcGetProcAddress = w.Kernel32.NewProc("GetProcAddress")
	w.ProcEnumProcesses = w.Psapi.NewProc("EnumProcesses")
	w.ProcOpenProcess = w.Kernel32.NewProc("OpenProcess")
	w.ProcCreateFile = w.Kernel32.NewProc("CreateFileW")
	w.ProcCloseHandle = w.Kernel32.NewProc("CloseHandle")
	w.ProcGetModuleBaseName = w.Psapi.NewProc("GetModuleBaseNameW")
	w.ProcWriteProcessMemory = w.Kernel32.NewProc("WriteProcessMemory")
	w.ProcGetCurrentProcess = w.Kernel32.NewProc("GetCurrentProcess")

	w.ProcIsDebuggerPresent = w.Kernel32.NewProc("IsDebuggerPresent")
	w.ProcCheckRemoteDebugger = w.Kernel32.NewProc("CheckRemoteDebuggerPresent")
	w.ProcNtSetDebugFilterState = w.Ntdll.NewProc("NtSetDebugFilterState")
	w.ProcNtQueryInformationProcess = w.Ntdll.NewProc("NtQueryInformationProcess")

	w.ProcRtlInitUnicodeString = w.Ntdll.NewProc("RtlInitUnicodeString")
	w.ProcRtlUnicodeStringToAnsiString = w.Ntdll.NewProc("RtlUnicodeStringToAnsiString")
	w.ProcLdrGetDllHandleEx = w.Ntdll.NewProc("LdrGetDllHandleEx")
	w.ProcLdrGetProcedureAddressForCall = w.Ntdll.NewProc("LdrGetProcedureAddressForCaller")
	w.ProcEnumDisplaySettingsW = w.User32.NewProc("EnumDisplaySettingsW")
	w.ProcSetLastError = w.Kernel32.NewProc("SetLastError")
	w.ProcGetLastError = w.Kernel32.NewProc("GetLastError")
	w.ProcOutputDebugString = w.Kernel32.NewProc("OutputDebugStringW")

	w.ProcFopen = syscall.NewLazyDLL("ucrtbase.dll").NewProc("fopen")
	w.ProcFclose = syscall.NewLazyDLL("ucrtbase.dll").NewProc("fclose")

	w.ProcVirtualProtect = w.Kernel32.NewProc("VirtualProtect")
	w.ProcVirtualFree = w.Kernel32.NewProc("VirtualFree")

	return w
}

func (w *WinAPI) RtlInitUnicodeString(destinationString *UNICODE_STRING, sourceString string) {
	sourcePtr, _ := syscall.UTF16PtrFromString(sourceString)
	syscall.Syscall(w.ProcRtlInitUnicodeString.Addr(), 2, uintptr(unsafe.Pointer(destinationString)), uintptr(unsafe.Pointer(sourcePtr)), 0)
}

func (w *WinAPI) RtlUnicodeStringToAnsiString(destinationString *ANSI_STRING, unicodeString *UNICODE_STRING, allocateDestinationString bool) {
	syscall.Syscall(w.ProcRtlUnicodeStringToAnsiString.Addr(), 3, uintptr(unsafe.Pointer(destinationString)), uintptr(unsafe.Pointer(unicodeString)), uintptr(boolToInt(allocateDestinationString)))
}

func (w *WinAPI) GetModuleHandle(dllName string) uintptr {
	dllNamePtr, _ := syscall.UTF16PtrFromString(dllName)
	handle, _, _ := w.ProcGetModuleHandle.Call(uintptr(unsafe.Pointer(dllNamePtr)))
	return handle
}

func (w *WinAPI) GetProcAddress(moduleHandle uintptr, procName string) uintptr {
	procNamePtr, _ := syscall.BytePtrFromString(procName)
	addr, _, _ := w.ProcGetProcAddress.Call(moduleHandle, uintptr(unsafe.Pointer(procNamePtr)))
	return addr
}

func (w *WinAPI) LdrGetProcedureAddressForCaller(moduleHandle uintptr, procedureName *ANSI_STRING, procedureNumber uint16, functionHandle *uintptr, flags uint64, callback uintptr) uint32 {
	ret, _, _ := w.ProcLdrGetProcedureAddressForCall.Call(moduleHandle, uintptr(unsafe.Pointer(procedureName)), uintptr(procedureNumber), uintptr(unsafe.Pointer(functionHandle)), uintptr(flags), callback)
	return uint32(ret)
}

func (w *WinAPI) LdrGetDllHandleEx(flags uint64, dllPath string, dllCharacteristics string, libraryName *UNICODE_STRING, dllHandle *uintptr) uint32 {
	ret, _, _ := w.ProcLdrGetDllHandleEx.Call(uintptr(flags), 0, 0, uintptr(unsafe.Pointer(libraryName)), uintptr(unsafe.Pointer(dllHandle)))
	return uint32(ret)
}

func (w *WinAPI) LowLevelGetModuleHandle(library string) uintptr {
	var hModule uintptr
	var unicodeString UNICODE_STRING
	w.RtlInitUnicodeString(&unicodeString, library)
	w.LdrGetDllHandleEx(0, "", "", &unicodeString, &hModule)
	return hModule
}

func (w *WinAPI) LowLevelGetProcAddress(hModule uintptr, function string) uintptr {
	var functionHandle uintptr
	var unicodeString UNICODE_STRING
	var ansiString ANSI_STRING
	w.RtlInitUnicodeString(&unicodeString, function)
	w.RtlUnicodeStringToAnsiString(&ansiString, &unicodeString, true)
	w.LdrGetProcedureAddressForCaller(hModule, &ansiString, 0, &functionHandle, 0, 0)
	return functionHandle
}

func (w *WinAPI) WriteProcessMemory(baseAddress uintptr, buffer []byte) bool {
	currentProcess, _, _ := w.ProcGetCurrentProcess.Call()
	if currentProcess == 0 {
		return false
	}

	var bytesWritten uintptr
	ret, _, _ := w.ProcWriteProcessMemory.Call(
		currentProcess,
		baseAddress,
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(len(buffer)),
		uintptr(unsafe.Pointer(&bytesWritten)),
	)
	return ret != 0 && bytesWritten == uintptr(len(buffer))
}

func (w *WinAPI) IsDebuggerPresent() bool {
	flag, _, _ := w.ProcIsDebuggerPresent.Call()
	return flag != 0
}

func (w *WinAPI) CheckRemoteDebugger() (bool, error) {
	var isDebuggerPresent bool
	r1, _, err := w.ProcCheckRemoteDebugger.Call(^uintptr(0), uintptr(unsafe.Pointer(&isDebuggerPresent)))
	if r1 == 0 {
		return false, err
	}
	return isDebuggerPresent, nil
}

func (w *WinAPI) EnumProcesses() ([]uint32, error) {
	var processIds [1024]uint32
	var bytesReturned uint32
	
	ret, _, err := w.ProcEnumProcesses.Call(
		uintptr(unsafe.Pointer(&processIds[0])),
		uintptr(len(processIds)*4),
		uintptr(unsafe.Pointer(&bytesReturned)),
	)
	
	if ret == 0 {
		return nil, err
	}
	
	numProcesses := bytesReturned / 4
	result := make([]uint32, numProcesses)
	copy(result, processIds[:numProcesses])
	
	return result, nil
}

func (w *WinAPI) GetProcessName(processId uint32) (string, error) {
	handle, _, err := w.ProcOpenProcess.Call(
		windows.PROCESS_QUERY_INFORMATION|windows.PROCESS_VM_READ,
		0,
		uintptr(processId),
	)
	
	if handle == 0 {
		return "", err
	}
	defer w.ProcCloseHandle.Call(handle)
	
	var processName [260]uint16
	ret, _, err := w.ProcGetModuleBaseName.Call(
		handle,
		0,
		uintptr(unsafe.Pointer(&processName[0])),
		uintptr(len(processName)),
	)
	
	if ret == 0 {
		return "", err
	}
	
	return syscall.UTF16ToString(processName[:]), nil
}

func (w *WinAPI) GetRunningProcessNames() ([]string, error) {
	processIds, err := w.EnumProcesses()
	if err != nil {
		return nil, err
	}
	
	var processNames []string
	for _, pid := range processIds {
		if pid == 0 {
			continue
		}
		
		name, err := w.GetProcessName(pid)
		if err == nil && name != "" {
			processNames = append(processNames, name)
		}
	}
	
	return processNames, nil
}

func (w *WinAPI) GetRunningProcessCount() (int, error) {
	var ids [1024]uint32
	var needed uint32
	r1, _, err := w.ProcEnumProcesses.Call(
		uintptr(unsafe.Pointer(&ids[0])),
		uintptr(len(ids)*4),
		uintptr(unsafe.Pointer(&needed)),
	)
	if r1 == 0 {
		return 0, err
	}
	return int(needed / 4), nil
}

func (w *WinAPI) CallProc(proc uintptr, args ...uintptr) (uintptr, uintptr, error) {
	return syscall.SyscallN(proc, args...)
}

func (w *WinAPI) LastError() error {
	return syscall.GetLastError()
}

func (w *WinAPI) CreateFile(fileName string, desiredAccess uint32, shareMode uint32, securityAttributes uintptr, creationDisposition uint32, flagsAndAttributes uint32, templateFile uintptr) uintptr {
	fileNamePtr, _ := syscall.UTF16PtrFromString(fileName)
	handle, _, _ := w.ProcCreateFile.Call(
		uintptr(unsafe.Pointer(fileNamePtr)),
		uintptr(desiredAccess),
		uintptr(shareMode),
		securityAttributes,
		uintptr(creationDisposition),
		uintptr(flagsAndAttributes),
		templateFile,
	)
	return handle
}

func (w *WinAPI) Fopen(filename string, mode string) uintptr {
	filenamePtr, _ := syscall.BytePtrFromString(filename)
	modePtr, _ := syscall.BytePtrFromString(mode)
	file, _, _ := w.ProcFopen.Call(uintptr(unsafe.Pointer(filenamePtr)), uintptr(unsafe.Pointer(modePtr)))
	return file
}

func (w *WinAPI) Fclose(file uintptr) int {
	ret, _, _ := w.ProcFclose.Call(file)
	return int(ret)
}

func (w *WinAPI) ZeroOutMemory(start uintptr, length int) bool {
	var oldProtect uint32
	
	ret, _, _ := w.ProcVirtualProtect.Call(
		start,
		uintptr(length),
		uintptr(0x04), 
		uintptr(unsafe.Pointer(&oldProtect)),
	)
	if ret == 0 {
		return false
	}
	
	for i := 0; i < length; i++ {
		*(*byte)(unsafe.Pointer(start + uintptr(i))) = 0
	}
	
	ret, _, _ = w.ProcVirtualProtect.Call(
		start,
		uintptr(length),
		uintptr(oldProtect),
		uintptr(unsafe.Pointer(&oldProtect)),
	)
	
	return ret != 0
}

func (w *WinAPI) FreeMemory(address uintptr) bool {
	ret, _, _ := w.ProcVirtualFree.Call(address,0, uintptr(0x8000))
	return ret != 0
}

func (w *WinAPI) GetCurrentProcess() uintptr {
	handle, _, _ := w.ProcGetCurrentProcess.Call()
	return handle
}

func (w *WinAPI) SetLastError(errCode uint32) {
	w.ProcSetLastError.Call(uintptr(errCode))
}

func (w *WinAPI) GetLastError() uint32 {
	ret, _, _ := w.ProcGetLastError.Call()
	return uint32(ret)
}

func (w *WinAPI) SetDebugFilterState(componentId uint64, level uint32, state bool) bool {
	ret, _, _ := w.ProcNtSetDebugFilterState.Call(
		uintptr(componentId),
		uintptr(level),
		uintptr(boolToInt(state)),
	)
	return ret == 0
}

func (w *WinAPI) QueryInformationProcess(handle syscall.Handle, infoClass uint32, info uintptr, infoLen uint32) (uintptr, uintptr, error) {
	return w.ProcNtQueryInformationProcess.Call(
		uintptr(handle),
		uintptr(infoClass),
		info,
		uintptr(infoLen),
		0,
	)
}

func (w *WinAPI) GetDisplayRefreshRate() (uint32, error) {
	const ENUM_CURRENT_SETTINGS = 0xFFFFFFFF

	var devMode DEVMODE
	devMode.DmSize = uint16(unsafe.Sizeof(devMode))

	ret, _, err := w.ProcEnumDisplaySettingsW.Call(
		0,
		uintptr(ENUM_CURRENT_SETTINGS),
		uintptr(unsafe.Pointer(&devMode)),
	)

	if ret == 0 {
		return 0, err
	}

	return devMode.DmDisplayFrequency, nil
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}