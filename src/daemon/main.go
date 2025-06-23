package main

import (
	"daemon/process"
	"fmt"
	"os"
)

func main() {
	// Create new process (daemon)
	proc, err := process.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize daemon: %v\n", err)
		os.Exit(1)
	}

	// Run the daemon
	if err := proc.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "daemon exited with error: %v\n", err)
		os.Exit(1)
	}
}
