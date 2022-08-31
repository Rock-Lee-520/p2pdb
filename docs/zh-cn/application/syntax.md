
与关系数据库交互一致,P2PDB 命令类似于 SQL。命令包括 CREATE、SELECT、INSERT、UPDATE、DELETE 和 DROP,在默认的file 文件存储引擎中,底层使用到了SQLITE 存储引擎,所以某些场景上P2PDB的SQL语法及支持的程度可以跟SQLITE一致。
这些命令基于它们的操作性质可分为以下几种：

### DDL - 数据定义语言
命令|描述|是否支持
-|-|-
CREATE|创建一个新的表，一个表的视图，或者数据库中的其他对象|:white_check_mark:
ALTER|修改数据库中的某个已有的数据库对象，比如一个表|:black_square_button: 
DROP|删除整个表，或者表的视图，或者数据库中的其他对象|:white_check_mark:

### DML - 数据操作语言
命令|描述|是否支持
-|-|-
INSERT|创建一条记录|:white_check_mark:
UPDATE|修改记录|:white_check_mark:
DELETE|删除记录|:white_check_mark:


### DQL - 数据查询语言
命令|描述|是否支持
-|-|-
SELECT|从一个或多个表中检索某些记录|:white_check_mark:



### 支持的SQL语法

P2PDB 数据类型是一个用来指定任何对象的数据类型的属性。P2PDB 中的每一列，每个变量和表达式都有相关的数据类型。

您可以在创建表的同时使用这些数据类型。P2PDB 使用一个更普遍的动态类型系统。在 P2PDB 中，值的数据类型与值本身是相关的，而不是与它的容器相关。

### 支持的类型
每个存储在 SQLite 数据库中的值都具有以下存储类之一：

存储类	描述
NULL	值是一个 NULL 值。
INTEGER	值是一个带符号的整数，根据值的大小存储在 1、2、3、4、6 或 8 字节中。
REAL	值是一个浮点值，存储为 8 字节的 IEEE 浮点数字。
TEXT	值是一个文本字符串，使用数据库编码（UTF-8、UTF-16BE 或 UTF-16LE）存储。
BLOB	值是一个 blob 数据，完全根据它的输入存储。
SQLite 的存储类稍微比数据类型更普遍。INTEGER 存储类，例如，包含 6 种不同的不同长度的整数数据类型。