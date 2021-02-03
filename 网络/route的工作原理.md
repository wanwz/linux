## Linux系统路由的工作原理

### route
    1. route -n
        ```shell
        ]# route -n
        Kernel IP routing table
        Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
        0.0.0.0         172.21.0.1      0.0.0.0         UG    0      0        0 eth0
        169.254.0.0     0.0.0.0         255.255.0.0     U     1002   0        0 eth0
        172.21.0.0      0.0.0.0         255.255.240.0   U     0      0        0 eth0
        
        Flags标志说明：
        U Up表示此路由当前为开启状态
        H Host表示此网关为一主机
        G Gateway表示此网关为一路由器
        R Reinstate Route使用动态路由重新初始化的路由
        D Dynamically此路由是动态性的写入
        M Modified此路由是由路由守护程序或导向器动态修改
        ！表示此路由当前为关闭状态
        ```
    2. Destination|Genmask 共同指定一个目标地址/目标网络地址段，即将报文发往的地方
    3. Gateway 网关，指定报文要经过这个网关发往目标地址

### 命令参数
    1. route add -net 192.168.0.0 netmask 255.255.0.0 dev eth0
    2. route add -host 192.168.1.2 dev eth1
        1. add 添加路由
        2. -net 添加网络段
        3. -host 添加主机地址
        4. dev 指定网卡
        5. netmask 子网掩码
    3. route del -net 192.168.0.0 netmask 255.255.0.0 dev eth0
        1. del 删除路由
    4. route add -net 192.168.1.0 netmask 255.255.255.0 reject
        1. reject 屏蔽路由
    5. route add default gw 192.168.1.1
        1. 添加默认网关
