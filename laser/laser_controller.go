package laser

import (
	"log"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
	"time"
)

type LaserController struct {
	Address    uint16
	Bus        i2c.BusCloser
	Device     i2c.Dev
	PinChanges chan bool
}

func NewLaserController(address uint16) (*LaserController, error) {
	if _, err := host.Init(); err != nil {
		return nil, err
	}

	bus, err := i2creg.Open("1")
	if err != nil {
		return nil, err
	}

	dev := i2c.Dev{Bus: bus, Addr: 0x21}

	return &LaserController{
		Address:    address,
		Bus:        bus,
		Device:     dev,
		PinChanges: make(chan bool),
	}, nil
}

func (lc *LaserController) StartPinsPolling() {
	var readBuf [1]byte
	var lastState bool

	for {
		if err := lc.Device.Tx(nil, readBuf[:]); err != nil {
			log.Printf("Failed to read from device: %v\n", err)
			break
		}

		currentState := readBuf[0]&0x01 == 0x01

		if currentState != lastState {
			lc.PinChanges <- currentState
			lastState = currentState
		}

		time.Sleep(1 * time.Second)
	}
}
