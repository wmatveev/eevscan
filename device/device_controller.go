package device

import (
	"log"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

type DeviceController struct {
	deviceAddr uint16
	pinAddr    uint16
	bus        i2c.BusCloser
	device     i2c.Dev
}

func NewDeviceController(deviceAddr uint16, pinAddr uint16) (*DeviceController, error) {
	if _, err := host.Init(); err != nil {
		return nil, err
	}

	bus, err := i2creg.Open("1")
	if err != nil {
		return nil, err
	}

	dev := i2c.Dev{Bus: bus, Addr: deviceAddr}

	return &DeviceController{
		deviceAddr: deviceAddr,
		pinAddr:    pinAddr,
		bus:        bus,
		device:     dev,
	}, nil
}

func (dc *DeviceController) ReadingFromDevice() (uint8, error) {
	var readData [1]byte

	if err := dc.device.Tx(nil, readData[:]); err != nil {
		log.Printf("Failed to read from device: %v\n", err)
		return 0, err
	}

	return readData[0], nil
}

func (dc *DeviceController) WriteToDevice(pinState uint16) error {
	writeData := []byte{byte(pinState)}

	if err := dc.device.Tx(writeData, nil); err != nil {
		log.Fatalf("Failed to write to the device: %v", err)
		return err
	}

	return nil
}
