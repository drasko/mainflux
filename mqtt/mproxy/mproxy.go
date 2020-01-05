package mproxy

import (
	"fmt"
	"strings"

	"github.com/mainflux/mainflux"
	"github.com/mainflux/mainflux/logger"
	"github.com/mainflux/mproxy/pkg/events"
	opentracing "github.com/opentracing/opentracing-go"
)

var _ events.Event = (*Event)(nil)

// Event implements events.Event interface
type Event struct {
	tc     mainflux.ThingsServiceClient
	mp     mainflux.MessagePublisher
	tracer opentracing.Tracer
	logger logger.Logger
}

// New creates new Event entity
func New(tc mainflux.ThingsServiceClient, mp mainflux.MessagePublisher, logger logger.Logger, tracer opentracing.Tracer) *Event {
	return &Event{
		tc:     tc,
		mp:     mp,
		tracer: tracer,
		logger: logger,
	}
}

// AuthRegister is called on device connection,
// prior forwarding to the MQTT broker
func (e *Event) AuthRegister(clientID, username *string, password *[]byte) error {
	e.logger.Info(fmt.Sprintf("AuthRegister() - clientID: %s, username: %s, password: %s", *clientID, *username, string(*password)))
	return nil
}

// AuthPublish is called on device publish,
// prior forwarding to the MQTT broker
func (e *Event) AuthPublish(clientID string, topic *string) error {
	e.logger.Info(fmt.Sprintf("AuthPublish() - clientID: %s, topic: %s", clientID, *topic))
	return nil
}

// AuthSubscribe is called on device publish,
// prior forwarding to the MQTT broker
func (e *Event) AuthSubscribe(clientID string, topics *[]string) error {

	e.logger.Info(fmt.Sprintf("AuthSubscribe() - clientID: %s, topics: %s", clientID, strings.Join(*topics, ",")))
	return nil
}

// Register - after client sucesfully connected
func (e *Event) Register(clientID string) {
	e.logger.Info(fmt.Sprintf("Register() - clientID: %s", clientID))
}

// Publish - after client sucesfully published
func (e *Event) Publish(clientID, topic string) {
	e.logger.Info(fmt.Sprintf("Publish() - clientID: %s, topic: %s", clientID, topic))
}

// Subscribe - after client sucesfully subscribed
func (e *Event) Subscribe(clientID string, topics []string) {
	e.logger.Info(fmt.Sprintf("Subscribe() - clientID: %s, topics: %s", clientID, strings.Join(topics, ",")))
}

// Unubscribe - after client unsubscribed
func (e *Event) Unubscribe(clientID string, topics []string) {

	e.logger.Info(fmt.Sprintf("Unubscribe() - clientID: %s, topics: %s", clientID, strings.Join(topics, ",")))
}
