### SELECT 语句

使用SQL SELECT语句来查询数据。
#### 基本语法:
以下为在P2PDB数据库中查询数据通用的 SELECT 语法：
```sql
SELECT column_name,column_name
FROM table_name
[WHERE Clause]
[LIMIT N][ OFFSET M]

```

#### 解释:
* 查询语句中你可以使用一个或者多个表，表之间使用逗号(,)分割，并使用WHERE语句来设定查询条件。

* SELECT 命令可以读取一条或者多条记录。

* 你可以使用星号（*）来代替其他字段，SELECT语句会返回表的所有字段数据

* 你可以使用 WHERE 语句来包含任何条件。

* 你可以使用 LIMIT 属性来设定返回的记录数。

* 你可以通过OFFSET指定SELECT语句开始查询的数据偏移量。默认情况下偏移量为0。


#### 例子:
以下实例我们将通过 SQL SELECT 命令来获取 P2PDB 数据表 emails 的数据：

```bash
mysql> use p2pdb;
Database changed
mysql>  CREATE TABLE test (id integer , name TEXT, address  TEXT);
Query OK, 0 rows affected (0.02 sec)

mysql> insert into test(id,name,address) values(1,'Alice','Texas');
Query OK, 1 row affected (0.03 sec)

mysql> insert into test values(2,'Alice','Texas');
Query OK, 1 row affected (0.01 sec)

mysql> select * from test;
+------+-------+-------+
| id   | name  | address |
+------+-------+-------+
|    1 | Alice | Texas |
|    2 | Alice | Texas |
+------+-------+-------+
2 rows in set (0.01 sec)
```

### WHERE 子句

### LIKE 子句

### GLOB 子句

### LIMIT 子句

### ORDER BY 子句

### GROUP BY 子句

### Having 子句

### Distinct 关键字
