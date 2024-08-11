package events

// Event событие, передаваемое между издателем и подписчиком
type Event struct {
	Type    string
	Payload interface{}
}

// Subscriber функция, которая будет вызвана при возникновении события
type Subscriber func(event Event)
