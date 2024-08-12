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
	log.Println(event)
	//log.Println("Object Entered To Zone")

	sm.laserController.Pause()

	sm.scannerController.ActivateScanner()
	sm.portController.StartPortsReading()
	sm.scannerController.DeactivateScanner()

	sm.laserController.Resume()
}

func (sm *StateManager) Start() {
	go sm.laserController.StartPinsPolling(sm.eventManager)
}
