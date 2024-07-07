package UsernameCheck

import (
	"os"
	"strings"
)

// CheckForBlacklistedNames checks if the current username matches any blacklisted names.
// It returns true if a blacklisted name is found, otherwise false.
func CheckForBlacklistedNames() bool {
	bn := []string{"Johnson", "Miller", "malware", "maltest", "CurrentUser", "Sandbox", "virus", "John Doe", "test user", "sand box", "WDAGUtilityAccount", "Bruno", "george", "Harry Johnson"}
	user := strings.ToLower(os.Getenv("USERNAME"))
	for _, name := range bn {
		if user == strings.ToLower(name) {
			return true
		}
	}
	return false
}
