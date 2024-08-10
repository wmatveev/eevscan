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

	//if _, err := host.Init(); err != nil {
	//	return nil, err
	//}
	//
	//bus, err := i2creg.Open("1")
	//if err != nil {
	//	return nil, err
	//}
	//
	//dev := i2c.Dev{Bus: bus, Addr: address}
	//
	//return &Controller{
	//	addr:       address,
	//	bus:        bus,
	//	device:     dev,
	//	PinChanges: make(chan bool),
	//}, nil
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

	//var readBuf [1]byte
	//var lastState bool
	//
	//for {
	//	if err := lc.device.Tx(nil, readBuf[:]); err != nil {
	//		log.Printf("Failed to read from device: %v\n", err)
	//		break
	//	}
	//
	//	currentState := readBuf[0]&0x01 == 0x01
	//
	//	if currentState != lastState {
	//		lc.PinChanges <- currentState
	//		lastState = currentState
	//	}
	//
	//	time.Sleep(1 * time.Second)
	//}
}
