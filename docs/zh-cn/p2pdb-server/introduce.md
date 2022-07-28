
**p2pdb-server** 一个mysql 语法的服务器端, 兼容90%mysql 语法，可使用任一一种mysql 客户端连接,包括PHP、JAVA、GO、RUBY、PYTHON等主流语言的兼容（正积极测试中）,需要注意的是,p2pdb-server 只是一个sql查询接收器,本身并不提供真实的数据存储,在p2pdb中 数据存储使用[p2pdb-store](https://github.com/Rock-liyi/p2pdb-store)模块实现


### SQL 语法
p2pdb-server的目标是支持 MySQL 所做的 100% 的语句。我们不断向引擎添加更多功能，但并非所有功能都能支持。
具体支持的语法参考[SUPPORTED](zh-cn/p2pdb-server/SUPPORTED.md)


### 客户端兼容：

我们支持并积极测试某些第三方客户端，以确保它们与 p2pdb-server 之间的兼容性。您可以查看[SUPPORTED_CLIENTS](zh-cn/p2pdb-server/SUPPORTED_CLIENTS.md)文件中受支持的第三方客户端列表 以及有关如何使用它们连接到 p2pdb-server 的一些示例。

### 架构
查看[架构](zh-cn/p2pdb-server/ARCHITECTURE.md)



## 致谢

**p2pdb-server** 最原始的代码来源于{source-d} 这个组织，但是这个仓库fork 来自 go-mysql-server,在这里
真诚的感谢source-d组织以及dolthub组织的贡献，是因为你们的前面的铺垫才有了后来的p2pdb-server

## 大部分代码来源
[go-mysql-server](http://github.com/Rock-liyi/p2pdb-server):遵守Apache License 2.0协议

