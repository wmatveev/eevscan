package laser

import (
	"eevscan/device"
	"eevscan/events"
	"log"
	"sync"
	"time"
)

type Controller struct {
	deviceController *device.DeviceController
	mu               sync.Mutex
	paused           bool
}

func NewLaserController(deviceAddr uint16) (*Controller, error) {
	dc, err := device.NewDeviceController(deviceAddr, 0x01)
	if err != nil {
		log.Fatalf("Failed to initialize laser controller: %v", err)
	}

	_ = dc.WriteToDevice(0x00)

	return &Controller{
		deviceController: dc,
		paused:           false,
	}, nil
}

func (lc *Controller) StartPinsPolling(eventManager *events.EventManager) {
	var lastState bool

	for {
		lc.mu.Lock()
		if lc.paused {
			lc.mu.Unlock()
			time.Sleep(10 * time.Millisecond)
			continue
		}
		lc.mu.Unlock()

		readData, err := lc.deviceController.ReadingFromDevice()
		if err != nil {
			log.Fatalf("Failed to read from device: %v", err)
		}

		currentState := readData&0x03 != 0x00

		if currentState != lastState {
			if currentState == true {
				eventManager.Publish(events.Event{
					Type:    events.EventObjectEnteredToZone,
					Payload: currentState,
				})
			}

			lastState = currentState
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func (lc *Controller) Pause() {
	lc.mu.Lock()
	lc.paused = true
	lc.mu.Unlock()
}

func (lc *Controller) Resume() {
	lc.mu.Lock()
	lc.paused = false
	lc.mu.Unlock()
}
