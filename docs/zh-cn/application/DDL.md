
> ### 连接数据库


```bash
mysql -u root -p -h 127.0.0.1
Enter password:
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 2

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql>
```
#### 解释
以上命令演示使用mysql-client 进行登陆后的响应结果



> ### 显示所有数据库

```sql 
SHOW DATABSES
```
#### 解释
该命令会显示当前连接下的所有数据库,包含一个在内存中维护元数据信息的information_schema库



#### 例子
```sql
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| p2pdb              |
+--------------------+
2 rows in set (0.01 sec)

```

> ### 选择数据库

```SQL 
USE DATABASE_NAME
```

#### 解释
选择其中一个数据库进行操作

#### 例子

```bash
mysql> use p2pdb;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
```

> ### 创建新数据库(暂不支持)
```
CREATE DATABASE 数据库名;
```
#### 解释
通过自定义名字创建一个新的数据库

#### 例子
```bash

mysql> create DATABASE p2pdb_2;

```

> ### 删除数据库(暂不支持)
```
DROP DATABASE 数据库名;
```
#### 解释
删除一个数据库,注意`内置数据库information_schema 不可被删除`

#### 例子
```bash

mysql> drop database p2pdb_2;

```



> ### 创建表
#### 基本语法:
 <!-- -- column1 datatype  PRIMARY KEY(one or more columns), -->
```sql
CREATE TABLE database_name.table_name(
  
   column1 datatype,
   column2 datatype,
   .....
   columnN datatype,
);

```

#### 解释:
CREATE TABLE 是告诉数据库系统创建一个新表的关键字，CREATE TABLE 语句后跟着表的唯一的名称或标识。您也可以选择指定带有 table_name 的 database_name


#### 例子:

```bash
mysql> CREATE TABLE test (id integer , keys TEXT, values TEXT);
Query OK, 0 rows affected (0.03 sec)
```

> ### 删除表
#### 基本语法:
```sql
DROP TABLE database_name.table_name;

```

#### 解释:
DROP TABLE 语句的基本语法如上，您可以选择指定带有表名的数据库名称，如上所示：

#### 例子:
先确认一下库是否存在
```bash
mysql> show tables;
+--------+
| Table  |
+--------+
| test  |
+--------+
1 rows in set (0.00 sec)
```

这意味着 test 表已存在数据库中，接下来让我们把它从数据库中删除，如下：
```bash
mysql> drop table test;
Query OK, 0 rows affected (0.03 sec)

mysql> show tables;
+--------+
| Table  |
+--------+
1 rows in set (0.01 sec)
```

> ### 修改表(暂不支持)


### 显示建表语句
#### 基本语法:
```sql
SHOW TABLE CREATE.TABLE_NAME;

```

#### 解释:
SHOW TABLE 可显示表结构信息：

#### 例子:

```bash
mysql> show create table test;
+-------+----------------------------------------------------------------------------------------------------------+
| Table | Create Table                                                                                             |
+-------+----------------------------------------------------------------------------------------------------------+
| test  | CREATE TABLE `test` (
  `id` int,
  `keys` text,
  `values` text
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 |
+-------+----------------------------------------------------------------------------------------------------------+
1 row in set (0.01 sec)
```
