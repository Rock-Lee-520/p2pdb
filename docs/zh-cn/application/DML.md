> ### INSERT 新增

INSERT INTO 语句用于向数据库的某个表中添加新的数据行。
#### 基本语法:
INSERT INTO 语句有两种基本语法，如下所示：
```sql
INSERT INTO TABLE_NAME (column1, column2, column3,...columnN)
VALUES (value1, value2, value3,...valueN);

```

#### 解释:
在这里，column1, column2,...columnN 是要插入数据的表中的列的名称。

如果要为表中的所有列添加值，您也可以不需要在 P2PDB 查询中指定列名称,但要确保值的顺序与列在表中的顺序一致。
```sql
INSERT INTO TABLE_NAME VALUES (value1,value2,value3,...valueN);
```

#### 例子:

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

> ### UPDATE 更新

UPDATE 查询用于修改表中已有的记录。可以使用带有 WHERE 子句的 UPDATE 查询来更新选定行，否则所有的行都会被更新
#### 基本语法:
带有 WHERE 子句的 UPDATE 查询的基本语法如下：
```sql
UPDATE table_name
SET column1 = value1, column2 = value2...., columnN = valueN
WHERE [condition];

```

#### 解释:
可以使用 AND 或 OR 运算符来结合 N 个数量的条件。


#### 例子:
假设test 表有以下记录：
```bash
mysql> select * from test;
+------+-------+-------+
| id   | name  | address |
+------+-------+-------+
|    1 | Alice | Texas |
|    2 | Alice | Texas |
+------+-------+-------+
2 rows in set (0.01 sec)
```

下面是一个实例，它会更新 ID 为 2 的name 为bob

```bash
mysql> update test set name="bob" where id=2;
Query OK, 0 rows affected (0.06 sec)
Rows matched: 0  Changed: 0  Warnings: 0

mysql> select * from test;
+------+-------+-------+
| id   | name  | address |
+------+-------+-------+
|    1 | Alice | Texas |
|    2 | bob   | Texas |
+------+-------+-------+
2 rows in set (0.00 sec)
```

> ### DELETE 删除

DELETE 查询用于删除表中已有的记录。可以使用带有 WHERE 子句的 DELETE 查询来删除选定行，否则所有的记录都会被删除
#### 基本语法:
```sql

DELETE FROM table_name
WHERE [condition];

```

#### 解释:
可以使用 AND 或 OR 运算符来结合 N 个数量的条件。


#### 例子:
假设test 表有以下记录：
```bash
mysql> select * from test;
+------+-------+-------+
| id   | name  | address |
+------+-------+-------+
|    1 | Alice | Texas |
|    2 | bob   | Texas |
+------+-------+-------+
2 rows in set (0.01 sec)
```

下面是一个实例，它会删除 ID 为 2 的用户

```bash
mysql> delete from test where id=2;
Query OK, 1 rows affected (0.03 sec)

mysql> select * from test;
+------+-------+-------+
| id   | name  | alias |
+------+-------+-------+
|    1 | Alice | Texas |
+------+-------+-------+
1 row in set (0.00 sec)
```

如果您想要从 test 表中删除所有记录，则不需要使用 WHERE 子句，DELETE 语句如下：
```bash
mysql> delete from test;
Query OK, 1 rows affected (0.01 sec)
```