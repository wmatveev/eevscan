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
			//"/dev/ttyACM1",
			//"/dev/ttyACM2",
			//"/dev/ttyACM3",
			//"/dev/ttyACM4",
			//"/dev/ttyACM5",
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

func (pc *PortController) StartPortsReading() {

	for _, port := range pc.portNames {
		select {
		case <-pc.QuitChannel:
			log.Println("Stopping ports reading")
			return

		default:
			barcode, err := ReadFromPort(port)
			if err != nil {
				log.Printf("Failed to read from port %s: %v", port, err)
				continue
			}

			log.Printf("--->1 Barcode bytes: %s", string(barcode))
			if barcode != nil {
				log.Printf("--->2 Barcode bytes: %s", string(barcode))
				pc.Barcode <- barcode
			}
		}
	}

	return
}

func ReadFromPort(portName string) ([]byte, error) {
	c := &serial.Config{Name: portName, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}
	defer s.Close()

	buf := make([]byte, 128)
	n, err := s.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf[:n], nil
}
