# 开始安装

## 环境要求

P2PDB 使用纯golang 语言开发，由于golang的跨平台特性,可以在Mac OS、Liunx 、Windows 等操作系统上运行

## golang 版本要求
golang version 大于等于 1.7

## 安装

使用git 拉取源代码包到本地运行

```
git  clone  https://github.com/Rock-liyi/p2pdb.git

```

> P2PDB  使用 `Mod` 来管理项目的依赖，在使用 P2PDB 之前，需要使用Mod 拉取依赖包
```shell
go mod init 
go mod tidy
```



> 启动一个 P2PDB 实例

```shell
go run  interface/cli/start.go
```

