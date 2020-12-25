## zabbix报错

### web登录页面显示"Database Error..."
#### 解决办法
1. 检查数据库登录权限
```shell
]# mysql -uzabbixuser -p
mysql> show databases;
```
2. 检查zabbix_sever.conf文件
```shell
]# vim /etc/zabbix/zabbix_server.conf
...
DBHost=localhost
DBName=zabbix
DBUser=zabbixuser
DBPassword=password
...
```
3. 检查php的zabbix.conf文件
```shell
]# cat /etc/zabbix/web/zabbix.conf.php
...
$DB['TYPE']                             = 'MYSQL';
$DB['SERVER']                   = 'localhost';
$DB['PORT']                             = '3306';
$DB['DATABASE']                 = 'zabbix';
$DB['USER']                             = 'zabbixuser';
$DB['PASSWORD']                 = 'password';
```
> 目前到第三个步骤检查没有错误，即可正常登录
