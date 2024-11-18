package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	p := os.Args[1]                    // directory
	r, err := strconv.Atoi(os.Args[2]) // number of random files creators
	if err != nil {
		fmt.Printf("Invalid number of random files: %v\n", err)
		return
	}

	if _, err := os.Stat(p); os.IsNotExist(err) {
		fmt.Printf("Directory %s does not exsit. Creating...\n", p)
		if err := os.MkdirAll(p, 0755); err != nil {
			fmt.Printf("Failed to create directory %s: %v\n", p, err)
			return
		}
	}

	stop := make(chan struct{})
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
		<-signalChan
		fmt.Println("Stopping...")
		close(stop)
	}()

	var wg sync.WaitGroup
	for i := 0; i < r; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			WriteRandomFiles(p, stop)
		}()
	}
	wg.Wait()
}
