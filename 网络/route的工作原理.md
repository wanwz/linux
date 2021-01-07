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
        ```
    2. Destination|Genmask 共同指定一个目标地址/目标网络地址段，即将报文发往的地方
    3. Gateway 网关，指定报文要经过这个网关发往目标地址
    
