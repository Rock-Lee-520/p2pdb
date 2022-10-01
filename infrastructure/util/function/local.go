package function

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
	uuid "github.com/satori/go.uuid"
)

var node *snowflake.Node

func GetSnowflakeId() int64 {
	var st time.Time

	// 格式化 1月2号下午3时4分5秒  2006年
	st, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	if err != nil {
		fmt.Println(err)
		return 1
	}

	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return 2
	}

	return node.Generate().Int64()
}

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
