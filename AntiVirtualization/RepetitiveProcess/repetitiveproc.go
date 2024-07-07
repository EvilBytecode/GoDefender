package RepetitiveProcess

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"syscall"
)

// Check checks if any process with the same name is running more than 15 times and exits if so.
func Check() (bool, error) {
	cmd := exec.Command("tasklist")
	var out bytes.Buffer
	cmd.Stdout = &out

	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	err := cmd.Run()
	if err != nil {
		log.Printf("Error running tasklist command: %v", err)
		return false, err
	}

	processLines := strings.Split(out.String(), "\n")
	processCounts := make(map[string]int)
	for _, line := range processLines {
		fields := strings.Fields(line)
		if len(fields) > 0 {
			processName := fields[0]
			processCounts[processName]++
		}
	}

	for _, count := range processCounts {
		if count > 15 {
			return true, nil
		}
	}

	return false, nil
}
