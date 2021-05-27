## io.c
  ```
  - 错误提示信息：
    unknown message 0:0 [sender]
    rsync error: error in rsync protocol data stream (code 12) at io.c(436) [sender=3.0.6]
    
    - 根据提示信息可以看出
      1. 传输端 rsync版本 3.0.6
      2. 由于rsync协议数据流导致
      
    - 问题发生时的情况
      1. 接收端服务器换过服务器，但是用户的密码没有改变，ip没有变，其它一切未知
      
    - 检查的方向
      1. 接收端服务器的系统版本（Ubuntu）、rsync版本（3.1.0）
      2. 接收端服务器的日志(auth.log)
      3. 接收端服务器的网卡信息和路由信息
      
    - 解决办法
      1. 保持对端的rsync版本一致
      
  ```
