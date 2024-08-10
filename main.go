package main

import (
	"eevscan/laser"
	"fmt"
	"log"
)

func main() {

	lc, err := laser.NewLaserController(21)
	if err != nil {
		log.Fatalf("Failed to initialize laser controller: %v", err)
	}

	go lc.StartPinsPolling()

	for change := range lc.PinChanges {
		fmt.Printf("Pin state changed to: %t\n", change)
	}
}
