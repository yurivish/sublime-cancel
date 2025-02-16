package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create channel to receive signals
	sigChan := make(chan os.Signal, 1)

	// Notify sigChan when these signals are received
	signal.Notify(sigChan,
		syscall.SIGTERM, // "kill pid" or "kill -15 pid"
		syscall.SIGINT,  // Ctrl+C
		syscall.SIGQUIT, // Ctrl+\
		syscall.SIGHUP,  // Terminal closed
		syscall.SIGABRT, // Abort signal
		syscall.SIGTRAP, // Trace/breakpoint trap
	)

	fmt.Println("Process started. PID:", os.Getpid())
	fmt.Println("Send signals to test (e.g., kill -SIGTERM", os.Getpid(), ")")
	fmt.Println("Press Ctrl+C to send SIGINT")
	fmt.Println("Press Ctrl+\\ to send SIGQUIT")

	// Wait for a signal
	for {
		sig := <-sigChan

		// Print the signal we received
		fmt.Printf("\nCaught signal: %v\n", sig)

		// Ensure output is written before exiting
		os.Stdout.Sync()

		// Exit with different status codes based on the signal
		switch sig {
		case syscall.SIGTERM:
			os.Exit(143) // Standard SIGTERM exit code
		case syscall.SIGINT:
			os.Exit(130) // Standard SIGINT exit code
		default:
			os.Exit(1)
		}
	}
}
