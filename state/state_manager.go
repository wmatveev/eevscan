package state

import (
	"eevscan/device"
	"eevscan/events"
	"eevscan/laser"
	"eevscan/scanner"
)

type StateManager struct {
	EventManager      *events.EventManager
	laserController   *laser.Controller
	scannerController *scanner.Controller
	portController    *device.PortController
}

func NewStateManager(lc *laser.Controller, sc *scanner.Controller, pc *device.PortController) *StateManager {
	eventManager := events.NewEventManager()
	stateManager := &StateManager{
		EventManager:      eventManager,
		laserController:   lc,
		scannerController: sc,
		portController:    pc,
	}

	eventManager.Subscribe(events.EventObjectEnteredToZone, stateManager.handleObjectEnteredToZone)
	eventManager.Subscribe(events.EventShutdown, stateManager.handleShutdown)

	return stateManager
}

func (sm *StateManager) handleObjectEnteredToZone(event events.Event) {
	sm.laserController.Pause()

	sm.scannerController.ActivateScanner()
	sm.portController.StartPortsReading()
	sm.scannerController.DeactivateScanner()

	sm.laserController.Resume()

	_ = event
}

func (sm *StateManager) Start() {
	go sm.laserController.StartPinsPolling(sm.EventManager)
}

func (sm *StateManager) handleShutdown(event events.Event) {

}
