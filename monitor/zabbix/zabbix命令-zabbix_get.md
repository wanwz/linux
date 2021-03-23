## zabbix命令使用方法
```
zabbix_get [-s|k|p]
  -s：客户端主机ip或主机名
  -p：客户端端口 默认10050
  -k：key值
```
### 示例
1. 获取客户端CPU的值
```shell
]# /usr/local/zabbix/bin/zabbix_get -s 192.168.153.130 -p 10050 -k system.cpu.load[all,avg1]
0.000
]# /usr/local/zabbix/bin/zabbix_agentd -p | grep system.cpu.load #如果不知道有哪些key可用，可用此命令查找
```
