
#### 日志数据结构
```json
{
    "nodeId":"",//节点id
    "instanceId":"",//数据库实例id
    "logicalClock":1,
    "peerId":"",
    "data":{
        "nodeType":"log",//节点类型
        "requestId":"",//请求id 用于幂等
        "logicalClock":12,//逻辑时钟值
        "lastNodeId":"",//上一个节点id
        "databaseId":"",//数据库id
        "tableId":"",//表id
        "opreation":"i",//i、u、r、d
        "properties":
        [
           "before": {//update 数据
            "id":12323,
            "name":"Anne",
             },
            "after":{//update 数据
                "id":12323,
                "name":"Anne",
            }
            ],
        "sql":["update set"]//原始sql
    }
    ,
    "links":[""]
 
   
}
```

实体映射

```golang

type DMLChangeType struct {
	INSERT string
	UPDATE string
	DELETE string
}

type DDLChangeType struct {
	CREATE          string
	ALTER           string
	DROP            string
	DDL_ACTION_TYPE DDL_ACTION_TYPE
}

type DDL_ACTION_TYPE struct {
	DATABASE string
	TABLE    string
}

type ChangeRowData interface {
	getDDLChangeType() *DDLChangeType
}
```

每条 ChangeRowData 都有一个元数据 DMLChangeType 4 种类型， 分别是插入(INSERT)、更新前镜像 (UPDATE_BEFORE)、更新后镜像 (UPDATE_AFTER)、删除 (DELETE)，这四种类型和数据库里面的 binlog 概念保持一致。

op 字段的取值也有四种，分别是 i、u、d、r，各自对应 insert、update、delete、read。对于代表更新操作的u，其数据部分同时包含了前镜像 (before) 和后镜像 (after)。
