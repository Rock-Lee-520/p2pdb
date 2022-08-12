### 语法说明
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

