

### 版本管理
Go module遵循语义化版本规范 2.0.0 major.minor.patch
版本号命名如下:
#### 英文版
MAJOR.MINOR.PATCH
#### 中文版
主版本号.次版本号.修订号

#### 版本号递增规则
- 主版本号(MAJOR Version): 出现不兼容的API变化
- 次版本号(MINOR Version): 新增向后兼容的功能
- 修订号(补丁版本号, PATCH Version): 向后兼容的bug



### 对CAP定理的满足条件

- [x] 可用性   
- [x] 分区容错性
- [x] 一致性
    - [] 实时一致性
    - [x] 最终一致性

说明:
P2PDB能做到全局的最终一致性,对于顺序问题，如同狭义相对论中所说，不同角度会观察到不同结果，仅当具有因果关系时，才可完全保证事件顺序，当不同节点出现并发操作事，P2PDB使用p2pdb-consistency处理排序，使得所有节点认可事件顺序，但这个事件顺序并不一定是真实的事件顺序。



### 对ACID的满足条件

- [x] 原子性(atomicity 或称不可分割线)   
- [ ] 一致性
    - [X] 本地实例实时一致性
    - [x] 远程实例最终一致性
- [x] 隔离性(isolation 又称独立性)   
- [x] 持久性(durability)   

P2PDB  的锁主要分为读和写，读锁是可以共存的，但是写锁是互斥的，也就是说， 同时只有一个地方能写入，这也就保证了隔离性，再加上日志，就实现了ACID,具体实现原理可参考[Sqlite](https://www.sqlite.org/atomiccommit.html)文章


