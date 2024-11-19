package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func CalculateSHA256(filePath string) (string, error) {
	// Open the file for reading.
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file for hashing: %v", err)
	}
	defer file.Close()

	hasher := sha256.New()
	reader := bufio.NewReader(file)

	// copy file content to hasher
	if _, err := reader.WriteTo(hasher); err != nil {
		return "", fmt.Errorf("failed to compute hash: %v", err)
	}

	// encode hash to hexadecimal
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash, nil
}
