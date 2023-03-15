package app

import (
	"encoding/json"
	"fmt"
)

type EventRouter struct {
	eventHandler iEventHandler
}

func New(eventHandler iEventHandler) *EventRouter {
	return &EventRouter{
		eventHandler: eventHandler,
	}
}

type iEventHandler interface {
	EmailSendTask(eventDetails string) error
	EmailBounce(eventDetails string) error
	EmailComplaint(eventDetails string) error
}

type Event struct {
	EventName     string `json:"eventType"`
	CorrelationId string `json:"correlationId"`
	EventDetails  string `json:"eventDetails"`
}

func (s *EventRouter) ProcessEvent(eventString string) error {
	event := Event{}
	err := json.Unmarshal([]byte(eventString), &event)
	if err != nil {
		return fmt.Errorf("error parsing event: %s", err.Error())
	}
	if event.EventName == "Bounce" || event.EventName == "Complaint" {
		event.EventDetails = eventString
	}
	switch event.EventName {
	case "email-send-task":
		err = s.eventHandler.EmailSendTask(event.EventDetails)
	case "Bounce":
		err = s.eventHandler.EmailBounce(event.EventDetails)
	case "Complaint":
		err = s.eventHandler.EmailComplaint(event.EventDetails)
	default:
		err = fmt.Errorf("unsupported event name")
	}
	if err != nil {
		return fmt.Errorf("error processing event (%s) => %v", event.EventName, err.Error())
	}
	return nil
}
