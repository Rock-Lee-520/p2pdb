
### 目录分层设计
遵循DDD领域驱动分层架构
```
interface 接口层
----cli 命令行执行工具
application 应用层
----event 事件
--------publish 发布事件
--------subscribe 订阅事件
domain 领域层， 核心逻辑
----common 公共领域
----server 服务端领域
infrastructure	基础设施层
----core 核心模块引用
--------server p2pdb-server 模块引用
--------log    p2pdb-log 模块引用
----util  公共工具，如日志
--------log
data 文件数据存储目录
```
