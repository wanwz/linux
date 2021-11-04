## openvpn

> CentOS 7.3
>
> openvpn-2.4

### openvpn-client

1. 路径

   ```shell
   ]# cd /etc/openvpn
   ]# tree
   .
   ├── client
   │   ├── xxx.cer
   │   ├── xxx.key
   │   ├── client.conf
   │   └── xxx_ca.cer
   └── server
   ```

2. 文件

   1. xxx.cer|xxx.key 由openvpn-server提供

   2. client.conf

      ```shell
      ]# rpm -qa openvpn
      openvpn-2.4.9-1.el7.x86_64
      # rpm -ql openvpn-2.4.9-1.el7.x86_64 |grep client.conf
      /usr/share/doc/openvpn-2.4.9/sample/sample-config-files/client.conf
      /usr/share/doc/openvpn-2.4.9/sample/sample-config-files/roadwarrior-client.conf
      /usr/share/doc/openvpn-2.4.9/sample/sample-config-files/xinetd-client-config
      ]# cp/usr/share/doc/openvpn-2.4.9/sample/sample-config-files/client.conf /etc/openvpn/client/client.conf
      根据具体需求修改client.conf
      ```

3. sytemd

   ```shell
   ]# systemctl start openvpn-client@client
   ]# systemctl stop openvpn-client@client
   ```

4. 测试

   ```shell
   ]# ip a
   ...
   tun0: <POINTOPOINT,MULTICAST,NOARP,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UNKNOWN qlen 100
   inet xxx.xxx.xxx.xxx peer xxx.xxx.xxx.xxx/xx scope global tun0
   
   ]# ping xxx.xxx.xx.xxx 通则表示正常
   ```



> 启示：
>
> 	1. 在完全不了解一个服务时，去官网看它的介绍
>  	2. linux  man命令查看使用方法
>  	3. rpm -ql rpmpkg 查看其是否有sample样例和README
