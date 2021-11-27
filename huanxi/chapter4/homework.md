# Chapter 4

## MYSQL

1.  数据库概念

    1. 数据库是长期存放在计算机内、有组织的、可共享的大量数据的集合。数据库中的数据按照一定的数据模型组织、描述和储存，具有较小的冗余度、较高的数据独立性和易拓展性，并可为各种用户共享。

    2. 数据库数据具有永久存储、有组织和可共享三个基本特点。

2.  什么关系型数据库？什么是非关系型

    1. 关系型数据库

       关系型数据库最典型的数据结构是表，由二维表及其之间的联系所组成的一个数据组

    织。

    2. 非关系型数据库

       一种数据结构化存储方法的集合，可以是文档或者键值对等.

3.  SQL 分类 ( [source from](https://zhuanlan.zhihu.com/p/215945247) )

| langage               | description                                                                                                                                                                                                                                             |
| --------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 数据库查询语言（DQL） | 数据查询语言 DQL 基本结构是由 SELECT 子句，FROM 子句，WHERE  子句组成的查询块，简称 DQL，Data Query Language。代表关键字为 select。                                                                                                                     |
| 数据库操作语言（DML） | 用户通过它可以实现对数据库的基本操作。简称 DML，Data Manipulation Language。代表关键字为 insert、delete 、update。                                                                                                                                      |
| 数据库定义语言（DDL） | 数据定义语言 DDL 用来创建数据库中的各种对象，创建、删除、修改表的结构，比如表、视图、索引、同义词、聚簇等，简称 DDL，Data Denifition Language。代表关键字为 create、drop、alter。和 DML 相比，DML 是修改数据库表中的数据，而 DDL 是修改数据中表的结构。 |
| 事务控制语言（TCL）   | TCL 经常被用于快速原型开发、脚本编程、GUI 和测试等方面，简称：TCL，Trasactional Control Languag。代表关键字为 commit、rollback。                                                                                                                        |
| 数据控制语言（DCL）   | 数据控制语言 DCL 用来授予或回收访问数据库的某种特权，并控制数据库操纵事务发生的时间及效果，对数据库实行监视等。简称：DCL，Data Control Language。代表关键字为 grant、revoke                                                                             |

4. SQL 基础语法

DML:

<div align=center>

| command     | description          |
| ----------- | -------------------- |
| SELECT      | 从数据库表中获取数据 |
| UPDATE      | 更新数据库表中的数据 |
| DELETE      | 从数据库表中删除数据 |
| INSERT INTO | 向数据库表中插入数据 |

</div>

DDL:

<div align=center>

| command         | description          |
| --------------- | -------------------- |
| CREATE DATABASE | 创建新数据库         |
| ALTER DATABASE  | 修改数据库           |
| CREATE TABLE    | 创建新表             |
| ALTER TABLE     | 变更（改变）数据库表 |
| DROP TABLE      | 删除表               |
| CREATE INDEX    | 创建索引（搜索键）   |
| DROP INDEX      | 删除索引             |

</div>

5. DDL—创建数据库，查询数据库，修改数据库，删除数据库

<div align=center>

| command         | description  |
| --------------- | ------------ |
| CREATE DATABASE | 创建新数据库 |
| USE DATABASE    | 查询数据库   |
| ALTER DATABASE  | 修改数据库   |
| DROP DATABASE   | 删除数据库   |

</div>

6. DDL—查询表，修改表，删除表数据

<div align=center>

| command                           | description          |
| --------------------------------- | -------------------- |
| SELECT \* FROM                    | 查询表               |
| UPDATE FROM SET id = ? TO ? WHERE | 更新数据库表中的数据 |
| DELETE ? FROM ? WHERE             | 从数据库表中删除数据 |

</div>

7. MYSQL 数据类型—数值类型，字符串类型，日期类型 ( [source from](https://www.runoob.com/mysql/mysql-data-types.html))

数值类型:

<table class="reference">
<tbody>
<tr>
<th width="10%">
类型
</th>
<th width="15%">
大小
</th>
<th width="30%">
范围（有符号）
</th>
<th width="30%">
范围（无符号）
</th>
<th width="15%">
用途
</th>
</tr>
<tr>
<td>
TINYINT
</td>
<td>
1 Bytes
</td>
<td>
(-128，127)
</td>
<td>
(0，255)
</td>
<td>
小整数值
</td>
</tr>
<tr>
<td>
SMALLINT
</td>
<td>
2 Bytes
</td>
<td>
(-32 768，32 767)
</td>
<td>
(0，65 535)
</td>
<td>
大整数值
</td>
</tr>
<tr>
<td>
MEDIUMINT
</td>
<td>
3  Bytes
</td>
<td>
(-8 388 608，8 388 607)
</td>
<td>
(0，16 777 215)
</td>
<td>
大整数值
</td>
</tr>
<tr>
<td>
INT或INTEGER
</td>
<td>
4  Bytes
</td>
<td>
(-2 147 483 648，2 147 483 647)
</td>
<td>
(0，4 294 967 295)
</td>
<td>
大整数值
</td>
</tr>
<tr>
<td>
BIGINT
</td>
<td>
8  Bytes
</td>
<td>
(-9,223,372,036,854,775,808，9 223 372 036 854 775 807)
</td>
<td>
(0，18 446 744 073 709 551 615)
</td>
<td>
极大整数值
</td>
</tr>
<tr>
<td>
FLOAT
</td>
<td>
4  Bytes
</td>
<td>
(-3.402 823 466 E+38，-1.175 494 351 E-38)，0，(1.175 494 351 E-38，3.402 823 466 351 E+38)
</td>
<td>
0，(1.175 494 351 E-38，3.402 823 466 E+38)
</td>
<td>
单精度<br>浮点数值
</td>
</tr>
<tr>
<td>
DOUBLE
</td>
<td>
8  Bytes
</td>
<td>
(-1.797 693 134 862 315 7 E+308，-2.225 073 858 507 201 4 E-308)，0，(2.225 073 858 507 201 4 E-308，1.797 693 134 862 315 7 E+308)
</td>
<td>
0，(2.225 073 858 507 201 4 E-308，1.797 693 134 862 315 7 E+308)
</td>
<td>
双精度<br>浮点数值
</td>
</tr>
<tr>
<td>
DECIMAL
</td>
<td>
对DECIMAL(M,D) ，如果M&gt;D，为M+2否则为D+2
</td>
<td>
依赖于M和D的值
</td>
<td>
依赖于M和D的值
</td>
<td>
小数值
</td>
</tr>
</tbody>
</table>

字符串类型:

<table class="reference">
<tbody>
<tr>
<th width="20%">
类型
</th>
<th width="25%">
大小
</th>
<th width="55%">
用途
</th>
</tr>
<tr>
<td>
CHAR
</td>
<td>
0-255 bytes
</td>
<td>
定长字符串
</td>
</tr>
<tr>
<td>
VARCHAR
</td>
<td>
0-65535 bytes
</td>
<td>
变长字符串
</td>
</tr>
<tr>
<td>
TINYBLOB
</td>
<td>
0-255 bytes
</td>
<td>
不超过 255 个字符的二进制字符串
</td>
</tr>
<tr>
<td>
TINYTEXT
</td>
<td>
0-255 bytes
</td>
<td>
短文本字符串
</td>
</tr>
<tr>
<td>
BLOB
</td>
<td>
0-65 535 bytes
</td>
<td>
二进制形式的长文本数据
</td>
</tr>
<tr>
<td>
TEXT
</td>
<td>
0-65 535 bytes
</td>
<td>
长文本数据
</td>
</tr>
<tr>
<td>
MEDIUMBLOB
</td>
<td>
0-16 777 215 bytes
</td>
<td>
二进制形式的中等长度文本数据
</td>
</tr>
<tr>
<td>
MEDIUMTEXT
</td>
<td>
0-16 777 215 bytes
</td>
<td>
中等长度文本数据
</td>
</tr>
<tr>
<td>
LONGBLOB
</td>
<td>
0-4 294 967 295 bytes
</td>
<td>
二进制形式的极大文本数据
</td>
</tr>
<tr>
<td>
LONGTEXT
</td>
<td>
0-4 294 967 295 bytes
</td>
<td>
极大文本数据
</td>
</tr>
</tbody>
</table>

日期类型:

<table class="reference">
<tbody>
<tr>
<th width="10%">
类型
</th>
<th width="10%">
大小<br>( bytes)
</th>
<th width="40%">
范围
</th>
<th width="20%">
格式
</th>
<th>
用途
</th>
</tr>
<tr>
<td width="10%">
DATE
</td>
<td width="10%">
3
</td>
<td>
1000-01-01/9999-12-31
</td>
<td>
YYYY-MM-DD
</td>
<td>
日期值
</td>
</tr>
<tr>
<td width="10%">
TIME
</td>
<td width="10%">
3
</td>
<td>
'-838:59:59'/'838:59:59'
</td>
<td>
HH:MM:SS
</td>
<td>
时间值或持续时间
</td>
</tr>
<tr>
<td width="10%">
YEAR
</td>
<td width="10%">
1
</td>
<td>
1901/2155
</td>
<td>
YYYY
</td>
<td>
年份值
</td>
</tr>
<tr>
<td width="10%">
DATETIME
</td>
<td width="10%">
8
</td>
<td width="40%">
1000-01-01 00:00:00/9999-12-31 23:59:59
</td>
<td>
YYYY-MM-DD HH:MM:SS
</td>
<td>
混合日期和时间值
</td>
</tr>
<tr>
<td width="10%">
TIMESTAMP
</td>
<td width="10%">
4
</td>
<td width="40%">
<p>1970-01-01 00:00:00/2038 </p>

<p>结束时间是第 <strong>2147483647</strong> 秒，北京时间 <strong>2038-1-19 11:14:07</strong>，格林尼治时间 2038年1月19日 凌晨 03:14:07</p>
</td>
<td>
YYYYMMDD HHMMSS
</td>
<td>
混合日期和时间值，时间戳
</td>
</tr>
</tbody>
</table>

8. MYSQL 约束— 什么是约束？什么是唯一约束，主键约束，联合主键约束，非空约束？ [source from](https://blog.csdn.net/z_johnny/article/details/113820405)

约束:

一种限制，用于限制表中的数据，为了保证表中的数据的准确和可靠性。MySQL 数据库通过约束(constraints)防止无效的数据进入到表中，以保护数据的实体完整性。

六种约束:

| constraints | description                                                                                                                                                                                                                                |
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| NOT NULL    | 非空约束，用于约束该字段的值不能为空。比如姓名、学号等。                                                                                                                                                                                   |
| DEFAULT     | 默认值约束，用于约束该字段有默认值，约束当数据表中某个字段不输入值时，自动为其添加一个已经设置好的值。比如性别。                                                                                                                           |
| PRIMARY KEY | 主键约束，用于约束该字段的值具有唯一性，至多有一个，可以没有，并且非空。比如学号、员工编号等。                                                                                                                                             |
| UNIQUE      | 唯一约束，用于约束该字段的值具有唯一性，可以有多个，可以没有，可以为空。比如座位号。                                                                                                                                                       |
| CHECK       | 检查约束，用来检查数据表中，字段值是否有效。比如年龄、性别。                                                                                                                                                                               |
| FOREIGN KEY | 外键约束，外键约束经常和主键约束一起使用，用来确保数据的一致性，用于限制两个表的关系，用于保证该字段的值必须来自于主表的关联列的值。在从表添加外键约束，用于引用主表中某列的值。比如学生表的专业编号，员工表的部门编号，员工表的工种编号。 |

    列级约束：NOT NULL | DEFAULT | PRIMARY KEY | UNIQUE | CHECK
    表级约束：PRIMARY KEY | UNIQUE | CHECK | FOREIGN KEY

9. MYSQL 约束—添加，修改，删除数据，如何做？

| command                                                                                                         | description    |
| --------------------------------------------------------------------------------------------------------------- | -------------- |
| alter table 表名 add constraint 主键 （形如：PK\_表名） primary key 表名(主键字段)                              | 添加主键约束   |
| alter table 从表 add constraint 外键（形如：FK*从表*主表） foreign key 从表(外键字段) references 主表(主键字段) | 添加外键约束   |
| alter table 表名 modify name varchar(25) default "none"                                                         | 修改默认值约束 |
| alter table 表名 drop primary key                                                                               | 删除主键约束   |

10. MYSQL—举例基础查询？有什么特殊说明的吗？

| command                | description                                                                  |
| ---------------------- | ---------------------------------------------------------------------------- |
| where 子句（条件查询） | 按照“条件表达式”指定的条件进行查询。                                         |
| group by 子句（分组）  | 按照“属性名”指定的字段进行分组。                                             |
| having 子句（筛选）    | 有 group by 才能 having 子句，只有满足“条件表达式”中指定的条件的才能够输出。 |
| order by 子句（排序）  | 按照“属性名”指定的字段进行排序。默认为升序 sec, 降序 desc                    |
| limit（限制结果集）。  | offset 与 limit 进行分页                                                     |

注意: 最好不要使用 select \*, 影响性能.

11. MYSQL—举例 Where 查询

`select name, age, email from Tutors where name = 'huanxi'`

12. MYSQL—举例 Like 查询,有什么弊端？有什么替代技术吗？

使用'%', 匹配任意数量字符, '\_', 匹配单个数量字符. 可以匹配的模式有限.

可以使用正则表达式进行替代, 获得更高的灵活性, 但是一些特性根据引擎的支持与否决定.

13. MYSQL—分组查询

```sql
-- 按照降序列出班级平均分
select avg(score)
from student
group by classid
order by avg(score) desc;
```

14. MYSQL—Having 如何使用

```sql
-- 按照降序列出平均分大于60的班级平均分
select avg(score)
from student
group by classid
having avg(score) > 60
order by avg(score) desc;
```

15. 如何进行一对一，一对多，多对多关系查询

一对一:

`select name form Tutors where id = 1`

一对多:

学生有多门课程, courses_stu 表有 stuid 与 course 联合主键.

```sql
select course
from student, courses_stu
where student.id = courses.stuid and student.id = 1
```

多对多:

老师教授多门课程, 学生有多门课程, courses_stu 表有 stuid 与 course 联合主键.

```sql
-- 学生id为1的学生上的课程及其代课老师.

select student.name, tutor.name, courses_stu.course
from student, courses, tutor, courses_stu
where
  student.id = courses_stu.stuid
and
  tutor.id = courses.tutorid
and
  courses_stu.id = courses.id
and
  student.id = 1;

```

16. 什么是外键约束，有什么弊端？如何解决？

`外键约束:`

由数据库自身保证数据一致性，完整性，更可靠，因为程序很难 100％保证数据 的完整性，而用外键即使在数据库服务器当机或者出现其他问题的时候，也能够最大限度的保证数据的一致性和完整性。

不使用外键，会导致数据冗余，在级联最底层的表可能会重复好几层的数据。必然导致最底层的表数据量翻倍，IO 瓶颈是数据库性能瓶颈之一

`弊端:`

可以用触发器或应用程序保证数据的完整性(外键可代替)

过分强调或者说使用主键／外键会平添开发难度，导致表过多等问题

不用外键时数据管理简单，操作方便，性能高（导入导出等操作，在 insert, update, delete 数据的时候更快）

`解决方案:`

寻找替代外键的解决方案保证数据库数据的完整性.

数据完整性几乎都是业务的要求，理应由业务部分负责维护, 需要精心的设计, 以达到伸缩性与性能问题.

17. 索引的分类

<div align=center>

| index type         | description                                                                                                                                                         |
| ------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------ | -------------------- |
| 单值索引           | 单值索引是最通用,最简单的一种索引, 一个索引只包含一个列,一个表中可以有多个单列索引.                                                                                 |
| 唯一索引           | 索引列的值必须唯一,但可以为 null;                                                                                                                                   |
| 复合索引(联合索引) | mysql 从左到右的使用索引中的字段,一个查询可以只使用索引的一部分,但是只能从最左侧开始. 例如, 我们定义了复合索引 index(c1,c2,c3),则我们进行查找的时候可以 c1 , c2 ,c3 | c1 ,c2 | c1 这三种组合来查找, |

</div>

18. 如何创建，查看，删除索引

`创建单值索引:`

```sql
#外部创建
CREATE INDEX [indexname]ON t1(colname);

#创建表的时候创建
CREATE TABLE mytable(
ID INT NOT NULL,
username VARCHAR(16) NOT NULL,
INDEX [indexName] (username(length))
);

#alter语句添加
ALTER table tableName ADD INDEX indexName(columnName)
```

`创建唯一索引:`

```sql
#创建表时添加,可以使用标记约束和列级约束
CREATE TABLE mytable(
ID INT NOT NULL,   UNIQUE
username VARCHAR(16) NOT NULL,
UNIQUE [indexName] (rowname(length))
);

#使用alter 语句
ALTER TABLE T1  ADD  UNIQUE [indexname] (rowname(length))

#外部创建
CREATE UNIQUE INDEX [indexname] ON TABLENAME(rowname(length))
```

`创建符合索引:`

```sql
CREATE INDEX idx_c1_c2_c3 ON tablename(c1,c2,c3)
```

`查看索引:`

`SHOW INDEX FROM <表名> [ FROM <数据库名>]`

`删除索引:`

```sql
DROP INDEX index_name ON talbe_name

ALTER TABLE table_name DROP INDEX index_name
```

19. 介绍一下数据库事务及其特点

事务是应用程序中一系列严密的操作，所有操作必须成功完成，否则在每个操作中所作的所有更改都会被撤消。也就是事务具有原子性，一个事务中的一系列的操作要么全部成功，要么一个都不做。

事务的四大特征:

<div align=center>

| feature | description                                                                                                                                                                                                                                                                                         |
| ------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 原子性  | 整个事务中的所有操作，要么全部完成，要么全部不完成，不可能停滞在中间某个环节。事务在执行过程中发生错误，会被回滚（Rollback）到事务开始前的状态，就像这个事务从来没有执行过一样。                                                                                                                    |
| 一致性  | 在事务开始之前和事务结束以后，数据库的完整性约束没有被破坏。                                                                                                                                                                                                                                        |
| 隔离性  | 隔离状态执行事务，使它们好像是系统在给定时间内执行的唯一操作。如果有两个事务，运行在相同的时间内，执行 相同的功能，事务的隔离性将确保每一事务在系统中认为只有该事务在使用系统。这种属性有时称为串行化，为了防止事务操作间的混淆， 必须串行化或序列化请 求，使得在同一时间仅有一个请求用于同一数据。 |
| 持久性  | 在事务完成以后，该事务所对数据库所作的更改便持久的保存在数据库之中，并不会被回滚。                                                                                                                                                                                                                  |

</div>

20. 举例说明数据库事务案例

转账场景:

```sql
            create database bank;
            use bank;
            create table account(
                name varchar(20),
                money int
            );

            insert into account values('hejh','1000');
            insert into account values('swy','1000');

            set autocommit = off;
            update account set money = money - 100 where name='';
            update account set money = money + 100 where name='swy';
            commit;
```

21. 数据库的事务隔离级别

<div align=center>

| conception                      | description                                                                                      |
| ------------------------------- | ------------------------------------------------------------------------------------------------ |
| 脏读(dirty read)                | 当一个事务读取另一个事务尚未提交的事务时, 产生脏读                                               |
| 不可重复读 (nonrepeatable read) | 一种查询语句在一个事务中查询多次, 因为其他事务的**修改**或**删除**导致结果不一样, 产生不可重复读 |
| 幻读 (phantom read)             | 一种查询语句在一个事务中查询多次, 因为其他事务的**插入**导致结果不一样, 产生幻读                 |

</div>

<div align=center>

| 隔离级别         | dirty read | nonrepeatable read | phantom read | 加锁读 |
| ---------------- | ---------- | ------------------ | ------------ | ------ |
| Read uncommitted | ✅         | ✅                 | ✅           | ❌     |
| Read committed   | ❌         | ✅                 | ✅           | ❌     |
| Repeatable Read  | ❌         | ❌                 | ❌           | ❌     |
| Read uncommitted | ❌         | ❌                 | ❌           | ✅     |

</div>

22. MYSQL—举例聚合函数

- AVG()函数

AVG() 通过对表中行数计数并计算特定列值之和，求得该列的平均值。 AVG() 可用来返回所有列的平均值，也可以用来返回特定列或行的平均值。

```sql
select avg(price) as avg_price from h_info; -- 计算整张表的平均价格。
select avg(price) as avg_price from h_info where vent_id = 2; -- 计算vent_id=2的平均价格
```

- COUNT()函数

COUNT() 函数进行计数。可利用 COUNT() 确定表中行的数目或符合特定条件的行的数目。

使用 COUNT(\_) 对表中行的数目进行计数，不管表列中包含的是空值（ NULL ）还是非空值。
使用 COUNT(column) 对特定列中具有值的行进行计数，忽略 NULL 值。

```sql

select count(price) as count from h_info;  -- 表中一共有几条 price 数据。
select count(_) as count from h_info; -- 表中有几条记录
```

- MAX()函数，MIN()函数

MAX() 返回指定列中的最大值。 MAX() 要求指定列名。

```sql
select max(price) as max_price from h_info; -- 查找数值最大
select max(date) as max_date from h_info; -- 查找日期最大
```

- SUM()函数

SUM() 用来返回指定列值的和（总计）。

```sql
select sum(price) as sum_price from h_info; -- 查找数值最大
```

23. MYSQL—举例字符串函数

<div align=center><img src="https://tva4.sinaimg.cn/large/006cK6rNgy1gwkwjlqhg3j315t0kz1kx.jpg">

</div>

24. MYSQL—举例日期函数
<table cellspacing="0" cellpadding="2" border="1">
	<thead>
		<tr>
			<th>
				<strong>函数名称</strong></th>
			<th>
				<strong>描述</strong></th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_adddate.html">ADDDATE()</a></td>
			<td>
				相加日期</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_addtime.html">ADDTIME()</a></td>
			<td>
				相加时间</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_convert_tz.html">CONVERT_TZ()</a></td>
			<td>
				从一个时区转换到另一个时区</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_curdate.html">CURDATE()</a></td>
			<td>
				返回当前日期</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_current_date.html">CURRENT_DATE(), CURRENT_DATE</a></td>
			<td>
				CURDATE() 函数的同义词</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_current_time.html">CURRENT_TIME(), CURRENT_TIME</a></td>
			<td>
				CURTIME()&nbsp;函数的同义词</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_current_timestamp.html">CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP</a></td>
			<td>
				NOW()&nbsp;函数的同义词</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_curtime.html">CURTIME()</a></td>
			<td>
				返回当前时间</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql-date-time-functions.html#function_date-add">DATE_ADD()</a></td>
			<td>
				两个日期相加</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_date_format.html">DATE_FORMAT()</a></td>
			<td>
				按格式指定日期</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_date_sub.html">DATE_SUB()</a></td>
			<td>
				两个日期相减</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_date.html">DATE()</a></td>
			<td>
				提取日期或日期时间表达式的日期部分</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_datediff.html">DATEDIFF()</a></td>
			<td>
				两个日期相减</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_day.html">DAY()</a></td>
			<td>
				DAYOFMONTH() 函数的同义词</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_dayname.html">DAYNAME()</a></td>
			<td>
				返回星期的名字</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_dayofmonth.html">DAYOFMONTH()</a></td>
			<td>
				返回该月的第几天 (1-31)</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_dayofweek.html">DAYOFWEEK()</a></td>
			<td>
				返回参数的星期索引</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_dayofyear.html">DAYOFYEAR()</a></td>
			<td>
				返回一年中的天 (1-366)</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_extract.html">EXTRACT</a></td>
			<td>
				提取日期部分</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_from_days.html">FROM_DAYS()</a></td>
			<td>
				日期的数字转换为一个日期</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_from_unixtime.html">FROM_UNIXTIME()</a></td>
			<td>
				格式化日期为UNIX时间戳</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_hour.html">HOUR()</a></td>
			<td>
				提取小时部分</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_last_day.html">LAST_DAY</a></td>
			<td>
				返回该参数对应月份的最后一天</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_localtime.html">LOCALTIME(), LOCALTIME</a></td>
			<td>
				NOW()&nbsp;函数的同义词</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_localtimestamp.html">LOCALTIMESTAMP, LOCALTIMESTAMP()</a></td>
			<td>
				NOW()&nbsp;函数的同义词</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_makedate.html">MAKEDATE()</a></td>
			<td>
				从一年的年份和日期来创建日期</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql-date-time-functions.html#function_maketime">MAKETIME</a></td>
			<td>
				MAKETIME()</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_microsecond.html">MICROSECOND()</a></td>
			<td>
				从参数中返回微秒</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_minute.html">MINUTE()</a></td>
			<td>
				从参数返回分钟</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_month.html">MONTH()</a></td>
			<td>
				通过日期参数返回月份</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_monthname.html">MONTHNAME()</a></td>
			<td>
				返回月份的名称</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_now.html">NOW()</a></td>
			<td>
				返回当前日期和时间</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_period_add.html">PERIOD_ADD()</a></td>
			<td>
				添加一个周期到一个年月</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_period_diff.html">PERIOD_DIFF()</a></td>
			<td>
				返回两个时期之间的月数</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_quarter.html">QUARTER()</a></td>
			<td>
				从一个日期参数返回季度</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_sec_to_time.html">SEC_TO_TIME()</a></td>
			<td>
				转换秒为“HH:MM:SS'的格式</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_second.html">SECOND()</a></td>
			<td>
				返回秒 (0-59)</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_str_to_date.html">STR_TO_DATE()</a></td>
			<td>
				转换一个字符串为日期</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_subdate.html">SUBDATE()</a></td>
			<td>
				当调用三个参数时，它就是 DATE_SUB() 的代名词</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_subtime.html">SUBTIME()</a></td>
			<td>
				相减时间</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_sysdate.html">SYSDATE()</a></td>
			<td>
				返回函数执行时的时间</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_time_format.html">TIME_FORMAT()</a></td>
			<td>
				格式化为时间</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_time_to_sec.html">TIME_TO_SEC()</a></td>
			<td>
				将参数转换成秒并返回</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_time.html">TIME()</a></td>
			<td>
				提取表达式传递的时间部分</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_timediff.html">TIMEDIFF()</a></td>
			<td>
				相减时间</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_timestamp.html">TIMESTAMP()</a></td>
			<td>
				带一个参数，这个函数返回日期或日期时间表达式。有两个参数，参数的总和</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_timestampadd.html">TIMESTAMPADD()</a></td>
			<td>
				添加一个时间间隔到datetime表达式</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_timestampdiff.html">TIMESTAMPDIFF()</a></td>
			<td>
				从日期时间表达式减去的间隔</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_to_days.html">TO_DAYS()</a></td>
			<td>
				返回日期参数转换为天</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_unix_timestamp.html">UNIX_TIMESTAMP()</a></td>
			<td>
				返回一个UNIX时间戳</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_utc_date.html">UTC_DATE()</a></td>
			<td>
				返回当前UTC日期</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_utc_time.html">UTC_TIME()</a></td>
			<td>
				返回当前UTC时间</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_utc_timestamp.html">UTC_TIMESTAMP()</a></td>
			<td>
				返回当前UTC日期和时间</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_week.html">WEEK()</a></td>
			<td>
				返回周数</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_weekday.html">WEEKDAY()</a></td>
			<td>
				返回星期的索引</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_weekofyear.html">WEEKOFYEAR()</a></td>
			<td>
				返回日期的日历周 (1-53)</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_year.html">YEAR()</a></td>
			<td>
				返回年份</td>
		</tr>
		<tr>
			<td>
				<a href="http://www.yiibai.com/mysql/mysql_function_yearweek.html">YEARWEEK()</a></td>
			<td>
				返回年份和周</td>
		</tr>
	</tbody>
</table>

## GORM

1. 如何通过 Gorm 连接数据库

```go
package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func main() {
    // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

    dsn := "root:123456@tcp(127.0.0.1:3306)/huanxi?charset=utf8mb4&parseTime=True&loc=Local"
    _, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }
}
```

使用 logger 打印日志信息, 设置 Info 的日志级别

```go
package main

import (
    "log"
    "os"
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

func main() {
    // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
        logger.Config{
            SlowThreshold:             time.Second, // 慢 SQL 阈值
            LogLevel:                  logger.Info, // 日志级别
            IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
            Colorful:                  true,        // 启用彩色打印
        },
    )

    dsn := "root:1@tcp(127.0.0.1:3306)/huanxi?charset=utf8mb4&parseTime=True&loc=Local"
    _, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: newLogger,
    })
    if err != nil {
        panic(err)
    }
}
```

2. 如何通过 Gorm 生成表解构

声明模型时, 按照约定数据库的名称会变为结构体的复数, 例如, Tag 在数据库中为 tags. 使用 TableName 覆盖约定.

```go
package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

type Tag struct {
    gorm.Model
    Name string `json:"name"`
}

/* 等价与
type Tag struct {
    ID        uint `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt DeletedAt `gorm:"index"`
    Name string `json:"name"`
}*/

// TableName 会将 Tag 的表名重写为 `my_tag`
func (Tag) TableName() string {
    return "my_tag"
}

func init() {
    // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
        logger.Config{
            SlowThreshold:             time.Second, // 慢 SQL 阈值
            LogLevel:                  logger.Info, // 日志级别
            IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
            Colorful:                  true,        // 启用彩色打印
        },
    )

    dsn := "root:123456@tcp(127.0.0.1:3306)/huanxi?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: newLogger,
    })
    if err != nil {
        panic(err)
    }

    err = db.AutoMigrate(Tag{})
    if err != nil {
        panic("AutoMigrate Failed")
    }

}

func main() {
    fmt.Println("Init finished")
}
```

3. Gorm 中对零值的处理，update 和 updates 都有哪些坑？

```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}

/* 等价与
type Tag struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
	Name string `json:"name"`
}*/

// TableName 会将 Tag 的表名重写为 `my_tag`
func (Tag) TableName() string {
	return "my_tag"
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}

// TableName 会将 Tag 的表名重写为 `my_user`
func (User) TableName() string {
	return "my_user"
}

var db *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger.Error, // 日志级别
			IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,         // 启用彩色打印
		},
	)

	dsn := "root:1@tcp(127.0.0.1:3306)/huanxi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(Tag{}, User{})
	if err != nil {
		panic("AutoMigrate Failed")
	}

}

func main() {
	// 创建User
	//u1 := User{Name: "u1", Password: "passwd1"}
	//u2 := User{Name: "u2", Password: "passwd2"}
	//us := []User{u1, u2}
	//db.Create(&us)

	fmt.Println("=========Save===========")
	// Save 会保存所有的字段，即使字段是零值
	var t1 Tag
	db.First(&t1)
	fmt.Println(t1)

	t1.Name = "t1 modified"
	db.Save(&t1)
	fmt.Println(t1)

	fmt.Println("=========Update 更新单列===========")
	var t2 Tag
	db.First(&t2, 2) // id = 2
	fmt.Println(t2)

	db.Model(&Tag{}).Where("id = ?", 2).Update("name", "t2 modified")
	db.First(&t2, 2) // id = 2
	fmt.Println(t2)

	fmt.Println("=========Updates 更新多列===========")
	// Updates 方法支持 struct 和 map[string]interface{} 参数。
	// 当使用 struct 更新时，默认情况下，GORM 只会更新非零值的字段

	fmt.Println("=========结构体===========")
	var u1 User
	db.First(&u1)
	fmt.Println(u1)

	db.Model(&User{}).Where("id = ?", 1).Updates(User{Name: "u1 modified", Password: ""})
	db.First(&u1, 1)
	fmt.Println(u1)

	fmt.Println("=========MAP===========")
	var u2 User
	db.First(&u2, 2)
	fmt.Println(u2)

	db.Model(&User{}).Where("id = ?", 2).Updates(map[string]interface{}{"name": "u2 modified", "password": ""})
	db.First(&u2, 2)
	fmt.Println(u2)
}

```

<div align=center><img src="https://tvax4.sinaimg.cn/large/006cK6rNgy1gweodqray5j30x907yn2n.jpg"></div>

<div align=center><img src="https://tvax3.sinaimg.cn/large/006cK6rNgy1gweo94yvk0j30py0guajn.jpg"></div>

⚠️ 注意四点 :

    1. 新建 User 表, 可以看到 Save 会保存所有值, 即使是 0 值.

    2. Update 同样会保存 0 值, 没有在结构体或 map 中的字段, 会被忽略.

    3. Updates, 在使用结构体更新时, 不会更新 0 值的字段, map 则会更新所有存在的字段.

如果想使用结构体更新 0 值字段, 使用指针类型或 sql.NullXXX, 如 sql.NullString, sql.NullInt64, 包装一个 string or int64 类型的值, 以及一个 bool 值.

使用 map 更新字段时, 需要指定 model.

4. 添加数据有几种方式？

```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}

/* 等价与
type Tag struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
	Name string `json:"name"`
}*/

// TableName 会将 Tag 的表名重写为 `my_tag`
func (Tag) TableName() string {
	return "my_tag"
}

var db *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 启用彩色打印
		},
	)

	dsn := "root:1@tcp(127.0.0.1:3306)/huanxi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(Tag{})
	if err != nil {
		panic("AutoMigrate Failed")
	}

}

func main() {
	fmt.Println("=========单行插入===========")
	fmt.Println("==========结构体============")
	t1 := Tag{Name: "t1"}
	result := db.Create(&t1)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数
	fmt.Println("==========MAP============")
	t2 := map[string]interface{}{"name": "t2"}
	result = db.Model(&Tag{}).Create(t2)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数

	fmt.Println("=========多行插入===========")
	fmt.Println("==========结构体============")
	t3s := []Tag{
		{Name: "t3"},
		{Name: "t4"},
		{Name: "t5"},
	}
	result = db.Create(&t3s)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数
	fmt.Println("==========MAP============")
	t4s := []map[string]interface{}{
		{"name": "t6"},
		{"name": "t7"},
		{"name": "t8"},
	}
	result = db.Model(&Tag{}).Create(&t4s)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数

	fmt.Println("=========多行插入InBatch===========")
	fmt.Println("==========结构体============")
	t5s := []Tag{
		{Name: "t9"},
		{Name: "t10"},
		{Name: "t11"},
	}
	result = db.CreateInBatches(&t5s, 2) // 一次2条数据进行插入
	fmt.Println(result.Error)            // 返回 error
	fmt.Println(result.RowsAffected)     // 返回插入记录的条数
	fmt.Println("==========MAP============")
	t6s := []map[string]interface{}{
		{"name": "t12"},
		{"name": "t13"},
		{"name": "t14"},
	}
	result = db.Model(&Tag{}).CreateInBatches(t6s, 3) // 一次三条数据插入
	fmt.Println(result.Error)                         // 返回 error
	fmt.Println(result.RowsAffected)                  // 返回插入记录的条数

}

```

<div align=center><img src="https://tvax4.sinaimg.cn/large/006cK6rNgy1gwen6tlt88j30mc0baahm.jpg"></div>

⚠️ 注意三点:

    1. 使用 map 插入数据时, 需要指定 Model.

    2. 使用 map 插入数据时, 不会自动更新 created_at, updated_at 等子段. 原因是, gorm 只会记录 map 中的字段.

    3. 使用 map 插入数据时, association 不会被调用，且主键也不会自动填充.

5. 查询数据有几种方式？

```go
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}

/* 等价与
type Tag struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
	Name string `json:"name"`
}*/

// TableName 会将 Tag 的表名重写为 `my_tag`
func (Tag) TableName() string {
	return "my_tag"
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}

// TableName 会将 Tag 的表名重写为 `my_user`
func (User) TableName() string {
	return "my_user"
}

var db *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger.Error, // 日志级别
			IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,         // 启用彩色打印
		},
	)

	dsn := "root:1@tcp(127.0.0.1:3306)/huanxi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(Tag{}, User{})
	if err != nil {
		panic("AutoMigrate Failed")
	}

}

func main() {
	fmt.Println("=========查询单条数据===========")
	var u1 User
	var u2 User
	var u3 User
	// 获取第一条记录（主键升序）
	db.First(&u1)
	// SELECT * FROM users ORDER BY id LIMIT 1;
	fmt.Println(u1)

	// 获取一条记录，没有指定排序字段
	db.Take(&u2)
	// SELECT * FROM users LIMIT 1;
	fmt.Println(u2)

	// 获取最后一条记录（主键降序）
	db.Last(&u3)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	fmt.Println(u3)

	result := db.First(&u3)
	fmt.Println(result.RowsAffected) // 返回找到的记录数
	fmt.Println(result.Error)        // returns error or nil

	// 检查 ErrRecordNotFound 错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("record not found")
	} else {
		fmt.Println("record found successfully")
	}

	fmt.Println("=========查询多条数据===========")
	// 查询所有数据
	var users []User
	result = db.Find(&users)
	fmt.Println(result.RowsAffected) // 返回找到的记录数
	fmt.Println(users)

	result = db.Where("id in ?", []int64{1, 2}).Find(&users)
	fmt.Println(result.RowsAffected) // 返回找到的记录数
	fmt.Println(users)

}

```

<div align=center><img src="https://tvax4.sinaimg.cn/large/006cK6rNgy1gwep2lgroej30zk094wk7.jpg"></div>

⚠️:

查询时, 使用结构体需要传入指针或指定模型 Model(). 使用 map 时, 需要指定 Table()

6. Gorm 如何实现删除数据的？

```go
// Email 的 ID 是 `10`
db.Delete(&email)
// DELETE from emails where id = 10;

// 带额外条件的删除
db.Where("name = ?", "jinzhu").Delete(&email)
// DELETE from emails where id = 10 AND name = "jinzhu";

db.Delete(&User{}, 10)
// DELETE FROM users WHERE id = 10;

db.Delete(&User{}, "10")
// DELETE FROM users WHERE id = 10;

db.Delete(&users, []int{1,2,3})
// DELETE FROM users WHERE id IN (1,2,3);
```

⚠️ :

    1. GORM 推荐使用软删除, 如果模型中有 gorm.DeletedAt, 则自动进行软删除.

给模型添加软删除:

```go
type User struct {
  ID      int
  Deleted gorm.DeletedAt
  Name    string
}
```

    2. 查询软删除的数据 Unscoped()

```go
db.Unscoped().Where("age = 20").Find(&users)
// SELECT * FROM users WHERE age = 20;
```

    3. 不使用软删除 Unscoped()

```go
db.Unscoped().Delete(&order)
// DELETE FROM orders WHERE id=10;
```

7. 如何进行一对一，一对多，多对多查询？

`db.Preload("Failed_Name").Find(&users)`
