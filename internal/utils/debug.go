package utils

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"
)

var (
	DebugEnabled bool = true
	api *WinAPI = NewWinAPI()
)

func DbgPrint(format string, args ...interface{}) {
	if !DebugEnabled {
		return
	}
	msg := fmt.Sprintf(format, args...)
	messagePtr, _ := syscall.UTF16PtrFromString(msg)
	api.CallProc(api.ProcOutputDebugString.Addr(), 
	uintptr(unsafe.Pointer(messagePtr)))
}

func Print(format string, args ...interface{}) {
	if !DebugEnabled {
		return
	}
	log.Printf("[DEBUG] %s", fmt.Sprintf(format, args...))
}

