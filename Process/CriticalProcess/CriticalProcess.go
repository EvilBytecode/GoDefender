package programutils

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	modadvapi32 = syscall.NewLazyDLL("advapi32.dll")
	modntdll    = syscall.NewLazyDLL("ntdll.dll")
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")
	procLookupPrivilegeValueW  = modadvapi32.NewProc("LookupPrivilegeValueW")
	procAdjustTokenPrivileges  = modadvapi32.NewProc("AdjustTokenPrivileges")
	procRtlSetProcessIsCritical = modntdll.NewProc("RtlSetProcessIsCritical")
	procGetCurrentProcess       = modkernel32.NewProc("GetCurrentProcess")
	procCloseHandle             = modkernel32.NewProc("CloseHandle")
)

const (
	SE_DEBUG_NAME       = "SeDebugPrivilege"
	SE_PRIVILEGE_ENABLED = 0x00000002
)

type LUID struct {
	LowPart  uint32
	HighPart int32
}

type LUIDAndAttributes struct {
	Luid       LUID
	Attributes uint32
}

type Tokenprivileges struct {
	PrivilegeCount uint32
	Privileges     [1]LUIDAndAttributes
}

func LookupPrivilegeValue(systemname *uint16, name *uint16, luid *LUID) error {
	r1, _, e1 := procLookupPrivilegeValueW.Call(uintptr(unsafe.Pointer(systemname)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(luid)))
	if r1 == 0 {
		return syscall.Errno(e1.(syscall.Errno))
	}
	return nil
}

func AdjustTokenPrivileges(token syscall.Handle, disableAllPrivileges bool, newstate *Tokenprivileges, buflen uint32, prevstate *Tokenprivileges, returnlen *uint32) error {
	var _p0 uint32
	if disableAllPrivileges {
		_p0 = 1
	}
	r1, _, e1 := procAdjustTokenPrivileges.Call(uintptr(token), uintptr(_p0), uintptr(unsafe.Pointer(newstate)), uintptr(buflen), uintptr(unsafe.Pointer(prevstate)), uintptr(unsafe.Pointer(returnlen)))
	if r1 == 0 {
		return syscall.Errno(e1.(syscall.Errno))
	}
	return nil
}

func SetDebugPrivilege() error {
	var token syscall.Handle
	r1, _, e1 := procGetCurrentProcess.Call()
	if r1 == 0 {
		return syscall.Errno(e1.(syscall.Errno))
	}
	token = syscall.Handle(r1)

	defer syscall.CloseHandle(token)

	var tokenHandle syscall.Token
	err := syscall.OpenProcessToken(token, syscall.TOKEN_ADJUST_PRIVILEGES|syscall.TOKEN_QUERY, &tokenHandle)
	if err != nil {
		return err
	}

	var luid LUID
	err = LookupPrivilegeValue(nil, syscall.StringToUTF16Ptr(SE_DEBUG_NAME), &luid)
	if err != nil {
		return err
	}

	tp := Tokenprivileges{
		PrivilegeCount: 1,
		Privileges: [1]LUIDAndAttributes{
			{
				Luid:       luid,
				Attributes: SE_PRIVILEGE_ENABLED,
			},
		},
	}

	err = AdjustTokenPrivileges(syscall.Handle(tokenHandle), false, &tp, uint32(unsafe.Sizeof(tp)), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func SetProcessCritical() error {
	err := SetDebugPrivilege()
	if err != nil {
		return err
	}

	r1, _, e1 := procRtlSetProcessIsCritical.Call(uintptr(1), 0, 0)
	if r1 != 0 {
		return nil
	}
	return syscall.Errno(e1.(syscall.Errno))
}
