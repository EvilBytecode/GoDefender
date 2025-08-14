package CleanEnvironmentDetection

import (
	"strings"

	"golang.org/x/sys/windows/registry"
)

// DetectCleanEnvironment checks if at least 10 programs are installed on the system.
func DetectCleanEnvironment() bool {
	total := countInstalledPrograms()
	return total >= 10
}

// countInstalledPrograms enumerates installed programs from registry uninstall keys.
func countInstalledPrograms() int {
	uninstallKeys := []struct {
		root registry.Key
		path string
	}{
		{registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`},
		{registry.LOCAL_MACHINE, `SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall`},
		{registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`},
	}

	count := 0
	seen := make(map[string]bool)

	for _, uk := range uninstallKeys {
		key, err := registry.OpenKey(uk.root, uk.path, registry.READ)
		if err != nil {
			continue
		}
		defer key.Close()

		names, err := key.ReadSubKeyNames(-1)
		if err != nil {
			continue
		}

		for _, name := range names {
			subKey, err := registry.OpenKey(uk.root, uk.path+`\`+name, registry.READ)
			if err != nil {
				continue
			}

			displayName, _, err := subKey.GetStringValue("DisplayName")
			subKey.Close()

			if err == nil && displayName != "" {
				normalized := strings.ToLower(displayName)
				if !seen[normalized] {
					seen[normalized] = true
					count++
				}
			}
		}
	}

	return count
}
