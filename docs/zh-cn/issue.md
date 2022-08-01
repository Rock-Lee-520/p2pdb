

#### P2PDB 对CAP定理的满足条件

- [x] 可用性   
- [x] 分区容错性
- [ ] 一致性
    - [ ] 实时一致性
    - [x] 最终一致性

说明:
P2PDB能做到全局的最终一致性,对于顺序问题，如同狭义相对论中所说，不同角度会观察到不同结果，仅当具有因果关系时，才可完全保证事件顺序，当不同节点出现并发操作事，P2PDB采用Merkel逻辑时钟处理排序，使得所有节点认可事件顺序，但这个事件顺序并不一定是真实的发送的事件顺序。
