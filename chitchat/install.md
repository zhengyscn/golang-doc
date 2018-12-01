# 部署测试

PostgreSQL数据库安装
centos:
```bash
yum install postgresql-server
yum install postgresql
```

ubuntu:
```
sudo add-apt-repository "deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main"
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
sudo apt-get update
sudo apt-get install postgresql-9.6
```

修改PostgreSQL数据库默认用户postgres数据库用户的密码为 manage
PostgreSQL数据库创建一个postgres用户作为数据库的管理员，密码随机，所以需要修改密码，方式如下：
打开客户端工具（psql）
```bash
sudo -u postgres psql
ALTER USER postgres WITH PASSWORD 'manage';
\q
```
修改ubuntu操作系统的postgres用户的密码（密码要与数据库用户postgres的密码相同）
```bash
sudo passwd -d postgres
sudo -u postgres passwd
```

修改pg的配置文件：
```text
# for centos file@/etc/postgresql/postgres.conf
# "local" is for Unix domain socket connections only
local   all             all                                     trust
# IPv4 local connections:
host    all             all             127.0.0.1/32           trust
```

```text
#for ubuntu  file@/etc/postgresql/9.6/main/pg_hba.conf
# "local" is for Unix domain socket connections only
local   all             all                                     trust
# IPv4 local connections:
host    all             all             127.0.0.1/32           trust
```

重启pg服务

创建测试数据库和用户
```
# psql -u postgres 
postgres=#  create user "chater" with password '123456';
postgres=#  create database "chitchat" with owner = "chater";
postgres=# \q
# psql -U chater -d chitchat
初始化数据表
```

启动 chitchat 程序测试