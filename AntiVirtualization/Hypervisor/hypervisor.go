package hypervisor

import (
	"fmt"

	"github.com/intel-go/cpuid"
)

// Coded by Sown / zwinplayer64
//
// Checks the cpuid bit for hypervisor presence
func CheckHypervisorBit() bool {
	// the hypervisor bit is always 0 on real cpus
	const hypervisor = cpuid.HYPERVISOR
	if hypervisor != 0 && cpuid.HYPERV == 0 {
		fmt.Printf("The hypervisor cpuid bit isn't 0 (%x)\n", hypervisor)
		return true
	}
	fmt.Println("Hypervisor bit is 0.")
	return false
}
