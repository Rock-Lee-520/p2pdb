

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



### 简介
P2PDB,一个为了web3.0 时代的而生的去中心化、分布式、点对点数据库、它是传统数据库技术与区块链技术的结合,P2PDB基于[merker-CRDT](https://research.protocol.ai/blog/2019/a-new-lab-for-resilient-networks-research/PL-TechRep-merkleCRDT-v0.1-Dec30.pdf)论文,使用IPFS-libp2p构建去中心化网络和IPFS-pubsub与对等节点同步数据。P2PDB愿景是打造一个工业级别的去中心化式数据库，使P2PDB 成为离线应用程序，去中心化应用程序(dApps)、和边缘计算应用程序数据存储的绝佳选择, P2PDB基于[白皮书](zh-cn/whitebook.md)实现



#### P2PDB核心流程：

![alt 属性文本](https://github.com/Rock-liyi/p2pdb/raw/master/p2pdb_core_flowchart.png)


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

1. 兼容mysql 通用SQL语法，你可以使用任何mysql 客户端连接
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



### 引用说明
p2pdb是一个从0到1的项目，也是较早研究去中心化数据库领域，并且从理论到实践的项目,但是在实现的过程中，我们google了一遍全球关于去中心化数据库的文章及代码实现,可以说p2pdb 是站在巨人肩膀上，是因为有众多优秀的前辈的铺路，才有了p2pdb项目，在这里真诚的向所有研究去中心化数据库前辈，以及IPFS 团队、协议实验室的付出表示感谢，如果有任何引用代码，或者实现思路未在项目中注明来源，请联系我们，我们查证后将非常乐意注明，荣幸之至。



### 引用仓库关联组织

- [libp2p](https://github.com/libp2p) 
- [ipfs](https://github.com/ipfs)
- [dolthub](https://github.com/dolthub)
- [berty](https://github.com/berty/go-ipfs-log)

说明:由于本项目使用到的ipfs,libp2p等代码,引用的代码部分遵循代码相关协议,感谢协议实验室为web3.0做出的贡献,其余在独立模块中引用的代码，或者使用代码均在独立模块的LICENSE包含引用说明



### 维护说明

维护人员：[@Rock](https://github.com/Rock-liyi) [@Panda](https://github.com/PandaLIU-1111) [@CbYip](https://github.com/CbYip) [@her-cat](https://github.com/her-cat)

当前维护P2PDB项目的是`kkguan`这个组织,由于某些原因考虑,仓库暂为private状态

注意：项目仍然属于快速迭代阶段，意味着大部分的API都有被重构的风险, 当前可作为学习跟研究使用，不建议运用在生产上,开发团队会尽最大的努力去推进1.0.0版本的发布，但是我们无法预估什么时候发布稳定版本。

### 如何贡献

非常欢迎你的加入！[提一个 Issue](https://github.com/Rock-liyi/p2pdb/issues) 或者提交一个 Pull Request。

标准 Readme 遵循 [Contributor Covenant](http://contributor-covenant.org/version/1/3/0/) 行为规范。


<!-- 感谢以下参与项目的人： ### 贡献者
<a href="graphs/contributors"><img src="https://opencollective.com/standard-readme/contributors.svg?width=890&button=false" /></a> -->


### 使用许可

Apache License Version 2.0 see http://www.apache.org/licenses/LICENSE-2.0.html












