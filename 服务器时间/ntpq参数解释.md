## ntpq
1. 用来监视ntpd操作，ntpq -p查询网络中的NTP服务器，同时显示客户端和每个服务器的关系

|    标志     |                             含义                             |
| :---------: | :----------------------------------------------------------: |
|      *      |               响应的NTP服务器和最精确的服务器                |
|      +      |                 响应这个查询请求的NTP服务器                  |
| blank(空格) |                     没有响应的NTP服务器                      |
|   remote    |                响应这个请求的NTP服务器的名称                 |
|    refid    |             NTP服务器使用的更高一级服务器的名称              |
|     st      |                正在响应请求的NTP服务器的级别                 |
|    when     |                上一次成功请求之后到现在的秒数                |
|    poll     | 本地和远程服务器多少时间进行一次同步，单位秒，在一开始运行NTP的时候这个poll值会比较小，服务器同步的频率大，可以尽快调整到正确的时间范围，之后poll值会逐渐增大，同步的频率也就会相应减小 |
|    reach    | 用来测试能否和服务器连接，是一个八进制值，每成功连接一次它的值就会增加 |
|    delay    |          从本地机发送同步要求到ntp服务器的往返时间           |
|   offset    | 主机通过NTP时钟同步与所同步时间源的时间偏移量，单位为毫秒，offset越接近于0，主机和ntp服务器的时间越接近 |
|   jitter    | 统计了在特定个连续的连接数里offset的分布情况。简单地说这个数值的绝对值越小，主机的时间就越精确 |

2. ntpq -p 报错提示"timed out, nothing received"
    1. 重启ntpd服务后，立即执行ntpq -p会报上述错误，等待5分钟后再次执行即可
    2. 参考[ntpq_timed_out_nothing_received.md]: https://github.com/huataihuang/cloud-atlas-draft/blob/master/service/ntp/ntpq_timed_out_nothing_received.md
  
