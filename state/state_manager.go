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
}

func (sm *StateManager) Start() {
	log.Println("---> state_manager 5")
	go sm.laserController.StartPinsPolling(sm.eventManager)
}
