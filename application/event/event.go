package event

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Event interface {
	NewEvent(topic string)
	Register(topic string)
	PrintDataEvent(ch string, data DataEvent)
	PublishEvent(topic string, data interface{})
}

type EventFuncs struct {
}

var eb = &EventBus{
	subscribers: map[string]DatachannelSlice{},
}

func (e *EventFuncs) Register(topic string, chanEvent chan DataEvent) {
	//	chanEvent := make(chan DataEvent)

	eb.Subscribe(topic, chanEvent)

}

func (e *EventFuncs) NewEvent(topic string) {
	chanEvent := make(chan DataEvent)

	eb.Subscribe(topic, chanEvent)
}

type Message struct {
	Messagetype string
	Data        []byte
}

func (e *EventFuncs) PrintDataEvent(ch string, data DataEvent) {
	fmt.Printf("Channel: %s; Topic: %s; DataEvent %v\n", ch, data.Topic, data.Data)
}

func (e *EventFuncs) PublishEvent(topic string, message string) {
	for {
		eb.Publish(topic, message)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

type DataEvent struct {
	Data  interface{}
	Topic string
}

type DataChannel chan DataEvent

type DatachannelSlice []DataChannel

type EventBus struct {
	subscribers map[string]DatachannelSlice
	rm          sync.RWMutex
}

func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}

func (eb *EventBus) Publish(topic string, data interface{}) {

	eb.rm.RLock()
	if chans, found := eb.subscribers[topic]; found {

		channels := append(DatachannelSlice{}, chans...)
		go func(data DataEvent, dataChannelSlice DatachannelSlice) {
			for _, ch := range dataChannelSlice {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	eb.rm.RUnlock()
}
