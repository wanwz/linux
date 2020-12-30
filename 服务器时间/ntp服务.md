## 概述
1. ntpd是一个守护进程
2. 用于校对本地系统和时钟源服务器之间的时间
3. 步进式的逐渐校正时间

## 安装
```shell
]# rpm -qa |grep ntp #有以下输出则表示已安装，无回显则yum安装
fontpackages-filesystem-1.44-8.el7.noarch
ntpdate-4.2.6p5-29.el7.centos.2.x86_64
ntp-4.2.6p5-29.el7.centos.2.x86_64
]# yum install ntp -y
```

## 配置
```shell
]# grep -Ev "^#|^$" /etc/ntp.conf
driftfile /var/lib/ntp/drift #系统与BIOS时间偏差记录
restrict default nomodify notrap nopeer noquery #默认拒绝client所有操作
restrict 127.0.0.1  #允许本地一切操作
restrict ::1
includefile /etc/ntp/crypto/pw
keys /etc/ntp/keys
disable monitor #增强安全性
server time1.tencentyun.com iburst #时钟源服务器地址
server time2.tencentyun.com iburst
server time3.tencentyun.com iburst
server time4.tencentyun.com iburst
server time5.tencentyun.com iburst
interface ignore wildcard #增强安全性
interface listen eth0 #增强安全性
```

## 启动
```shell
]# systemctl start ntpd
]# systemctl enable ntpd
```


## NTP服务器搭建(作为时钟源服务器)
> 主要是修改配置文件，安装启动与上述相同
1. 修改配置文件
```shell
]# cp /etc/ntp.conf{,.bak}
]# vim /etc/ntp.conf
driftfile /var/lib/ntp/drift #系统与BIOS时间偏差记录
restrict default nomodify notrap nopeer noquery #默认拒绝client所有操作
restrict 127.0.0.1  #允许本地一切操作
restrict 192.168.9.0 mask 255.255.255.0 nomodify #允许192.168.9.0网段的client连接到这台服务器同步时间，但是拒绝client修改时间
#restrict 0.0.0.0 mask 0.0.0.0 nomodify 允许所有的client连接到这台服务器同步时间，但是拒绝client修改时间

server 0.centos.pool.ntp.org
server 1.centos.pool.ntp.org
server 2.centos.pool.ntp.org

server 127.127.1.0
fudge 127.127.1.0 stratum 10
includefile /etc/ntp/crypto/pw
keys /etc/ntp/keys

]# vim /etc/sysconfig/ntpd
SYNC_HWCLOCK=yes #将同步好的系统时间写入BIOS时间
```
## 注意
1. ntpd默认为客户端运行方式，不作为时钟源服务器，主要看配置文件
2. ntp服务器端重启后，客户端需等待5-10分钟再同步，否则提示"no server suitable for synchronization found"错误
3. client使用ntpdate时，一定要关闭ntp服务，否则提示端口被占用
