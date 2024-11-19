package main

import (
	"fmt"
	"log"
	"os"

	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"
)

func closeWatcher(watcher *fsnotify.Watcher) {
	if err := watcher.Close(); err != nil {
		fmt.Printf("Error closing watcher: %v\n", err)
	}
}
func main() {

	dir := os.Args[1]

	// validate directory
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatalf("Directory does not exist: %s\n", dir)
	}

	// init a watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Failed to initialize watcher: %v\n", err)
		return
	}
	defer closeWatcher(watcher)

	// add dir to the watcher
	err = watcher.Add(dir)
	if err != nil {
		log.Fatalf("Failed to add directory to watcher: %v\n", err)
	}

	fmt.Printf("Monitoring directory: %s\n", dir)

	// gracful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// event proccess
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				fmt.Println("Watcher events channel closed")
				return
			}

			// enable create and writes events
			if event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Write == fsnotify.Write {
				ProcessEvent(event)
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				fmt.Println("Watcher errors channel closed")
				return
			}
			fmt.Printf("Watcher error: %v\n", err)

		case <-sigChan:
			fmt.Println("shutting down...")
			return
		}
	}
}
