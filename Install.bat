@Echo off
title DOWNLOADING MODULES
go get github.com/EvilBytecode/GoDefender/AntiVirtualization/TriageDetection
go get github.com/EvilBytecode/GoDefender/AntiVirtualization/MonitorMetrics
go get github.com/EvilBytecode/GoDefender/AntiVirtualization/VirtualboxDetection
go get github.com/EvilBytecode/GoDefender/AntiVirtualization/VMWareDetection
go get github.com/EvilBytecode/GoDefender/AntiVirtualization/KVMCheck
go get github.com/EvilBytecode/GoDefender/AntiVirtualization/UsernameCheck
go get github.com/EvilBytecode/GoDefender/AntiDebug/IsDebuggerPresent
go get github.com/EvilBytecode/GoDefender/AntiDebug/RemoteDebugger
go get github.com/EvilBytecode/GoDefender/AntiDebug/pcuptime
go get github.com/EvilBytecode/GoDefender/AntiDebug/CheckBlacklistedWindowsNames
go get github.com/EvilBytecode/GoDefender/AntiDebug/RunningProcesses
go get github.com/EvilBytecode/GoDefender/AntiDebug/ParentAntiDebug
go get github.com/EvilBytecode/GoDefender/AntiDebug/KillBadProcesses
go get github.com/EvilBytecode/GoDefender/Process/CriticalProcess
go get github.com/EvilBytecode/GoDefender/AntiDebug/UserAntiAntiDebug
pause
