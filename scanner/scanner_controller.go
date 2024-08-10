package scanner

//type Controller struct {
//	deviceAddr	uint16
//	pinAddr		uint16
//	bus			i2c.BusCloser
//	device		i2c.Dev
//}
//
//func NewScannerController(devAddress uint16) *Controller {
//	if _, err := host.Init(); err != nil {
//		return nil, err
//	}
//
//	bus, err := i2creg.Open("1")
//	if err != nil {
//		return nil, err
//	}
//
//	dev := i2c.Dev{Bus: bus, Addr: devAddress}
//
//	return &Controller{
//		deviceAddr: devAddress,
//		pinAddr:    bus,
//		device:     dev,
//		PinChanges: make(chan bool),
//	}, nil
//}
