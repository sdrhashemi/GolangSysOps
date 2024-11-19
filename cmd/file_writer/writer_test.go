package main

import (
	"os"
	"testing"
	"time"
)

func TestWriteFiles(t *testing.T) {
	directory := t.TempDir() // create a temporary directory for testing
	stop := make(chan struct{})

	go func() {
		time.Sleep(2 * time.Second)
		close(stop)
	}()

	WriteRandomFiles(directory, stop)

	files, err := os.ReadDir(directory)
	if err != nil {
		t.Errorf("Failed to read directory %s: %v", directory, err)
	}

	if len(files) == 0 {
		t.Errorf("No files were created in directory %s", directory)
	}
}
