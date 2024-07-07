package AllTokens

import (
    "golang.org/x/sys/windows"
)

var tokens = []string{
    "SeAssignPrimaryTokenPrivilege",
    "SeAuditPrivilege",
    "SeBackupPrivilege",
    "SeChangeNotifyPrivilege",
    "SeCreateGlobalPrivilege",
    "SeCreatePagefilePrivilege",
    "SeCreatePermanentPrivilege",
    "SeCreateSymbolicLinkPrivilege",
    "SeCreateTokenPrivilege",
    "SeDebugPrivilege",
    "SeDelegateSessionUserImpersonatePrivilege",
    "SeEnableDelegationPrivilege",
    "SeImpersonatePrivilege",
    "SeIncreaseQuotaPrivilege",
    "SeIncreaseBasePriorityPrivilege",
    "SeIncreaseWorkingSetPrivilege",
    "SeLoadDriverPrivilege",
    "SeLockMemoryPrivilege",
    "SeMachineAccountPrivilege",
    "SeManageVolumePrivilege",
    "SeProfileSingleProcessPrivilege",
    "SeRelabelPrivilege",
    "SeRemoteShutdownPrivilege",
    "SeRestorePrivilege",
    "SeSecurityPrivilege",
    "SeShutdownPrivilege",
    "SeSyncAgentPrivilege",
    "SeSystemtimePrivilege",
    "SeSystemEnvironmentPrivilege",
    "SeSystemProfilePrivilege",
    "SeTakeOwnershipPrivilege",
    "SeTcbPrivilege",
    "SeTimeZonePrivilege",
    "SeTrustedCredManAccessPrivilege",
    "SeUndockPrivilege",
}

func Enable() {
    hProcess := windows.CurrentProcess()
    var hToken windows.Token

    err := windows.OpenProcessToken(hProcess, windows.TOKEN_ADJUST_PRIVILEGES|windows.TOKEN_QUERY, &hToken)
    if err != nil {
        return
    }
    defer hToken.Close()

    for _, token := range tokens {
        var luid windows.LUID
        err := windows.LookupPrivilegeValue(nil, windows.StringToUTF16Ptr(token), &luid)
        if err != nil {
            continue
        }

        tp := windows.Tokenprivileges{
            PrivilegeCount: 1,
            Privileges: [1]windows.LUIDAndAttributes{
                {Luid: luid, Attributes: windows.SE_PRIVILEGE_ENABLED},
            },
        }

        err = windows.AdjustTokenPrivileges(hToken, false, &tp, 0, nil, nil)
        if err != nil {
            continue
        }
    }
}
