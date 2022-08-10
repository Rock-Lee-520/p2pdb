# Supported Sqlite SQL Syntax

SQLite 数据类型是一个用来指定任何对象的数据类型的属性。SQLite 中的每一列，每个变量和表达式都有相关的数据类型。

您可以在创建表的同时使用这些数据类型。SQLite 使用一个更普遍的动态类型系统。在 SQLite 中，值的数据类型与值本身是相关的，而不是与它的容器相关。

## Supported Sqlite types
每个存储在 SQLite 数据库中的值都具有以下存储类之一：

存储类	描述
NULL	值是一个 NULL 值。
INTEGER	值是一个带符号的整数，根据值的大小存储在 1、2、3、4、6 或 8 字节中。
REAL	值是一个浮点值，存储为 8 字节的 IEEE 浮点数字。
TEXT	值是一个文本字符串，使用数据库编码（UTF-8、UTF-16BE 或 UTF-16LE）存储。
BLOB	值是一个 blob 数据，完全根据它的输入存储。
SQLite 的存储类稍微比数据类型更普遍。INTEGER 存储类，例如，包含 6 种不同的不同长度的整数数据类型。

## Data manipulation statements

- DELETE
- INSERT
- SELECT
- UPDATE

## Data definition statements

- ADD COLUMN
- ALTER COLUMN
- ALTER TABLE
- CHANGE COLUMN
- CREATE INDEX
- CREATE TABLE
- CREATE VIEW
- DESCRIBE TABLE
- DROP COLUMN
- DROP INDEX
- DROP TABLE
- DROP VIEW
- MODIFY COLUMN
- RENAME COLUMN
- SHOW CREATE TABLE
- SHOW CREATE VIEW
- SHOW DATABASES
- SHOW SCHEMAS
- SHOW TABLES

## Transactional statements

- BEGIN
- COMMIT
- LOCK TABLES
- START TRANSACTION
- UNLOCK TABLES

## Session management statements

- SET

## Utility statements

- EXPLAIN
- USE

## Standard expressions

- WHERE
- HAVING
- LIMIT
- OFFSET
- GROUP BY 
- ORDER BY
- DISTINCT 
- ALL
- AND
- NOT
- OR
- IF
- CASE / WHEN
- NULLIF
- COALESCE 
- IFNULL
- LIKE
- IN / NOT IN
- IS NULL / IS NOT NULL
- INTERVAL
- Scalar subqueries
- Column ordinal references (standard MySQL extension)

## Comparison expressions
- !=
- ==
- \>
- <
- \>=
- <=
- BETWEEN
- IN
- NOT IN
- REGEXP
- IS NOT NULL
- IS NULL

## Aggregate functions

- AVG
- COUNT and COUNT(DISTINCT)
- MAX
- MIN
- SUM (always returns DOUBLE)

## Join expressions

- CROSS JOIN
- INNER JOIN
- LEFT INNER JOIN
- RIGHT INNER JOIN
- NATURAL JOIN

## Arithmetic expressions

- \+ (including between dates and intervals)
- \- (including between dates and intervals)
- \*
- \/
- <<
- \>>
- &
- \|
- ^
- div
- %

## Subqueries

Supported both as a table and as expressions but they can't access the
parent query scope.

## Functions

See README.md for the list of supported functions.

# Notable limitations

The engine is missing many features. The most important ones are noted
below. Our goal over time is 100% compatibility, which means adding
support for the items in this list.

Some features are relatively easy to support, some are more
difficult. Please browse / file issues explaining your use case to
make your case for prioritizing missing features, or feel free to
discuss an implementation plan with us and submit a PR.

## Missing features

- Prepared statements / Execute
- Outer joins
- `AUTO INCREMENT`
- Transaction snapshotting / rollback
- Check constraint 
- Window functions
- Common table expressions (CTEs)
- Stored procedures
- Events
- Cursors
- Triggers
- Users / privileges / `GRANT` / `REVOKE` (via SQL)
- `CREATE TABLE AS`
- `DO`
- `HANDLER`
- `IMPORT TABLE`
- `LOAD DATA` / `LOAD XML`
- `SELECT FOR UPDATE`
- `TABLE` (alternate select syntax)
- `TRUNCATE`
- Alter index
- Alter view
- Create function
