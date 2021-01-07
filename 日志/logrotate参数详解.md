## logrotate日志参数详解
    1. 程序产生日志后，通过logrorate来进行自动切割及保留
    2. 新的配置文件尽量保存在/etc/logrotate.d/目录下，不动主配置文件
### 参数
    ```shell
    ]# cat /etc/logrotate.d/syslog # 以此为例
    /var/log/cron           # 日志的路径
    /var/log/maillog
    /var/log/messages
    /var/log/secure
    /var/log/spooler
    {
            sharedscripts   # 共享脚本
            dateext         # 转储文件以日期来结束
            rotate 25       # 保留文件份数
            size 40M        # 超过40M，切换日志
            compress        # 压缩
            dateformat  -%Y%m%d%s   # 日志切换后保留格式 /var/log/messages-20210101
            postrotate              # 转换后，执行以下脚本一次
                    /bin/kill -HUP `cat /var/run/syslogd.pid 2> /dev/null` 2> /dev/null || true
            endscript               # 结束
    }
    
    # 其它参数
    daily|weekly|monthly    # 每日|周|月切换一次
    create 0600 root root   # 切换后创建新文件的权限，属主，属组
    minsize 1M              # 日志大于1M才会切换
    missingok               # 日志丢失，不报错
    compress                # 日志压缩
    delaycompress           # 和compress 一起使用时，转储的日志文件到下一次转储时才压缩
    nodelaycompress         # 覆盖 delaycompress 选项，转储同时压缩
    notifempty              # 当日志文件为空时，不进行轮转
    mail address            # 把转储的日志文件发送到指定的E-mail 地址
    olddir directory        # 转储后的日志文件放入指定的目录，必须和当前日志文件在同一个文件系统
    ```
### 举例
    1. 
