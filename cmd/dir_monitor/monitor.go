package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

// process a single notify event
func ProcessEvent(event fsnotify.Event) {
	filePath := event.Name

	time.Sleep(100 * time.Millisecond)

	// validate regular file
	_, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("Failed to stat file %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("Detected file: %s\n", filePath)

	// get PID
	pid, err := GetPID(filePath)
	if err != nil {
		fmt.Printf("Cant get PID for file %s: %v\n", filePath, err)
	} else {
		fmt.Printf("PID of file: %d\n", pid)
	}

	// calculate hash
	hash, err := CalculateSHA256(filePath)
	if err != nil {
		fmt.Printf("Error calculating SHA256 for %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("SHA256: %s\n", hash)

	// delete file
	err = os.Remove(filePath)
	if err != nil {
		fmt.Printf("Failed to delete file %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("File %s deleted successfully.\n", filePath)
}

// get PID for the file
func GetPID(filePath string) (int, error) {

	cmd := exec.Command("lsof", "-F", "p", "--", filePath)

	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("lsof error: %v", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "p") {

			pidStr := strings.TrimPrefix(line, "p")
			var pid int
			_, err := fmt.Sscanf(pidStr, "%d", &pid)
			if err != nil {
				continue
			}
			return pid, nil
		}
	}

	// no PID
	return 0, fmt.Errorf("no PID found for file: %s", filePath)
}
