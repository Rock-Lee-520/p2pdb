package function

import (
	uuid "github.com/satori/go.uuid"
)

func GetLocalHost() string {
	return "127.0.0.1"
}

func GetLocalInstanceId() string {
	id := uuid.NewV4()

	return id.String()
}

func GetLocalPeerId() string {
	return "localhost1111"
}

func GetGlobalLogicalClock() int64 {
	return 0
}

func GetUUID() string {
	id := uuid.NewV4()

	return id.String()
}
