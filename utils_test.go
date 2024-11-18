package main

import (
	"testing"
)

func TestCreateRandomFileName(t *testing.T) {
	t.Run("Testing Randomnes of file name", func(t *testing.T) {
		fileName1 := CreateRandomFileName(FileNameLength)
		fileName2 := CreateRandomFileName(FileNameLength)

		if fileName1 == fileName2 {
			t.Errorf("Expected different file names, got %s", fileName1)
		}
	})

	t.Run("Testing Length of file name", func(t *testing.T) {
		fileName := CreateRandomFileName(FileNameLength)
		if len(fileName) != FileNameLength {
			t.Errorf("Expected file name length %d, got %d", FileNameLength, len(fileName))
		}
	})

}

func TestCreateRandomData(t *testing.T) {
	t.Run("Testing data size to be 1MB", func(t *testing.T) {
		data := CreateRandomData(FileSize)
		if len(data) != FileSize {
			t.Errorf("Expected data length %d, got %d", FileSize, len(data))
		}
	})

	t.Run("Testing data in not empty", func(t *testing.T) {
		isEmpty := true
		data := CreateRandomData(FileSize)
		for _, b := range data {
			if b != 0 {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			t.Errorf("Expected data not to be empty")
		}
	})
}
