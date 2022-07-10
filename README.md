

# P2PDB

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
<p align="center">
    <a href="https://opensource.org/licenses/Apache-2.0">
        <img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg"
            alt="License"></a>
    <a href="https://godoc.org/github.com/Rock-liyi/p2pdb">
        <img src="https://img.shields.io/badge/godoc-reference-blue.svg"
            alt="GoDoc"></a>
</p>

中文 | [English](./README-EN.md)


### 简介
P2PDB,一个为了web3.0 时代的而产生的去中心化、分布式、点对点数据库、它是传统数据库技术融合区块链技术产生的新一代数据库,P2PDB基于[merker-CRDT](https://research.protocol.ai/blog/2019/a-new-lab-for-resilient-networks-research/PL-TechRep-merkleCRDT-v0.1-Dec30.pdf)论文,使用IPFS-libp2p构建分布式网络和IPFS-pubsub与对等节点同步数据。P2PDB愿景是打造一个工业级别的去中心化式数据库，使P2PDB 成为离线应用程序，去中心化应用程序(dApps)、和边缘计算应用程序数据存储的绝佳选择, P2PDB基于[白皮书](doc/zh-cn/%E7%99%BD%E7%9A%AE%E4%B9%A6.md)实现

P2PDB主要由以下模块构成：

![alt 属性文本](./p2pdb.png)


- `p2pdb-server`: [p2pdb-server](https://github.com/Rock-liyi/p2pdb-server)一个mysql 语法的服务器端, 兼容90%mysql 语法，可使用任一一种mysql 客户端连接,包括PHP、JAVA、GO、RUBY、PYTHON等主流语言的兼容（已实现）

- `p2pdb-log`: 基于[merker-crdt](https://research.protocol.ai/blog/2019/a-new-lab-for-resilient-networks-research/PL-TechRep-merkleCRDT-v0.1-Dec30.pdf)协议实现的不可篡改日志组件（已实现）


- `p2pdb-pubsub`: 基于[libp2p-pubsub](github.com/libp2p/go-libp2p-pubsub)实现的消息广播组件,用于对等节点中数据的主动传播,采用了Gossip流言广播算法（开发中）。


- `p2pdb-consistency`: 基于[crdt](https://github.com/Rock-liyi/p2pdb/blob/main/doc/zh-cn/CRDT%E5%8D%8F%E8%AE%AE.md)，ot等协议用于消息顺序一致性判断的组件,主要用于当事件没有因果关系时（并发）,作为全序判断的模块，集成了常用的顺序判断规则，如内置规则无法满足，你也可以根据该模块规范增加新的协议（开发中）。


- `p2pdb-discovery`: [p2pdb-discovery](https://github.com/Rock-liyi/p2pdb-discovery) 是对等节点的服务发现注册,用于检索对等节点,基于[libp2p](https://github.com/libp2p/go-libp2p)的mdns模块实现（开发中）。


- `p2pdb-store`: [p2pdb-store](https://github.com/Rock-liyi/p2pdb-store) 用于数据实际存储的模块,类似mysql的数据存储一样，提供索引检索,数据的增删改查等,这是一个抽象的模块,目的是将来如果当前的存储引擎无法满足你的存储需求,可以提供更多的DB驱动如clickhouse、postgresql、TDngine等数据库，当前提供内存存储（开发模式使用）以及文件存储（生产使用）(已实现)。


### 核心模块解释

#### p2pdb-log
p2pdb基于[p2pdb-log](https://github.com/Rock-liyi/p2pdb-log)之上实现，p2pdb-log是一种只允许追加写入日志，并且不可窜改。基于操作的无冲突复制数据结构 (CRDT)与Merkle DAG（有向无环图）实现。如果所有 P2PDB 数据库类型都不符合您的需求和/或您需要特定于案例的功能，您可以轻松使用日志模块实现您想要的数据库(开发中)。

#### p2pdb-server
[p2pdb-server](https://github.com/Rock-liyi/p2pdb-server) 可以理解为一个mysql 的服务器端，用于执行mysql的指令,
被p2pdb-server执行的指令都会记录在p2pdb-log中,并广播到所有对等节点,p2pdb-server模拟了mysql协议的实现，因此你可以使用任何一种mysql的客户端进行连接,甚至是编程语言。


### 快速使用
```
git  clone  https://github.com/Rock-liyi/p2pdb.git

```
需要安装golang 1.6及以上环境,默认端口3306,ip 127.0.0.1,可以使用任意一种mysql 客户端进行链接

```
go mod init 
go mod tidy
go run  interface/cli/start.go

```

<!-- 
## 内容列表

- [背景](#背景)
- [目标](#目标)
	- [使用场景](#使用场景)
- [特性](#特性)
	- [数据库特性](#数据库特性)    
- [架构](#架构)
	- [目录分层](#目录分层)
- [使用说明](#使用说明)
	- [安装](#安装)
	- [文档](#理解CDRT协议)
- [徽章](#徽章)
- [示例](#示例)
- [相关仓库](#相关仓库)
- [维护者](#维护者)
- [如何贡献](#如何贡献)
- [使用许可](#使用许可) -->

### 背景
P2PDB最早源于某集团离线收银项目的研发，早期采用了Raft 强一致性的协议+Sqlite 实现一个轻量级的分布式数据库，随着对Raft协议的深入研究，发现Raft协议对于离线场景存在较大缺陷（故障节点数量大于50%，整个集群无法工作）,随后开启了深入研究分布式数据库协议的道路,在一年后发现基于默克尔有向无环图+CRDT实现的逻辑时钟[merker-CRDT](https://research.protocol.ai/blog/2019/a-new-lab-for-resilient-networks-research/PL-TechRep-merkleCRDT-v0.1-Dec30.pdf)可以实现最终一致性的去中心化、分布式、点对点数据库，并且针对离线场景做了充分的支持。


P2PDB 是一个点对点数据库，这意味着每个对等点都有自己的特定数据库实例。数据库在对等点之间自动复制，从而在任何对等点更新时生成数据库的最新视图。也就是说，数据库被拉到客户端。

这意味着每个应用程序都包含他们正在使用的完整数据库。与客户端-服务器模型相比，这反过来又改变了数据建模，客户端-服务器模型通常为所有数据使用一个大数据库：在 P2PDB 中，应该根据该数据的访问权限进行存储、“分区”或“分片”。例如，在类似微博的应用程序中，推文不会保存在数百万用户同时写入的全局“推文”数据库中，而是每个用户都有自己的微博数据库。要订阅别的用户推文，只需关注主题（topic）


—— [跟ipfs的关系](https://www.ipfs.io/)    

> ipfs协议 用于构建分布式低延迟的消息传输网络，而P2PDB 使用ipfs协议实现网络通信、数据传输等能力。

<!-- —— [跟filecoin的区别](https://filecoin.io/)
> P2PDB类似filecoin实现文件交换网络一样，目的是为了实现全球去中心化的数据交换网络。不同的是， P2PDB只接受一段数据流的存储而不是文件，相对filecoin来说，P2PDB更轻量级，数据交换速度更快（数据体积更小），P2PDB可以理解为是一款去中心化存储的轻量级关系型数据库，当然P2PDB也支持非关系性数据库中key=>value 键值对，以及类似mongdb的文档型数据存储格式。 -->



### 目标
 这个数据库的目标是：

1. 一个**Dapp应用数据存储解决方案**
2. 一个**去中心化分布式数据库解决方案**
3. 一个**边缘数据存储解决方案**
4. 一个**离线应用数据存储解决方案**

 使用场景
* 1、离线web应用存储
* 2、多人互动游戏
* 3、多人协作文档
* 4、多人聊天应用
* 5、边缘缓存存储
* 6、Dapp应用存储
* 7、NFT应用存储
* 8、更多.........


### 特性

1. 兼容mysql 90%以上语法支持，你可以使用任何mysql 客户端连接
2. 去中心化，无中心服务器
3. 历史数据一旦生成，无法窜改
4. 数据加密传输，公私钥加解密，授权支持到字段级别
5. 最终一致性，支持强最终一致
6. 故障恢复，定时快照本份，自动同步数据
7. 版本证明，类似区块链算法，保证数据一旦生成永不丢失。



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




### 相关文档
- [文档](https://github.com/Rock-liyi/p2pdb/tree/master/doc)



### 本数据库使用到的部分仓库

- [libp2p](https://github.com/libp2p/go-libp2p) 
- [ipfs](https://github.com/ipfs/go-ipfs)
说明:由于本项目使用到的ipfs,libp2p等代码,引用的代码部分遵循代码相关协议,感谢协议实验室为去中心化网络做出的贡献

### 特别说明
p2pdb 是一个从0到1的项目，也是较早从事去中心化数据库从理论到实践的项目,但是在实现的过程中，我们google了一遍全球关于去中心化数据库的文章及代码实现,可以说p2pdb 是站在巨人肩膀上，是因为有众多优秀的前辈的铺路，才有了p2pdb项目，在这里真诚的向所有研究去中心化数据库前辈，以及IPFS 团队、协议实验室的付出表示感谢，如果有任何引用代码，或者实现思路参考未在项目中注明，请联系我们，我们查证后将非常乐意注明，荣幸之至。

### 维护者

[@Rock](https://github.com/Rock-liyi)
[@Panda](https://github.com/PandaLIU-1111)
[@CbYip](https://github.com/CbYip)

### 如何贡献

非常欢迎你的加入！[提一个 Issue](https://github.com/Rock-liyi/p2pdb) 或者提交一个 Pull Request。

标准 Readme 遵循 [Contributor Covenant](http://contributor-covenant.org/version/1/3/0/) 行为规范。


<!-- 感谢以下参与项目的人： ### 贡献者
<a href="graphs/contributors"><img src="https://opencollective.com/standard-readme/contributors.svg?width=890&button=false" /></a> -->


### 使用许可

Apache License Version 2.0 see http://www.apache.org/licenses/LICENSE-2.0.html












