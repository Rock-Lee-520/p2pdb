package event

import (
	"bytes"
	"math/rand"
	"testing"
	"time"

	_ "github.com/Rock-liyi/p2pdb/application/event/subscribe" //注册事件监听
	commonEntity "github.com/Rock-liyi/p2pdb/domain/common/entity"
	"github.com/Rock-liyi/p2pdb/domain/common/entity/value_object"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/function"
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

func TestPublishAsyncEvent(t *testing.T) {
	data := randUpString(19)
	//start a subscribe service	 for the application
	PublishAsyncEvent("TestPublishAsyncEvent", data)

	time.Sleep(5 * time.Second)
}

func TestPublishSyncEvent(t *testing.T) {
	//data := "StoreDeleteEvent"
	//data := "StoreDeleteEvent"
	var data = commonEntity.NewData()
	data.SQLStatement = " INSERT INTO p2pdb.test6 (name2,email2,id) VALUES('123','123',123)"
	data.DatabaseName = "test6"
	data.TableName = "test6"
	data.DMLType = "insert"
	PublishSyncEvent(value_object.StoreInsertEvent, function.JsonEncode(data))
}
