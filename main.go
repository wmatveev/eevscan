package main

import (
	"eevscan/device"
	"eevscan/events"
	"eevscan/laser"
	"eevscan/scanner"
	"eevscan/state"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	pc, err := device.NewPortController()
	if err != nil {
		log.Fatalf("Failed to initialize port controller: %v", err)
	}

	stateManager := state.NewStateManager(lc, sc, pc)
	stateManager.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Ожидание сигнала завершения
	<-sigChan

	// Публикация события завершения работы
	stateManager.EventManager.Publish(events.Event{
		Type: events.EventShutdown,
	})

	//select {}
}
