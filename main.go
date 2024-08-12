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

	rs, err := device.NewRS232Controller()
	if err != nil {
		log.Fatalf("Failed to initialize rs232 controller: %v", err)
	}

	stateManager := state.NewStateManager(lc, sc, pc, rs)
	stateManager.Start()

	// Создаем канал для обработки сигнала завершения приложения
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	stateManager.EventManager.Publish(events.Event{
		Type: events.EventShutdown,
	})

	//select {}
}
