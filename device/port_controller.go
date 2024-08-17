package device

import (
	"github.com/tarm/serial"
	"log"
)

type PortController struct {
	portNames   []string
	Barcode     chan []byte
	QuitChannel chan struct{}
}

func NewPortController() (*PortController, error) {

	return &PortController{
		portNames: []string{
			"/dev/ttyACM0",
			"/dev/ttyACM1",
			"/dev/ttyACM2",
			"/dev/ttyACM3",
			"/dev/ttyACM4",
			"/dev/ttyACM5",
		},
		Barcode:     make(chan []byte),
		QuitChannel: make(chan struct{}),
	}, nil
}

func (pc *PortController) RestartPortsReading() {
	close(pc.QuitChannel)
	pc.QuitChannel = make(chan struct{})
	go pc.StartPortsReading()
}

func (pc *PortController) StartPortsReading() []byte {

	for {
		for i := 0; i < len(pc.portNames); i++ {

			log.Println(pc.portNames[i])

			barcode, err := ReadFromPort(pc.portNames[i])
			if err != nil {
				log.Printf("Failed to read from port %s: %v", "/dev/ttyACM0", err)
				continue
			}

			if barcode != nil {
				log.Printf("Barcode bytes: %s", string(barcode))
				return barcode
			}
		}
	}

	return nil
}

func ReadFromPort(portName string) ([]byte, error) {

	config := &serial.Config{
		Name:        portName,
		Baud:        9600,
		Size:        8,
		Parity:      serial.ParityNone,
		StopBits:    serial.Stop1,
		ReadTimeout: 0,
	}

	port, err := serial.OpenPort(config)
	if err != nil {
		log.Printf("Failed to open port: %s", err)
		return nil, err
	}

	defer port.Close()

	buf := make([]byte, 128)
	n, err := port.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf[:n], nil

	//c := &serial.Config{Name: portName, Baud: 9600, ReadTimeout: time.Millisecond * 10}
	//s, err := serial.OpenPort(c)
	//if err != nil {
	//	return nil, err
	//}
	//defer func(s *serial.Port) {
	//	err := s.Close()
	//	if err != nil {
	//	}
	//}(s)
	//
	//buf := make([]byte, 128)
	//n, err := s.Read(buf)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return buf[:n], nil
}
