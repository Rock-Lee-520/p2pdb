# 库相关

### 连接数据库


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



### 显示所有数据库

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

### 选择数据库

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

### 创建新数据库(暂不支持)
```
CREATE DATABASE 数据库名;
```
#### 解释
通过自定义名字创建一个新的数据库

#### 例子
```bash

mysql> create DATABASE p2pdb_2;

```

### 删除数据库(暂不支持)
```
DROP DATABASE 数据库名;
```
#### 解释
删除一个数据库,注意`内置数据库information_schema 不可被删除`

#### 例子
```bash

mysql> drop database p2pdb_2;

```