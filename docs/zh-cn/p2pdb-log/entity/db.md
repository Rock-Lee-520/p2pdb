
### DB文件目录说明
默认地址
```
p2pdb-server/data

```
默认文件
p2pdb_instance   存放实例信息目录
    config  实例配置信息表
    table 实例表信息表
    log 存储表日志目录
        {db}/{table}/node 节点信息
        {db}/{table}/object 节点数据对象
        {db}/{table}/link   节点连接

db        用户自定义数据库
    table 用户自定义数据表