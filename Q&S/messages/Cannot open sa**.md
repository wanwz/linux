## 问题现象
root总是收到邮件提醒如下内容：
"Subject: Cron <root@manager> /usr/lib/sa/sa1 1 1
Content-Type: text/plain; charset=UTF-8
X-Cron-Env: <SHELL=/bin/sh>
X-Cron-Env: <HOME=/root>
X-Cron-Env: <PATH=/usr/bin:/bin>
X-Cron-Env: <LOGNAME=root>
X-Cron-Env: <USER=root>

Cannot open /var/log/sa/sa20: No such file or directory"

## 解决办法
1.创建sa目录
```shell
]# mkdir /var/log/sa
```
2.手动生成sa文件
```shell
]# cd /var/log/sa
]# sar -o 20 &>/dev/null &
```
3.查看生成的信息
```shell
]# sar -f sa20
Linux x.x.xx.x.CentOSx.x.xx (hostname)     12/20/2020

02:02:29 PM       CPU     %user     %nice   %system   %iowait     %idle
02:02:49 PM       all      0.02      0.00      0.02      0.03     99.94
Average:          all      0.02      0.00      0.02      0.03     99.94
```
## 问题原因
执行定时任务sysstat，但没有sa文件
```shell
]# cat /etc/cron.d/sysstat
# run system activity accounting tool every 10 minutes
*/10 * * * * root /usr/lib/sa/sa1 1 1
# generate a daily summary of process accounting at 23:53
53 23 * * * root /usr/lib/sa/sa2 -A
```

> 如果不需要系统监控信息，也可以将其注释，取消定时执行
