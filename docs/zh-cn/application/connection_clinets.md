## 测试的客户端

理论上所有支持mysql的客户端都可以连接p2pdb,以下这些是我们积极测试以检查他们是否兼容p2pdb的客户端，

!>  注意，由于没有足够的人手投入,P2PDB并不支持在一些可视化客户端中执行DDL语句(mysql-cli除外),如Navicat,DBeaver等,使用时需要注意,如果你对P2PDB感兴趣,欢迎加入P2PDB团队来完善它.


- Mysql-cli
  - [mysql-cli](#mysql-cli)
- Python
  - [pymysql](#pymysql)
  - [mysql-connector](#python-mysql-connector)
  - [sqlalchemy](#python-sqlalchemy)
- Ruby
  - [ruby-mysql](#ruby-mysql)
- [PHP](#php)
- Node.js
  - [mysqljs/mysql](#mysqljs)
- .NET Core
  - [MysqlConnector](#mysqlconnector)
- Java/JVM
  - [mariadb-java-client](#mariadb-java-client)
- Go
  - [go-mysql-driver/mysql](#go-sql-drivermysql)
- C
  - [mysql-connector-c](#mysql-connector-c)

## 客户端使用例子

> ### mysql-cli
```bash
mysql -h 127.0.0.1 -u root -p
Enter password:
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 1
Server version: 5.7.9-Vitess

Copyright (c) 2000, 2022, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| p2pdb              |
+--------------------+
2 rows in set (0.00 sec)

```

> ### pymysql

```bash
import pymysql.cursors

connection = pymysql.connect(host='127.0.0.1',
                             user='root',
                             password='',
                             db='mydb',
                             cursorclass=pymysql.cursors.DictCursor)

try:
    with connection.cursor() as cursor:
        sql = "SELECT * FROM mytable LIMIT 1"
        cursor.execute(sql)
        rows = cursor.fetchall()

        # use rows
finally:
    connection.close()
```

> ### Python mysql-connector

```bash
import mysql.connector

connection = mysql.connector.connect(host='127.0.0.1',
                                user='root',
                                passwd='',
                                port=3306,
                                database='mydb')

try:
    cursor = connection.cursor()
    sql = "SELECT * FROM mytable LIMIT 1"
    cursor.execute(sql)
    rows = cursor.fetchall()

    # use rows
finally:
    connection.close()
```

> ### Python sqlalchemy

```bash
import pandas as pd
import sqlalchemy

engine = sqlalchemy.create_engine('mysql+pymysql://root:@127.0.0.1:3306/mydb')
with engine.connect() as conn:
     repo_df = pd.read_sql_table("mytable", con=conn)
     for table_name in repo_df.to_dict():
        print(table_name)
```

> ### ruby-mysql

```bash
require "mysql"

conn = Mysql::new("127.0.0.1", "root", "", "mydb")
resp = conn.query "SELECT * FROM mytable LIMIT 1"

# use resp

conn.close()
```

> ### php

```php
try {
    $conn = new PDO("mysql:host=127.0.0.1:3306;dbname=mydb", "root", "");
    $conn->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);

    $stmt = $conn->query('SELECT * FROM mytable LIMIT 1');
    $result = $stmt->fetchAll(PDO::FETCH_ASSOC);

    // use result
} catch (PDOException $e) {
    // handle error
}
```

> ### mysqljs

```js
import mysql from 'mysql';

const connection = mysql.createConnection({
    host: '127.0.0.1',
    port: 3306,
    user: 'root',
    password: '',
    database: 'mydb'
});
connection.connect();

const query = 'SELECT * FROM mytable LIMIT 1';
connection.query(query, function (error, results, _) {
    if (error) throw error;

    // use results
});

connection.end();
```

> ### MysqlConnector

```bash
using MySql.Data.MySqlClient;
using System.Threading.Tasks;

namespace something
{
    public class Something
    {
        public async Task DoQuery()
        {
            var connectionString = "server=127.0.0.1;user id=root;password=;port=3306;database=mydb;";

            using (var conn = new MySqlConnection(connectionString))
            {
                await conn.OpenAsync();

                var sql = "SELECT * FROM mytable LIMIT 1";

                using (var cmd = new MySqlCommand(sql, conn))
                using (var reader = await cmd.ExecuteReaderAsync())
                while (await reader.ReadAsync()) {
                    // use reader
                }
            }
        }
    }
}
```

> ### mariadb-java-client

```bash
package org.testing.mariadbjavaclient;

import java.sql.*;

class Main {
    public static void main(String[] args) {
        String dbUrl = "jdbc:mariadb://127.0.0.1:3306/mydb?user=root&password=";
        String query = "SELECT * FROM mytable LIMIT 1";

        try (Connection connection = DriverManager.getConnection(dbUrl)) {
            try (PreparedStatement stmt = connection.prepareStatement(query)) {
                try (ResultSet rs = stmt.executeQuery()) {
                    while (rs.next()) {
                        // use rs
                    }
                }
            }
        } catch (SQLException e) {
            // handle failure
        }
    }
}
```

> ### go-sql-driver/mysql

```go
package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		// handle error
	}

	rows, err := db.Query("SELECT * FROM mytable LIMIT 1")
	if err != nil {
		// handle error
    }

    // use rows
}
```

> ### mysql-connector-c

```bash
#include <my_global.h>
#include <mysql.h>

void finish_with_error(MYSQL *con)
{
    fprintf(stderr, "%s\n", mysql_error(con));
    mysql_close(con);
    exit(1);
}

int main(int argc, char **argv)
{
    MYSQL *con = NULL;
    MYSQL_RES *result = NULL;
    int num_fields = 0;
    MYSQL_ROW row;

    printf("MySQL client version: %s\n", mysql_get_client_info());

    con = mysql_init(NULL);
    if (con == NULL) {
        finish_with_error(con);
    }

    if (mysql_real_connect(con, "127.0.0.1", "root", "", "mydb", 3306, NULL, 0) == NULL) {
        finish_with_error(con);
    }

    if (mysql_query(con, "SELECT name, email, phone_numbers FROM mytable")) {
        finish_with_error(con);
    }

    result = mysql_store_result(con);
    if (result == NULL) {
        finish_with_error(con);
    }

    num_fields = mysql_num_fields(result);
    while ((row = mysql_fetch_row(result))) {
        for(int i = 0; i < num_fields; i++) {
            printf("%s ", row[i] ? row[i] : "NULL");
        }
        printf("\n");
    }

    mysql_free_result(result);
    mysql_close(con);

    return 0;
}
```
