## 部署zabbix.5.0.7

### 环境
1. CentOS7.4
2. 清华yum源和epel源

### 安装步骤
> zabbix5.0版本安装竟然这么简单...
#### 服务器端配置
```shell
]# yum install mariadb mariadb-server -y

]# rpm -Uvh https://repo.zabbix.com/zabbix/5.0/rhel/7/x86_64/zabbix-release-5.0-1.el7.noarch.rpm -y
]# yum install zabbix-server-mysql zabbix-agent -y
]# yum install centos-release-scl -y
]# vim /etc/yum.repos.d/zabbix.repo #修改zabbix-frontend配置
...
[zabbix-frontend]
enabled=1
...
]# yum install zabbix-web-mysql-scl zabbix-apache-conf-scl -y

]# mysql_secure_installation #初始化数据库一路回车即可（仅限第一次安装数据库进行此操作）
]# mysql -uroot -p
mysql> create database zabbix character set utf8 collate utf8_bin;
mysql> create user zabbix@localhost identified by 'password';
mysql> grant all privileges on zabbix.* to zabbix@localhost;
mysql> quit;
]# zcat /usr/share/doc/zabbix-server-mysql*/create.sql.gz | mysql -uzabbix -p zabbix

]# vim /etc/zabbix/zabbix_server.conf
...
DBPassword=password
]# vim /etc/opt/rh/rh-php72/php-fpm.d/zabbix.conf #修改时区
...
php_value[date.timezone] = Asia/Shanghai

]# systemctl restart zabbix-server zabbix-agent httpd rh-php72-php-fpm
]# systemctl enable zabbix-server zabbix-agent httpd rh-php72-php-fpm

]# yum install wqy-microhei-fonts -y #解决图形数据显示乱码问题
]# cp /usr/share/fonts/wqy-microhei/wqy-microhei.ttc /usr/share/fonts/dejavu/DejaVuSans.ttf
```
#### web配置
1. 登录http://server_ip_or_name/zabbix
2. 默认用户名：Admin/zabbix
