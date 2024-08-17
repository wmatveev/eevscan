package scanner

import (
	"eevscan/device"
	"log"
)

type Controller struct {
	deviceController *device.DeviceController
}

func NewScannerController(devAddress uint16) (*Controller, error) {
	dc, err := device.NewDeviceController(devAddress, 0x01)
	if err != nil {
		log.Fatalf("Failed to initialize scanner controller: %v", err)
	}

	return &Controller{
		deviceController: dc,
	}, nil
}

func (sc *Controller) ActivateScanner() {
	err := sc.deviceController.WriteToDevice(0x01)
	if err != nil {
		log.Fatalf("Failed to activate scanner controller: %v", err)
		return
	}

	return
}

func (sc *Controller) DeactivateScanner() {
	err := sc.deviceController.WriteToDevice(0x00)
	if err != nil {
		log.Fatalf("Failed to deactivate scanner controller: %v", err)
		return
	}

	return
}
