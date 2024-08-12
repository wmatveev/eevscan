package device

import (
	"github.com/tarm/serial"
	"log"
)

type RS232Controller struct {
	portName []string
}

func NewRS232Controller() (*RS232Controller, error) {

	return &RS232Controller{
		portName: make([]string, 0),
	}, nil
}

func (rs *RS232Controller) Write(data []byte) {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Printf("RS232 OpenPort err:%v", err)
		return
	}
	defer func(s *serial.Port) {
		err := s.Close()
		if err != nil {
		}
	}(s)

	log.Printf("Send to RS232 serial port:%b", data)
	_, err = s.Write(data)
	if err != nil {
		log.Printf("RS232 Write err:%v", err)
		return
	}

	return
}
