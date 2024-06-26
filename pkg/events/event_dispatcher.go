package events

import (
	"errors"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (e *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if e.Has(eventName, handler) {
		return ErrHandlerAlreadyRegistered
	}

	e.handlers[eventName] = append(e.handlers[eventName], handler)

	return nil
}

func (e *EventDispatcher) Clear() {
	e.handlers = make(map[string][]EventHandlerInterface)
}

func (e *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := e.handlers[eventName]; ok {
		for _, registeredHandler := range e.handlers[eventName] {
			if registeredHandler == handler {
				return true
			}
		}
	}
	return false
}

func (e *EventDispatcher) Dispatch(event EventInterface) {
	eventName := event.GetName()
	if handlers, ok := e.handlers[eventName]; ok {
		for _, handler := range handlers {
			handler.Handle(event)
		}
	}
}

// func (e *EventDispatcher) Dispatch(event EventInterface, dispatchedEventCh chan EventHandlerInterface) {
// 	eventName := event.GetName()
// 	if handlers, ok := e.handlers[eventName]; ok {
// 		for _, handler := range handlers {
// 			go func() {
// 				handler.Handle(event)
// 				dispatchedEventCh <- handler
// 			}()
// 		}
// 	}
// }

func (e *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) {
	if e.Has(eventName, handler) {
		for i, h := range e.handlers[eventName] {
			if h == handler {
				e.handlers[eventName] = append(e.handlers[eventName][:i], e.handlers[eventName][i+1:]...)
				return
			}
		}
	}
}
