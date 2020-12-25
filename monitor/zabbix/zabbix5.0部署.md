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

#### 添加邮件监控
1. 配置 --> 动作
2. 创建动作
3. 动作
  1. 名称：服务器异常通知 --> 添加
  2. 操作：操作下 --> 添加 --> 用户群组（Zabbix administrators） --> 仅送到(admin) --> Custom message(勾上) --> 主题（恢复{TRIGGER.STATUS}, 服务器:{HOSTNAME1}: {TRIGGER.NAME}已恢复!） --> 消息（告警主机:{HOSTNAME1} 告警时间:{EVENT.DATE} {EVENT.TIME} 告警等级:{TRIGGER.SEVERITY} 告警信息: {TRIGGER.NAME} 告警项目:{TRIGGER.KEY1} 问题详情:{ITEM.NAME}:{ITEM.VALUE} 当前状态:{TRIGGER.STATUS}:{ITEM.VALUE1} 事件ID:{EVENT.ID}）--> Update
  3. 操作：恢复操作下 --> 添加 --> 用户群组（Zabbix administrators） --> 仅送到(admin) --> Custom message(勾上) --> 主题（恢复{TRIGGER.STATUS}, 服务器:{HOSTNAME1}: {TRIGGER.NAME}已恢复!）--> 消息（告警主机:{HOSTNAME1} 告警时间:{EVENT.DATE} {EVENT.TIME} 告警等级:{TRIGGER.SEVERITY} 告警信息: {TRIGGER.NAME} 告警项目:{TRIGGER.KEY1} 问题详情:{ITEM.NAME}:{ITEM.VALUE} 当前状态:{TRIGGER.STATUS}:{ITEM.VALUE1} 事件ID:{EVENT.ID}） --> Update
  4. 更新
