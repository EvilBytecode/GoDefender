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
	modkernel32                    = syscall.NewLazyDLL("kernel32.dll")
	procSetProcessMitigationPolicy = modkernel32.NewProc("SetProcessMitigationPolicy")
)

// SetProcessMitigationPolicy calls the Windows API SetProcessMitigationPolicy
func SetProcessMitigationPolicy(policy int, lpBuffer *PROCESS_MITIGATION_BINARY_SIGNATURE_POLICY, size uint32) (bool, error) {
	ret, _, err := procSetProcessMitigationPolicy.Call(
		uintptr(policy),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(size),
	)
	if ret != 0 {
		return true, nil
	}
	if err != nil && err.Error() != "The operation completed successfully." {
		return false, err
	}
	return false, nil
}

func ConfigureProcessMitigationPolicy() {
	var policy PROCESS_MITIGATION_BINARY_SIGNATURE_POLICY
	policy.MicrosoftSignedOnly = 1

	ok, err := SetProcessMitigationPolicy(ProcessSignaturePolicyMitigation, &policy, uint32(unsafe.Sizeof(policy)))
	if err != nil {
		fmt.Printf("Failed to set mitigation policy: %v\n", err)
	} else if ok {
		fmt.Println("Mitigation policy set successfully.")
	} else {
		fmt.Println("Failed to set mitigation policy: Unknown error.")
	}
}
