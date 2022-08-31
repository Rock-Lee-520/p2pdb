
### 路线图
版本|当前状态|预计发布时间
-|-|-
1.0-beta|开发中|2022年12月30日|
2.0|计划中|2023年12月30日|
3.0|调研中|2024年12月30日|



### 1.0-beta版本

#### 背景
为了实现基于SQL语句的去中心化关系型数据库,支持常规非金融级别的去中心化业务场景,如去中心化聊天、邮件、多人协作、多人游戏等系统。


#### 目标
实现白皮书中约定的主要功能,达到MVP最小可用标准
#### 实现思路
查看[白皮书](zh-cn/whitebook.md)
#### 注意事项
1.0-beta 版本包含了P2PDB 大部分核心功能, 但是缺乏安全设计及稳定性测试,仍然不建议使用在商业项目上, 欢迎对P2PDB感兴趣的同学下载试用,并在github上反馈问题,开发团队将会积极回复和解决问题。

### 2.0版本

#### 背景
P2PDB的本质是一个去中心化的关系型数据库,需要不同对等节点进行数据的同步及备份,节点之间通过写入数据时设定的规则进行排序来达成全序共识,所以数据都是非加密且共享的,虽然数据的透明度增加了,却带来了数据的隐私问题,部分数据提供方可能并不希望自己的数据全部公开,比如用户身份、订单金额、合同等比较敏感的数据，这不仅包括个人隐私数据，还包括金融或者供应链系统的各种数据。因此，为了在更广阔的使用P2PDB运用在各种各样的去中心化场景,需要解决数据的隐私保护问题。
#### 目标
* 实现隐私保护和数据安全控制
* 单测覆盖90%的使用场景,搭建CI实现自动化回归测试
* 完成性能测试,输出性能测试报告
* 达到生产可用级别,并孵化一个以上商业案例

#### 实现思路
针对这个目标，近年来有不少的成果。其中,一种被称为零知识证明（zero knowledge proof）的加密技术被广泛的研究及运用,用户可以只通过传输是定义的信息来隐藏真实的数据信息，它能够在不向验证者提供任何有用信息的情况下，使验证者相信某个论断是正确的。

### 3.0版本
#### 背景
在1.0及2.0版本中，P2PDB 基于CRDT等算法实现了最终一致性的共识协议,满足了非交易型的大部分业务场景,但是在一些金融级别的业务场景，仍然无法很好的支撑,所以需要一种强一致性的共识算法,主流共识协议中,有Paxos、Raft、Pow、Pos等协议,由于Paxos跟Raft 在去中心化场景中无法解决恶意节点作恶(拜占庭容错问题), 所以P2PDB主要采用对接Pow和Pos的公链实现强一致性共识算法。

#### 目标
实现支持金融级别业务的强一致性共识协议

#### 实现思路
对接以太坊等主流公链,将P2PDB 数据库事务翻译成智能合约代码,使得执行事务通过执行智能合约的方式在公链上完成事务记录