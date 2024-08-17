package state

import (
	"eevscan/device"
	"eevscan/events"
	"eevscan/laser"
	"eevscan/scanner"
	"log"
)

type StateManager struct {
	EventManager      *events.EventManager
	laserController   *laser.Controller
	scannerController *scanner.Controller
	portController    *device.PortController
	rs232Controller   *device.RS232Controller
}

func NewStateManager(lc *laser.Controller, sc *scanner.Controller, pc *device.PortController,
	rs *device.RS232Controller) *StateManager {
	eventManager := events.NewEventManager()
	stateManager := &StateManager{
		EventManager:      eventManager,
		laserController:   lc,
		scannerController: sc,
		portController:    pc,
		rs232Controller:   rs,
	}

	eventManager.Subscribe(events.EventObjectEnteredToZone, stateManager.handleObjectEnteredToZone)
	eventManager.Subscribe(events.EventShutdown, stateManager.handleShutdown)
	//eventManager.Subscribe(events.EventSendBarcodeToRS232, stateManager.handleSendBarcodeToRS232)

	return stateManager
}

func (sm *StateManager) Start() {
	go sm.laserController.StartPinsPolling(sm.EventManager)
}

func (sm *StateManager) handleObjectEnteredToZone(event events.Event) {
	sm.laserController.Pause()

	log.Println("---> Scanner Activated")

	sm.scannerController.ActivateScanner()

	barcode := sm.portController.StartPortsReading()
	if barcode != nil {
		log.Printf("---> barcode = %b\n", barcode)
		//sm.EventManager.Publish(events.Event{
		//	Type:    events.EventSendBarcodeToRS232,
		//	Payload: "JM-SH-OD-0145",
		//})
	}

	sm.scannerController.DeactivateScanner()

	log.Println("---> Scanner Deactivated")

	sm.laserController.Resume()
	_ = event
}

func (sm *StateManager) handleSendBarcodeToRS232(event events.Event) {
	if _, ok := event.Payload.(string); ok {
		sm.rs232Controller.Write("12345")
	} else {
		log.Println("Invalid payload for barcode, expected []byte")
	}
	_ = event
}

func (sm *StateManager) handleShutdown(event events.Event) {
	sm.scannerController.DeactivateScanner()

	log.Println("---> handleShutdown, Scanner Deactivated")

	_ = event
}
