# Go Defender

![Go Defender](GoDefender.png)

## GoDefender

This Go package provides functionality to detect and defend against various forms of debugging tools, virtualization environments.
btw for quick setup run install.bat
### Anti-Virtualization

- **Triage Detection**: Detects if the system is running in a triage or analysis environment.
- **Monitor Metrics**: Monitors system metrics to identify abnormal behavior indicative of virtualization.
- **VirtualBox Detection**: Detects the presence of Oracle VirtualBox.
- **VMWare Detection**: Detects the presence of VMware virtualization software.
- **KVM Check**: Checks for Kernel-based Virtual Machine (KVM) hypervisor.
- **Username Check**: Verifies if the current user is a default virtualization user.

### Anti-Debug

This module includes functions to detect and prevent debugging and analysis of the running process.

- **IsDebuggerPresent**: Checks if a debugger is currently attached to the process.
- **Remote Debugger**: Detects if a remote debugger is connected to the process.
- **PC Uptime**: Monitors system uptime to detect debugging attempts based on system restarts.
- **Check Blacklisted Windows Names**: Verifies if the process name matches any blacklisted names commonly used by debuggers.
- **Running Processes**: Retrieves a list of running processes and identifies potential malicious ones.
- **Parent Anti-Debug**: Detects if the parent process is attempting to debug the current process.
- **Kill Bad Processes**: Terminates known malicious processes detected on the system.
- **Detects Usermode AntiAntiDebuggers**: ScyllaHide.. (BASIC)
### Process

This module focuses on critical processes that should be monitored or protected.
- **Critical Process**: Implements functionality to manage critical processes essential for system operation.
- **SetDebugPrivilege**: Grants better permissions.
# Quick Nutshell
- Detecting Most Anti Anti-Debugging Hooking Methods on Common Anti-Debugging Functions by checking for Bad Instructions on Functions Addresses (Most Effective on x64) and it detects user-mode anti anti-debuggers like scyllahide, and it can also detect some sandboxes which uses hooking to monitor application behaviour/activity (like [Tria.ge](https://tria.ge/))

# V1.0.5 (soon)
- usb check
- internet check
- recent file activity
- soon more but codes are uploaded but not pushed yet.
## Credits:
- https://github.com/AdvDebug = Inspired me to start making this package, without him it wouldnt be here, check out his github.
- https://github.com/MmCopyMemory = Giving ideas and many more, check out his github.
