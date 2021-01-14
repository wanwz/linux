## logrotate
    1.linux系统自带此工具
    2.对日志进行自动切割轮转保存

### 工作原理
    1.执行顺序：crond --> /etc/cron.daily/logrotate --> /etc/logrotate.conf --> /etc/logrotate.d/*
    2.logrotate依托crontab来进行自动切割保存等操作，默认以天为单位(也可根据实际需求进行修改)
    3.配置文件首先以/etc/logrotate.conf为准，其次在/etc/logrotate.d/下可以更加细粒度的配置，如果与父配置文件相冲突的配置，以子配置文件为准
    4.若规定每周轮转切割一次，后面有规定日志大小达到1M才会切割，则以后面日志大小为准
    
### 其它
    1.logrotate -f /etc/logrotate.d/* 强制按照某子配置文件进行一次日志轮转
    2.logrotate -d /etc/logrotate.d/ 检查配置文件是否有错误
    3.其它重要参数：
        ```doc
        compress                        通过gzip 压缩转储以后的日志
        nocompress                      不做gzip压缩处理
        copytruncate                    用于还在打开中的日志文件，把当前日志备份并截断；是先拷贝再清空的方式，拷贝和清空之间有一个时间差，可能会丢失部分日志数据。
        nocopytruncate                  备份日志文件不过不截断
        create mode owner group         轮转时指定创建新文件的属性，如create 0777 nobody nobody
        nocreate                        不建立新的日志文件
        delaycompress                   和compress 一起使用时，转储的日志文件到下一次转储时才压缩
        nodelaycompress                 覆盖 delaycompress 选项，转储同时压缩。
        missingok                       如果日志丢失，不报错继续滚动下一个日志
        errors address                  专储时的错误信息发送到指定的Email 地址
        ifempty                         即使日志文件为空文件也做轮转，这个是logrotate的缺省选项。
        notifempty                      当日志文件为空时，不进行轮转
        mail address                    把转储的日志文件发送到指定的E-mail 地址
        nomail                          转储时不发送日志文件
        olddir directory                转储后的日志文件放入指定的目录，必须和当前日志文件在同一个文件系统
        noolddir                        转储后的日志文件和当前日志文件放在同一个目录下
        sharedscripts                   运行postrotate脚本，作用是在所有日志都轮转后统一执行一次脚本。如果没有配置这个，那么每个日志轮转后都会执行一次脚本
        prerotate                       在logrotate转储之前需要执行的指令，例如修改文件的属性等动作；必须独立成行
        postrotate                      在logrotate转储之后需要执行的指令，例如重新启动 (kill -HUP) 某个服务！必须独立成行
        daily                           指定转储周期为每天
        weekly                          指定转储周期为每周
        monthly                         指定转储周期为每月
        rotate count                    指定日志文件删除之前转储的次数，0 指没有备份，5 指保留5 个备份
        dateext                         使用当期日期作为命名格式
        dateformat .%s                  配合dateext使用，紧跟在下一行出现，定义文件切割后的文件名，必须配合dateext使用，只支持 %Y %m %d %s 这四个参数
        [size|minsize] log-size         当日志文件到达指定的大小时才转储，log-size能指定bytes(缺省)及KB(sizek)或MB(sizem)，当日志文件 >= log-size 的时候就转储。以下为合法格式：（其他格式的单位大小写没有试过）
                                        size = 5 或 size 5 （>= 5 个字节就转储）
                                        size = 100k 或 size 100k
                                        size = 100M 或 size 100M
        ```
