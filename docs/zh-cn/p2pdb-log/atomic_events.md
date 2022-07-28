
## 数据结构
```
{
    "id": "163538817972916224", // 原子事件ID
    "resource_id":""//资源id
    "timestamp": 1652788631, // 本机创建时间戳
    "cid": QmcRD4wkPPi6dig81r5sLj9Zm1gDCL4zgpEj9CfuRrGbzF, // 节点cid
    "last_id":"163538817972916224"  //因果关系事件id(上一个关联事件id),为空代表没有因果关系
    "revision": 5808, // 资源版本,单调递增
    "type": "log", // 原子事件自定义类型
    "operations": [ // 自定义操作集
        ...........
    ],
    "properties": { // 自定义属性信息
        ...........
    }
   
}
```

