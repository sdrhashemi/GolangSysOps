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
		// waiting for signal to shutdwon
		case <-stop:
			fmt.Println("Stop file creator")
			return

		default:
			// getting the file name
			filename := filepath.Join(directory, CreateRandomFileName(FileNameLength))
			// getting data with fixed size
			fileRandomData := CreateRandomData(FileSize)

			// write the file with 0644 permission for reading and writing
			if err := os.WriteFile(filename, fileRandomData, 0644); err != nil {
				fmt.Printf("Failed to write file %s: %v\n", filename, err)
				return
			}

			fmt.Printf("Wrote file %s\n", filename)
			time.Sleep(Interval * time.Second) // for the 1 second in the task
		}

	}
}
