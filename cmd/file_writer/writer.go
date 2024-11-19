package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func WriteRandomFiles(directory string, stop <-chan struct{}) {
	for {
		select {

		case <-stop:
			fmt.Println("Stop file creator")
			return

		default:
			filename := filepath.Join(directory, CreateRandomFileName(FileNameLength))
			fileRandomData := CreateRandomData(FileSize)

			if err := os.WriteFile(filename, fileRandomData, 0644); err != nil {
				fmt.Printf("Failed to write file %s: %v\n", filename, err)
				return
			}

			fmt.Printf("Wrote file %s\n", filename)
			time.Sleep(Interval * time.Second) // for the 1 second in the task
		}

	}
}
