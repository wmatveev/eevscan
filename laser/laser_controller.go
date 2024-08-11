package laser

import (
	"eevscan/device"
	"eevscan/events"
	"log"
	"time"
)

type Controller struct {
	deviceController *device.Controller
	Pause            chan bool
	Resume           chan bool
}

func NewLaserController(deviceAddr uint16) (*Controller, error) {
	dc, err := device.NewDeviceController(deviceAddr, 0x01)
	if err != nil {
		log.Fatalf("Failed to initialize laser controller: %v", err)
	}

	_ = dc.WriteToDevice(0x00)

	return &Controller{
		deviceController: dc,
		Pause:            make(chan bool),
		Resume:           make(chan bool),
	}, nil
}

func (lc *Controller) StartPinsPolling(eventManager *events.EventManager) {
	var lastState bool
	pause := false

	for {
		select {
		case <-lc.Pause:
			pause = true
		case <-lc.Resume:
			pause = false
		default:
			if !pause {
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
	}
}
