## ntpdate
1. ntpdate -B -p 8 [serverip|hostname] (相当于软调整，步进很短)
```shell
]# ntpdate -B -p 8 0.centos.pool.ntp.org
30 Dec 12:10:16 ntpdate[2667]: adjust time server 10.255.243.181 offset 757.644670 sec
]# ntpdate -B -p 8 0.centos.pool.ntp.org
30 Dec 12:10:33 ntpdate[2670]: adjust time server 10.255.243.181 offset 757.641274 sec
```
    - 测试结果
        1. 时差超过半个小时，直接报错"Can't adjust the time of day"
        2. 时差在半小时以内，每次执行上述命令，只会缩减特别小的时间差
2. ntpdate [serverip|hostname]
```shell
]# ntpdate 0.centos.pool.ntp.org
30 Dec 12:17:44 ntpdate[2646]: step time server 10.255.243.181 offset 173825.539362 sec
```
    - 测试结果
        1. 不加任何参数，执行此命令后会直接同步时间 (相当于硬调整，一步到位)
3. 使用条件
    1. 指定的serverip|hostname需启动ntpd服务
    2. 该服务器需关闭ntpd服务
