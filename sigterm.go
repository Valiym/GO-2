package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-sigs:
			fmt.Println("Done")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Waiting for close signal")
		}
	}
}
