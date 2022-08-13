package function

import (
	"encoding/json"
	"log"
)

func JsonEncode(data interface{}) []byte {
	result, err := json.Marshal(data)
	if err != nil {
		log.Panic(err)
	}
	return result
}

func JsonDecode(data []byte, value interface{}) {

	err := json.Unmarshal(data, value)
	if err != nil {
		log.Panic(err)
	}
}
