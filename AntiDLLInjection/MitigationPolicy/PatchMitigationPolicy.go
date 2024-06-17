package MitigationPolicyPatch

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	ProcessSignaturePolicyMitigation = 8
)

type PROCESS_MITIGATION_BINARY_SIGNATURE_POLICY struct {
	MicrosoftSignedOnly uint32
}

var (
	modkernelbase                 = syscall.NewLazyDLL("kernelbase.dll")
	procSetProcessMitigationPolicy = modkernelbase.NewProc("SetProcessMitigationPolicy")
)

func SetProcessMitigationPolicy(policy int, lpBuffer *PROCESS_MITIGATION_BINARY_SIGNATURE_POLICY, size uint32) (bool, error) {
	ret, _, err := procSetProcessMitigationPolicy.Call(uintptr(policy),uintptr(unsafe.Pointer(lpBuffer)),uintptr(size),)
	if ret != 0 {
		return true, nil
	}
	if err != nil && err.Error() != "The operation completed successfully." {
		return false, err
	}
	return false, nil
}

func ConfigureProcessMitigationPolicy() {
	var OnlyMicrosoftBinaries PROCESS_MITIGATION_BINARY_SIGNATURE_POLICY
	OnlyMicrosoftBinaries.MicrosoftSignedOnly = 1

	success, err := SetProcessMitigationPolicy(ProcessSignaturePolicyMitigation,
		&OnlyMicrosoftBinaries,
		uint32(unsafe.Sizeof(OnlyMicrosoftBinaries)),
	)
	if err != nil {
		fmt.Println("Failed:", err.Error())
		return
	}
	if success {
		fmt.Println("Success")
	} else {
		fmt.Println("Failed")
	}
}
