package state

import (
	"eevscan/device"
	"eevscan/events"
	"eevscan/laser"
	"eevscan/scanner"
	"log"
)

type StateManager struct {
	eventManager      *events.EventManager
	laserController   *laser.Controller
	scannerController *scanner.Controller
	portController    *device.PortController
}

func NewStateManager(lc *laser.Controller, sc *scanner.Controller, pc *device.PortController) *StateManager {
	eventManager := events.NewEventManager()
	stateManager := &StateManager{
		eventManager:      eventManager,
		laserController:   lc,
		scannerController: sc,
		portController:    pc,
	}

	eventManager.Subscribe(events.EventObjectEnteredToZone, stateManager.handleObjectEnteredToZone)

	return stateManager
}

func (sm *StateManager) handleObjectEnteredToZone(event events.Event) {
	log.Println("Object Entered To Zone")

	select {
	case sm.laserController.Pause <- true:
		log.Println("Laser controller paused")
	default:
		log.Println("Laser controller already paused")
	}

	sm.portController.StartPortsReading()

	sm.scannerController.ActivateScanner()

	go func() {
		select {
		case barcode := <-sm.portController.Barcode:
			log.Printf("Received barcode: %s", string(barcode))
			close(sm.portController.QuitChannel)

		case <-sm.portController.QuitChannel:
			log.Println("Stopping barcode reading goroutine")
			return
		}
	}()

	log.Printf("---> 4")

	select {
	case sm.laserController.Resume <- true:
		log.Println("Laser controller resumed")
	default:
		log.Println("Laser controller was not paused")
	}
}

func (sm *StateManager) Start() {
	go sm.laserController.StartPinsPolling(sm.eventManager)
}
