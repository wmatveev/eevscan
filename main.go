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
		if change == true {
			_ = sc.ActivateScanner()
		} else {
			_ = sc.DeactivateScanner()
		}

		time.Sleep(1 * time.Second)

		fmt.Printf("Pin state changed to: %t\n", change)
	}
}
