package event

import (
	"fmt"
	"sync"
)

type EventFuncs struct {
}

var eb = &EventBus{
	subscribers: map[string]DatachannelSlice{},
}

func (e *EventFuncs) RegisterAsyncEvent(topic string, chanEvent chan DataEvent) {
	//	chanEvent := make(chan DataEvent)

	eb.SubscribeAsyncEvent(topic, chanEvent)

}

func (e *EventFuncs) NewAsyncEvent(topic string) {
	chanEvent := make(chan DataEvent)

	eb.SubscribeAsyncEvent(topic, chanEvent)
}

func (e *EventFuncs) PrintDataAsyncEvent(topic string, data DataEvent) {
	fmt.Printf("Channel: %s; Topic: %s; DataEvent %v\n", topic, data.Topic, data.Data)
}

func (e *EventFuncs) PublishAsyncEvent(topic string, message interface{}) {
	//for {

	eb.PublishAsyncEvent(topic, message)
	//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	//}
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

func (eb *EventBus) SubscribeAsyncEvent(topic string, ch DataChannel) {
	eb.rm.Lock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}

func (eb *EventBus) PublishAsyncEvent(topic string, data interface{}) {

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
