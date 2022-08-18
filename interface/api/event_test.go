package api

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
	//_ "github.com/Rock-liyi/p2pdb/application/event/subscribe" //注册事件监听
)

func randInt(min int, max int) byte {

	rand.Seed(time.Now().UnixNano())

	return byte(min + rand.Intn(max-min))

}

func randUpString(l int) []byte {

	var result bytes.Buffer

	var temp byte

	for i := 0; i < l; {

		if randInt(65, 91) != temp {

			temp = randInt(65, 91)

			result.WriteByte(temp)

			i++

		}

	}

	return result.Bytes()

}

var eventApi *EventApi

func TestPublishAsyncEvent(t *testing.T) {
	data := randUpString(19)
	//start a subscribe service	 for the application
	eventApi.PublishAsyncEvent("TestPublishAsyncEvent", data)

	time.Sleep(5 * time.Second)
}

func TestPublishSyncEvent(t *testing.T) {
	data := "StoreDeleteEvent"
	eventApi.PublishSyncEvent("StoreDeleteEvent", []byte(data))
}
