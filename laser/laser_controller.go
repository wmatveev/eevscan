package laser

import (
	"eevscan/device"
	"eevscan/events"
	"log"
	"time"
)

type Controller struct {
	deviceController *device.Controller
}

func NewLaserController(deviceAddr uint16) (*Controller, error) {
	dc, err := device.NewDeviceController(deviceAddr, 0x01)
	if err != nil {
		log.Fatalf("Failed to initialize laser controller: %v", err)
	}

	_ = dc.WriteToDevice(0x00)

	return &Controller{
		deviceController: dc,
	}, nil
}

func (lc *Controller) StartPinsPolling(eventManager *events.EventManager) {
	var lastState bool

	for {
		readData, err := lc.deviceController.ReadingFromDevice()
		if err != nil {
			log.Fatalf("Failed to read from device: %v", err)
		}

		log.Printf("---> readData: %b\n", readData)

		currentState := readData&0x03 != 0x00

		if currentState != lastState && currentState == true {
			eventManager.Publish(events.Event{
				Type:    events.EventObjectEnteredToZone,
				Payload: currentState,
			})

			lastState = currentState
		}

		time.Sleep(1 * time.Second)
	}
}
