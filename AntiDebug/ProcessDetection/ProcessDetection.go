package BadProcesses

import (
	"bytes"
	"os/exec"
	"strings"
)

// DetectProcesses checks for specific processes and returns true if any are found.
func Detect() (bool, error) {
	ptk := []string{
		"taskmgr.exe", "process.exe", "processhacker.exe", "ksdumper.exe", "fiddler.exe",
		"httpdebuggerui.exe", "wireshark.exe", "httpanalyzerv7.exe", "decoder.exe",
		"regedit.exe", "procexp.exe", "dnspy.exe", "vboxservice.exe", "burpsuite.exe",
		"DbgX.Shell.exe", "ILSpy.exe", "ollydbg.exe", "x32dbg.exe", "x64dbg.exe", "gdb.exe",
		"idaq.exe", "idag.exe", "idaw.exe", "ida64.exe", "idag64.exe", "idaw64.exe",
		"idaq64.exe", "windbg.exe", "immunitydebugger.exe", "windasm.exe",
	}

	for _, prg := range ptk {
		cmd := exec.Command("tasklist", "/FI", "IMAGENAME eq "+prg)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			return false, err
		}

		processLines := strings.Split(out.String(), "\n")
		for _, line := range processLines {
			if strings.Contains(line, prg) {
				return true, nil
			}
		}
	}

	return false, nil
}
