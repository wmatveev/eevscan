package main

import (
	"eevscan/laser"
	"eevscan/scanner"
	"fmt"
	"log"
	"time"
)

func main() {

	lc, err := laser.NewLaserController(0x21)
	if err != nil {
		log.Fatalf("Failed to initialize laser controller: %v", err)
	}

	sc, err := scanner.NewScannerController(0x20)
	if err != nil {
		log.Fatalf("Failed to initialize laser controller: %v", err)
	}

	go lc.StartPinsPolling()

	for change := range lc.PinChanges {
		_ = sc.ActivateScanner()

		time.Sleep(1 * time.Second)

		_ = sc.DeactivateScanner()

		fmt.Printf("Pin state changed to: %t\n", change)
	}
}
