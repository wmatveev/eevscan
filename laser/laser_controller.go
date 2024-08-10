package laser

import (
	"eevscan/device"
	"log"
	"time"
)

type Controller struct {
	deviceController *device.Controller
	PinChanges       chan bool
}

func NewLaserController(deviceAddr uint16) (*Controller, error) {
	dc, err := device.NewDeviceController(deviceAddr, 0x01)
	if err != nil {
		log.Fatalf("Failed to initialize laser controller: %v", err)
	}

	return &Controller{
		deviceController: dc,
		PinChanges:       make(chan bool),
	}, nil
}

func (lc *Controller) StartPinsPolling() {
	var lastState bool

	for {
		readData, err := lc.deviceController.ReadingFromDevice()
		if err != nil {
			log.Fatalf("Failed to read from device: %v", err)
		}

		currentState := readData&0x01 == 0x01

		if currentState != lastState {
			lc.PinChanges <- currentState
			lastState = currentState
		}

		time.Sleep(1 * time.Second)
	}
}
