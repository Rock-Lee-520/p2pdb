
<!-- [![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme) -->
<p align="center">
  <a href="https://opensource.org/licenses/Apache-2.0">
        <img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg"
            alt="License"></a>
    <a href="https://github.com/RichardLitt/standard-readme">
        <img src="https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square"
            alt="License"></a>
    <a href="https://godoc.org/github.com/Rock-liyi/p2pdb">
        <img src="https://img.shields.io/badge/godoc-reference-blue.svg"
            alt="GoDoc"></a>
                <a href="https://godoc.org/github.com/Rock-liyi/p2pdb">
        <img src="https://img.shields.io/badge/platform-Win32%20|%20GNU/Linux%20|%20macOS%20|%20FreeBSD%20-gold"
            alt="platform"></a>
</p>

# P2PDB


中文 | [English](./README.md)


### 简介
P2PDB,一个为了web3而生的去中心化、分布式、点对点数据库、它是传统数据库技术与区块链技术的结合,最终数据一致性算法基于[merker-CRDT](https://research.protocol.ai/blog/2019/a-new-lab-for-resilient-networks-research/PL-TechRep-merkleCRDT-v0.1-Dec30.pdf)论文实现,它使用IPFS-libp2p构建去中心化网络和IPFS-pubsub与对等节点同步数据。P2PDB愿景是打造一个工业级别的去中心化式数据库，使P2PDB 成为离线应用程序，去中心化应用程序(dApps)、和区块链应用程序数据存储的绝佳选择, P2PDB基于[白皮书](docs/zh-cn/whitebook.md)实现


### 使用教程
[使用教程](https://rock-liyi.github.io/p2pdb/#/zh-cn/)


### 核心模块

- :white_check_mark: `p2pdb-server`: [p2pdb-server](https://github.com/Rock-liyi/p2pdb-server)一个mysql 语法的服务器端, 兼容90%mysql 语法，可使用任一一种mysql 客户端连接,包括PHP、JAVA、GO、RUBY、PYTHON等主流语言的兼容 

- :white_check_mark: `p2pdb-log`: 基于[merker-crdt](https://research.protocol.ai/blog/2019/a-new-lab-for-resilient-networks-research/PL-TechRep-merkleCRDT-v0.1-Dec30.pdf)协议实现的不可篡改日志组件，p2pdb-log是一种只允许追加写入日志，并且不可窜改。基于操作的无冲突复制数据结构 (CRDT)与Merkle DAG（有向无环图）实现。如果所有 P2PDB 数据库类型都不符合您的需求和/或您需要特定于案例的功能，您可以轻松使用日志模块实现您想要的数据库

- :white_check_mark:  `p2pdb-pubsub`: 基于[libp2p-pubsub](github.com/libp2p/go-libp2p-pubsub)实现的消息广播组件,用于对等节点中数据的主动传播,采用了Gossip流言广播算法,内置了一个消息队列,可通过webhook的方式监听数据的变动,支持Websocket协议链接,将来计划对接到Kafka、RabbitMq、Nsq、Nats等主流的消息队列,实现更广泛的数据消费应用，同时它也是一个具备扩展性的模块，可以根据该模块提供的接口，进行二次开发，以满足更多的业务场景


- :black_square_button:  `p2pdb-consistency`: 基于[crdt](https://github.com/Rock-liyi/p2pdb/blob/main/doc/zh-cn/CRDT%E5%8D%8F%E8%AE%AE.md)，ot等协议用于消息顺序一致性判断的组件,主要用于当事件没有因果关系时（并发）,作为全序判断的模块，集成了常用的顺序判断规则，如内置规则无法满足，你也可以根据该模块规范增加新的协议。


- :white_check_mark:  `p2pdb-discovery`: [p2pdb-discovery](https://github.com/Rock-liyi/p2pdb-discovery) 是对等节点的服务发现注册,用于检索对等节点,基于[libp2p](https://github.com/libp2p/go-libp2p)的mdns模块实现。


- :white_check_mark:  `p2pdb-store`: [p2pdb-store](https://github.com/Rock-liyi/p2pdb-store) 用于数据实际存储的模块,类似mysql的数据存储一样，提供索引检索,数据的增删改查等,这是一个抽象的模块,目的是将来如果当前的存储引擎无法满足你的存储需求,可以提供更多的DB驱动如clickhouse、postgresql、TDngine等数据库，当前提供内存存储（开发模式使用）以及文件存储（生产使用）。


<!-- - :black_square_button:  `p2pdb-cdc`: 基于数据库事务日志Change Data Capture(CDC)技术，作为一种更为优雅和先进的实时数据同步方案，广泛用于增量数据集成中，用于进一步扩展消息的使用方式,内置了一个消息队列,可通过webhook的方式监听数据的变动,支持Websocket协议链接,将来计划对接到Kafka、RabbitMq、Nsq、Nats等主流的消息队列,实现更广泛的数据消费应用，同时它也是一个具备扩展性的模块，可以根据该模块提供的接口，进行二次开发，以满足更多的业务场景。 -->


—— [跟ipfs的关系](https://www.ipfs.io/)    

> ipfs协议 用于构建分布式低延迟的消息传输网络，而P2PDB 使用ipfs协议实现网络通信、数据传输等能力。



### 目标
 这个数据库的目标是：

1. 一个**Dapp应用数据存储解决方案**
2. 一个**去中心化分布式数据库解决方案**
3. 一个**边缘数据存储解决方案**
4. 一个**离线应用数据存储解决方案**


### 特性

1. 兼容mysql 通用SQL语法，可以使用任何mysql 客户端连接
2. 去中心化，无中心服务器
3. 历史数据一旦生成，无法窜改
4. 数据加密传输，公私钥加解密，授权支持到字段级别
5. 最终一致性，支持强最终一致
6. 故障恢复，定时快照本份，自动同步数据
7. 版本证明，类似区块链算法，保证数据一旦生成永不丢失。


### 使用场景
* 1、离线web应用存储
* 2、多人互动游戏
* 3、多人协作文档
* 4、多人聊天应用
* 5、边缘缓存存储
* 6、Dapp应用存储
* 7、NFT应用存储
* 8、更多.........







### 本数据库使用到的部分仓库的关联引用组织

- [libp2p](https://github.com/libp2p) 
- [ipfs](https://github.com/ipfs)
- [dolthub](https://github.com/dolthub)
- [berty](https://github.com/berty/go-ipfs-log)

说明:由于本项目使用到的ipfs,libp2p等代码,引用的代码部分遵循代码相关协议,感谢协议实验室为web3.0做出的贡献,其余在独立模块中引用的代码，或者使用代码均在独立模块的LICENSE包含引用说明


### 引用说明
p2pdb是一个从0到1的项目，也是较早研究去中心化数据库领域，并且从理论到实践的项目,但是在实现的过程中，我们google了一遍全球关于去中心化数据库的文章及代码实现,可以说p2pdb 是站在巨人肩膀上，是因为有众多优秀的前辈的铺路，才有了p2pdb项目，在这里真诚的向所有研究去中心化数据库前辈，以及IPFS 团队、协议实验室的付出表示感谢，如果有任何引用代码，或者实现思路未在项目中注明来源，请联系我们，我们查证后将非常乐意注明，荣幸之至。


### 维护说明

维护人员：[@Rock](https://github.com/Rock-liyi) [@Panda](https://github.com/PandaLIU-1111) [@CbYip](https://github.com/CbYip)  [@her-cat](https://github.com/her-cat)

当前维护P2PDB项目的是`kkguan`这个组织,由于某些原因考虑,仓库暂为private状态

注意：项目仍然属于快速迭代阶段，意味着大部分的API都有被重构的风险, 当前可作为学习跟研究使用，不建议运用在生产上,开发团队会尽最大的努力去推进1.0.0版本的发布，但是我们无法预估什么时候发布稳定版本。

### 如何贡献

非常欢迎你的加入！[提一个 Issue](https://github.com/Rock-liyi/p2pdb/issues) 或者提交一个 Pull Request。

标准 Readme 遵循 [Contributor Covenant](http://contributor-covenant.org/version/1/3/0/) 行为规范。


<!-- 感谢以下参与项目的人： ### 贡献者
<a href="graphs/contributors"><img src="https://opencollective.com/standard-readme/contributors.svg?width=890&button=false" /></a> -->


### 使用许可

Apache License Version 2.0 see http://www.apache.org/licenses/LICENSE-2.0.html


## Contributors ✨
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

Thanks goes to these wonderful people :

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/her-cat"><img src="https://avatars.githubusercontent.com/u/18332628?v=4?s=100" width="100px;" alt=""/><br /><sub><b>她和她的猫</b></sub></a><br /><a href="#infra-her-cat" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="https://github.com/Rock-liyi/p2pdb/commits?author=her-cat" title="Tests">⚠️</a> <a href="https://github.com/Rock-liyi/p2pdb/commits?author=her-cat" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/Rock-liyi"><img src="https://avatars.githubusercontent.com/u/35133768?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Rock</b></sub></a><br /><a href="#infra-Rock-liyi" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="https://github.com/Rock-liyi/p2pdb/commits?author=Rock-liyi" title="Tests">⚠️</a> <a href="https://github.com/Rock-liyi/p2pdb/commits?author=Rock-liyi" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/PandaLIU-1111"><img src="https://avatars.githubusercontent.com/u/26201936?v=4?s=100" width="100px;" alt=""/><br /><sub><b>pandaLIU</b></sub></a><br /><a href="#infra-PandaLIU-1111" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="https://github.com/Rock-liyi/p2pdb/commits?author=PandaLIU-1111" title="Tests">⚠️</a> <a href="https://github.com/Rock-liyi/p2pdb/commits?author=PandaLIU-1111" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/CbYip"><img src="https://avatars.githubusercontent.com/u/31086265?v=4?s=100" width="100px;" alt=""/><br /><sub><b>CbYip</b></sub></a><br /><a href="#infra-CbYip" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="https://github.com/Rock-liyi/p2pdb/commits?author=CbYip" title="Tests">⚠️</a> <a href="https://github.com/Rock-liyi/p2pdb/commits?author=CbYip" title="Code">💻</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!