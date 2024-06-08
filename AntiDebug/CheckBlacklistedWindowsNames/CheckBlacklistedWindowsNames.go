package CheckBlacklistedWindowsNames

import (
    "log"
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

var blacklistedWindows = []string{
    "proxifier", "graywolf", "extremedumper", "zed", "exeinfope", "dnspy",
    "titanHide", "ilspy", "titanhide", "x32dbg", "codecracker", "simpleassembly",
    "process hacker 2", "pc-ret", "http debugger", "Centos", "process monitor",
    "debug", "ILSpy", "reverse", "simpleassemblyexplorer", "process", "de4dotmodded",
    "dojandqwklndoqwd-x86", "sharpod", "folderchangesview", "fiddler", "die", "pizza",
    "crack", "strongod", "ida -", "brute", "dump", "StringDecryptor", "wireshark",
    "debugger", "httpdebugger", "gdb", "kdb", "x64_dbg", "windbg", "x64netdumper",
    "petools", "scyllahide", "megadumper", "reversal", "ksdumper v1.1 - by equifox",
    "dbgclr", "HxD", "monitor", "peek", "ollydbg", "ksdumper", "http", "wpe pro", "dbg",
    "httpanalyzer", "httpdebug", "PhantOm", "kgdb", "james", "x32_dbg", "proxy", "phantom",
    "mdbg", "WPE PRO", "system explorer", "de4dot", "X64NetDumper", "protection_id",
    "charles", "systemexplorer", "pepper", "hxd", "procmon64", "MegaDumper", "ghidra", "xd",
    "0harmony", "dojandqwklndoqwd", "hacker", "process hacker", "SAE", "mdb", "checker",
    "harmony", "Protection_ID", "PETools", "scyllaHide", "x96dbg", "systemexplorerservice",
    "folder", "mitmproxy", "dbx", "sniffer", "Process Hacker", "Process Explorer", "Sysinternals", "www.sysinternals.com", "binary ninja",
}

// CheckBlacklistedWindows checks for blacklisted window names and terminates the associated process if found.
func CheckBlacklistedWindows() {
    pew.Call(syscall.NewCallback(enumWindowsProc), 0)
}

// enumWindowsProc is the callback function that checks each window title against the blacklist.
func enumWindowsProc(hwnd syscall.Handle, lParam uintptr) uintptr {
    var pid uint32
    pgwtp.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&pid)))

    var title [256]byte
    pgwt.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&title[0])), uintptr(len(title)))
    wt := string(title[:])

    // Check if the window title contains any blacklisted strings
    for _, blacklisted := range blacklistedWindows {
        if contains(wt, blacklisted) {
            log.Printf("Detected blacklisted window: %s\n", wt)
            // If a blacklisted window is found, terminate the associated process
            proc, _, _ := pop.Call(syscall.PROCESS_TERMINATE, 0, uintptr(pid))
            if proc != 0 {
                ptp.Call(proc, 0)
                pch.Call(proc)
            }
        }
    }
    return 1 // Continue enumeration
}

func contains(s, substr string) bool {
    return len(s) >= len(substr) && s[:len(substr)] == substr
}
