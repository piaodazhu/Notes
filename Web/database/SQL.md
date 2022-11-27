# SQL quick start (sqlite3)
## A. create & drop
```sql
server01@server01:~/golearning/my-Go/Web/database$ sqlite3 testDB.db
SQLite version 3.37.2 2022-01-06 13:25:41
Enter ".help" for usage hints.
sqlite> .databses
Error: unknown command or invalid arguments:  "databses". Enter ".help" for help
sqlite> .databases
main: /home/server01/golearning/my-Go/Web/database/testDB.db r/w
sqlite> create table company(
   ...> id int primary key not null,
   ...> name text not null,
   ...> age int not null,
   ...> address char(50),
   ...> salary real
   ...> );
sqlite> .databases
main: /home/server01/golearning/my-Go/Web/database/testDB.db r/w
sqlite> .tables
company
sqlite> create table department(
   ...> id int primary key not null,
   ...> dept char(50) not null,
   ...> emp_id int not null
   ...> );
sqlite> .tables
company     department
sqlite> .schema
CREATE TABLE company(
id int primary key not null,
name text not null,
age int not null,
address char(50),
salary real
);
CREATE TABLE department(
id int primary key not null,
dept char(50) not null,
emp_id int not null
);
sqlite> .schema company
CREATE TABLE company(
id int primary key not null,
name text not null,
age int not null,
address char(50),
salary real
);
sqlite> create table tobedeleted(
   ...> id int primary key not null,
   ...> data float not null);
sqlite> .tables
company      department   tobedeleted
sqlite> drop table tobedeleted;
sqlite> .tables
company     department
```

## B. insert & select
```sql
sqlite> insert into company (id, name, age, address, salary)
   ...> values (1, 'Pual', 32, 'California', 20000.00);
sqlite> insert into company values (2, 'Allen', 25, 'Texas', 15000.00 );
sqlite> insert into company values (3, 'Teddy', 23, 'Norway', 20000.00 );
sqlite> insert into company values (4, 'Mark', 25, 'Rich-Mond ', 65000.00 );
sqlite> insert into company values (5, 'David', 27, 'Texas', 85000.00 );
sqlite> insert into company values (6, 'Kim', 22, 'South-Hall', 45000.00 );
sqlite> insert into company values (7, 'James', 24, 'Houston', 10000.00 );
sqlite> select * from company;
1|Pual|32|California|20000.0
2|Allen|25|Texas|15000.0
3|Teddy|23|Norway|20000.0
4|Mark|25|Rich-Mond |65000.0
5|David|27|Texas|85000.0
6|Kim|22|South-Hall|45000.0
7|James|24|Houston|10000.0
sqlite> .header on
sqlite> .mode column
sqlite> select * from company;
id  name   age  address     salary 
--  -----  ---  ----------  -------
1   Pual   32   California  20000.0
2   Allen  25   Texas       15000.0
3   Teddy  23   Norway      20000.0
4   Mark   25   Rich-Mond   65000.0
5   David  27   Texas       85000.0
6   Kim    22   South-Hall  45000.0
7   James  24   Houston     10000.0
sqlite> select id, name, salary from company;
id  name   salary 
--  -----  -------
1   Pual   20000.0
2   Allen  15000.0
3   Teddy  20000.0
4   Mark   65000.0
5   David  85000.0
6   Kim    45000.0
7   James  10000.0
sqlite> select * from company where age > 24 and id > 3;
id  name   age  address     salary 
--  -----  ---  ----------  -------
4   Mark   25   Rich-Mond   65000.0
5   David  27   Texas       85000.0
sqlite> select * from company where age > 24 or id > 3;
id  name   age  address     salary 
--  -----  ---  ----------  -------
1   Pual   32   California  20000.0
2   Allen  25   Texas       15000.0
4   Mark   25   Rich-Mond   65000.0
5   David  27   Texas       85000.0
6   Kim    22   South-Hall  45000.0
7   James  24   Houston     10000.0
sqlite> select * from company where name like "Da%";
id  name   age  address  salary 
--  -----  ---  -------  -------
5   David  27   Texas    85000.0
sqlite> select * from company where address like "%-%";
id  name  age  address     salary 
--  ----  ---  ----------  -------
4   Mark  25   Rich-Mond   65000.0
6   Kim   22   South-Hall  45000.0
sqlite> select * from company where age in (23,24);
id  name   age  address  salary 
--  -----  ---  -------  -------
3   Teddy  23   Norway   20000.0
7   James  24   Houston  10000.0
sqlite> select * from company where age between 25 and 27;
id  name   age  address     salary 
--  -----  ---  ----------  -------
2   Allen  25   Texas       15000.0
4   Mark   25   Rich-Mond   65000.0
5   David  27   Texas       85000.0
sqlite> select * from company limit 4;
id  name   age  address     salary 
--  -----  ---  ----------  -------
1   Pual   32   California  20000.0
2   Allen  25   Texas       15000.0
3   Teddy  23   Norway      20000.0
4   Mark   25   Rich-Mond   65000.0
sqlite> select * from company limit 4 offset 3;
id  name   age  address     salary 
--  -----  ---  ----------  -------
4   Mark   25   Rich-Mond   65000.0
5   David  27   Texas       85000.0
6   Kim    22   South-Hall  45000.0
7   James  24   Houston     10000.0
sqlite> select count(*) from company;
count(*)
--------
7       
sqlite> select count(*) from company where id >= 3;
count(*)
--------
5       
sqlite> select current_timestamp;
current_timestamp  
-------------------
2022-11-25 12:03:49
```

difference between 'like' and 'glob':
1. 'like' use % and _, but 'glob' use * and ?
2. 'like' ignore case, but 'glob' is case-sensitive

## C. update & delete

```sql
sqlite> create table test_company as select * from company;
sqlite> .tables
company       department    test_company
sqlite> update test_company set salary = 0 where age >= 30;
sqlite> select * from test_company;
id  name   age  address     salary 
--  -----  ---  ----------  -------
1   Pual   32   California  0.0    
2   Allen  25   Texas       15000.0
3   Teddy  23   Norway      20000.0
4   Mark   25   Rich-Mond   65000.0
5   David  27   Texas       85000.0
6   Kim    22   South-Hall  45000.0
7   James  24   Houston     10000.0
sqlite> delete from test_company where age >= 30;
sqlite> select * from test_company;
id  name   age  address     salary 
--  -----  ---  ----------  -------
2   Allen  25   Texas       15000.0
3   Teddy  23   Norway      20000.0
4   Mark   25   Rich-Mond   65000.0
5   David  27   Texas       85000.0
6   Kim    22   South-Hall  45000.0
7   James  24   Houston     10000.0
```

## D. order by & group by

```sql
sqlite> select * from company order by id desc;
id  name   age  address     salary 
--  -----  ---  ----------  -------
7   James  24   Houston     10000.0
6   Kim    22   South-Hall  45000.0
5   David  27   Texas       85000.0
4   Mark   25   Rich-Mond   65000.0
3   Teddy  23   Norway      20000.0
2   Allen  25   Texas       15000.0
1   Pual   32   California  20000.0
sqlite> select * from company order by salary asc, name desc;
id  name   age  address     salary 
--  -----  ---  ----------  -------
7   James  24   Houston     10000.0
2   Allen  25   Texas       15000.0
3   Teddy  23   Norway      20000.0
1   Pual   32   California  20000.0
6   Kim    22   South-Hall  45000.0
4   Mark   25   Rich-Mond   65000.0
5   David  27   Texas       85000.0
sqlite> select count(*), salary from company group by salary;
count(*)  salary 
--------  -------
1         10000.0
1         15000.0
2         20000.0
1         45000.0
1         65000.0
1         85000.0
sqlite> select count(*), salary from company having salary >= 20000 group by salary;
Error: in prepare, near "group": syntax error (1)
sqlite> select count(*), salary from company group by salary having salary >= 20000;
count(*)  salary 
--------  -------
2         20000.0
1         45000.0
1         65000.0
1         85000.0
sqlite> select distinct salary, * from company; 
salary   id  name   age  address     salary 
-------  --  -----  ---  ----------  -------
20000.0  1   Pual   32   California  20000.0
15000.0  2   Allen  25   Texas       15000.0
20000.0  3   Teddy  23   Norway      20000.0
65000.0  4   Mark   25   Rich-Mond   65000.0
85000.0  5   David  27   Texas       85000.0
45000.0  6   Kim    22   South-Hall  45000.0
10000.0  7   James  24   Houston     10000.0
```

## E. constrains
1. not null
2. default
3. unique
4. primary key
5. check

## F. join
There are 3 types:
1. cross join: *
2. inner join: and
3. outer join: or

```sql
sqlite> insert into department VALUES (1, 'IT Billing', 1 );
sqlite> insert into department VALUES (2, 'Engineering', 2 );
sqlite> insert into department VALUES (3, 'Finance', 7 );
sqlite> select * from department;
id  dept         emp_id
--  -----------  ------
1   IT Billing   1     
2   Engineering  2     
3   Finance      7    
sqlite> select * from company;
id  name   age  address     salary 
--  -----  ---  ----------  -------
1   Pual   32   California  20000.0
2   Allen  25   Texas       15000.0
3   Teddy  23   Norway      20000.0
4   Mark   25   Rich-Mond   65000.0
5   David  27   Texas       85000.0
6   Kim    22   South-Hall  45000.0
7   James  24   Houston     10000.0
sqlite> select emp_id, name, dept from company cross join department;
emp_id  name   dept       
------  -----  -----------
1       Pual   IT Billing 
2       Pual   Engineering
7       Pual   Finance    
1       Allen  IT Billing 
2       Allen  Engineering
7       Allen  Finance    
1       Teddy  IT Billing 
2       Teddy  Engineering
7       Teddy  Finance    
1       Mark   IT Billing 
2       Mark   Engineering
7       Mark   Finance    
1       David  IT Billing 
2       David  Engineering
7       David  Finance    
1       Kim    IT Billing 
2       Kim    Engineering
7       Kim    Finance    
1       James  IT Billing 
2       James  Engineering
7       James  Finance  
sqlite> select emp_id, name, dept from company inner join department on company.id = department.emp_id;emp_id  name   dept       
------  -----  -----------
1       Pual   IT Billing 
2       Allen  Engineering
7       James  Finance 
sqlite> select emp_id, name, dept from company left outer join department on company.id = department.emp_id;
emp_id  name   dept       
------  -----  -----------
1       Pual   IT Billing 
2       Allen  Engineering
        Teddy             
        Mark              
        David             
        Kim               
7       James  Finance 
```

## G. union

```sql
sqlite> select * from company where id > 3 union select * from company where salary >= 30000;
id  name   age  address     salary 
--  -----  ---  ----------  -------
4   Mark   25   Rich-Mond   65000.0
5   David  27   Texas       85000.0
6   Kim    22   South-Hall  45000.0
7   James  24   Houston     10000.0
sqlite> select * from company where id > 3 union all select * from company where salary >= 30000;
id  name   age  address     salary 
--  -----  ---  ----------  -------
4   Mark   25   Rich-Mond   65000.0
5   David  27   Texas       85000.0
6   Kim    22   South-Hall  45000.0
7   James  24   Houston     10000.0
4   Mark   25   Rich-Mond   65000.0
5   David  27   Texas       85000.0
6   Kim    22   South-Hall  45000.0
```
## H. alter
1) change name of a table. 2) add new column of a table.
```sql
sqlite> alter table test_company rename to testcompany;
sqlite> .tables
company      department   testcompany
sqlite> alter table testcompany add column sex char(1) default M not null;
sqlite> select * from testcompany;
id  name   age  address     salary   sex
--  -----  ---  ----------  -------  ---
2   Allen  25   Texas       15000.0  M  
3   Teddy  23   Norway      20000.0  M  
4   Mark   25   Rich-Mond   65000.0  M  
5   David  27   Texas       85000.0  M  
6   Kim    22   South-Hall  45000.0  M  
7   James  24   Houston     10000.0  M 
```

## I. functions
count() min() max() avg() sum() random() abs() upper() lower() length()

## J. index & indexed by

```sql
sqlite> create index salary_index on company (salary);
sqlite> .indexes
salary_index                   sqlite_autoindex_department_1
sqlite_autoindex_company_1   
sqlite> .indexes company 
salary_index                sqlite_autoindex_company_1
sqlite> select salary from company where salary > (select avg(salary) from company);
salary 
-------
45000.0
65000.0
85000.0
sqlite> explain query plan select salary from company where salary > (select avg(salary) from company);
QUERY PLAN
|--SEARCH company USING COVERING INDEX salary_index (salary>?)
`--SCALAR SUBQUERY 1
   `--SCAN company USING COVERING INDEX salary_index
sqlite> SELECT * FROM COMPANY INDEXED BY salary_index WHERE salary > 5000;
```

## K. view and transaction
```sql
sqlite> create view company_view as select id, name, age from company;
sqlite> .tables
company       company_view  department  
sqlite> select * from company_view 
   ...> ;
id  name   age
--  -----  ---
1   Paul   32 
2   Allen  25 
3   Teddy  23 
4   Mark   25 
5   David  27 
6   Kim    2

sqlite> begin;
sqlite> delete from company where age = 25;
sqlite> insert into company values (10, 'Bob', 35, 'China', 30000);
sqlite> commit;
sqlite> select * from company;
id  name   age  address     salary 
--  -----  ---  ----------  -------
1   Paul   32   California  20000.0
3   Teddy  23   Norway      20000.0
5   David  27   Texas       85000.0
6   Kim    22   South-Hall  45000.0
10  Bob    35   China       30000.0
```

## L. autoincrement

```sql
drop view company_view ;
sqlite> drop table company;
sqlite> create table company (
   ...> id int primary key autoincrement,
   ...> name text not null,
   ...> age int not null,
   ...> address char(50) default 'unknown',
   ...> salary real default 0
   ...> );
Error: in prepare, AUTOINCREMENT is only allowed on an INTEGER PRIMARY KEY (1)
sqlite> create table company (
   ...> id integer primary key autoincrement,
   ...> name text not null,
   ...> age int not null,
   ...> address char(50) default 'unknown',
   ...> salary real default 0
   ...> );
sqlite> insert into company (NAME,AGE,ADDRESS,SALARY) values ('Paul', 32, 'California', 20000.00);
sqlite> insert into company (NAME,AGE,ADDRESS,SALARY) values ('Allen', 25, 'Texas', 15000.00 );
sqlite> insert into company (NAME,AGE,ADDRESS,SALARY) values ('Teddy', 23, 'Norway', 20000.00 );
sqlite> insert into company (NAME,AGE,ADDRESS,SALARY) values ('Mark', 25, 'Rich-Mond', 65000.00);
sqlite> insert into company (NAME,AGE,ADDRESS,SALARY) values ('David', 27, 'Texas', 85000.00);
sqlite> select * from company;
id  name   age  address     salary 
--  -----  ---  ----------  -------
1   Paul   32   California  20000.0
2   Allen  25   Texas       15000.0
3   Teddy  23   Norway      20000.0
4   Mark   25   Rich-Mond   65000.0
5   David  27   Texas       85000.0
```
