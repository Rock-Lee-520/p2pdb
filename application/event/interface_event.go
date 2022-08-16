package event

type AsyncEvent interface {
	NewAsyncEvent(topic string)
	RegisterAsyncEvent(topic string)
	PrintDataAsyncEvent(topic string, data DataEvent)
	PublishAsyncEvent(topic string, data interface{})
}

type SyncEvent interface {
	OnEvent(param interface{})
	GlobalSyncEvent(param interface{})
	PublishSyncEvent(name string, message Message)
	RegisterSyncEvent(name string, callback func(Message))
}

type Message struct {
	Type string
	Data []byte
}
