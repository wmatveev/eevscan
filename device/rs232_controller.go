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

func (rs *RS232Controller) Write(data interface{}) {
	var byteData []byte

	// Определяем тип данных и выполняем преобразование
	switch v := data.(type) {
	case string:
		byteData = []byte(v)
	case []byte:
		byteData = v
	default:
		log.Printf("Unsupported data type")
		return
	}

	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Printf("RS232 OpenPort err:%v", err)
		return
	}
	defer func(s *serial.Port) {
		err := s.Close()
		if err != nil {
			log.Printf("RS232 Close err:%v", err)
		}
	}(s)

	// Логируем данные, которые отправляем в виде строки
	log.Printf("Send to RS232 serial port: %s", string(byteData))

	// Отправляем данные на последовательный порт
	_, err = s.Write(byteData)
	if err != nil {
		log.Printf("RS232 Write err:%v", err)
		return
	}
}
