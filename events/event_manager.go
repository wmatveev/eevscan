package events

// EventManager управляет подписками и публикацией событий
type EventManager struct {
	subscribers map[string][]Subscriber
}

func NewEventManager() *EventManager {
	return &EventManager{
		subscribers: make(map[string][]Subscriber),
	}
}

// Subscribe добавляет подписчика на конкретный тип события
func (em *EventManager) Subscribe(eventType string, subscriber Subscriber) {
	em.subscribers[eventType] = append(em.subscribers[eventType], subscriber)
}

// Publish публикует событие для всех подписчиков, подписанных на данный тип события
func (em *EventManager) Publish(event Event) {
	if subs, found := em.subscribers[event.Type]; found {
		for _, subscriber := range subs {
			subscriber(event)
		}
	}
}
