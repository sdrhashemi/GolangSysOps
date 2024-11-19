package main

import (
	"crypto/rand"
	"math/big"
)

// function for generating random file name
func CreateRandomFileName(fileNameLength int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, fileNameLength)

	for i := range result {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		result[i] = letters[index.Int64()]
	}

	return string(result)
}

// function for generating random data
func CreateRandomData(dataSize int) []byte {
	fileData := make([]byte, dataSize)

	_, err := rand.Read(fileData)
	if err != nil {
		panic(err)
	}

	return fileData
}
