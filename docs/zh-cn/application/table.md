
### 创建表
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

```sql
CREATE TABLE test (id integer , keys TEXT, values TEXT);
```

### 删除表
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

### 修改表(暂不支持)