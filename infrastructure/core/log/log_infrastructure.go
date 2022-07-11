package log

import "sync"

type LogInstance struct {
}

var instance *LogInstance
var _once sync.Once

func GetInstance() *LogInstance {

	_once.Do(func() {
		instance = &LogInstance{}
	})
	return instance
}

func (logInstance *LogInstance) NewLog() {

}
