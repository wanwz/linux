# rsync同步文件失败

## 问题描述

### 现象

1. 报错信息

   ```shell
   rsync error: error in rsync protocol data stream (code 12) at io.c(436) [sender=3.0.6]
   ```

2. 传输端执行的命令

   ```
   rsync --timeout=300 -avzc --delete-after --partial '-e ssh -p 22' /path/to/dir root@域名:/path/to/save/ >> /path/to/log 2>&1
   ```

3. 补充信息

   1. 接收端服务器：rsync版本3.1.0|Ubuntu


## 解决办法

1. 保持两端rsync版本一致后解决（因传输端一对多，故将接收端降低版本到3.0.6）

## 扩展

- rsync参数

  ```shell
  --delete-after:传输结束后删除不存在于源目标的文件
  --partial:允许恢复中断的传输，不使用这个参数时，rsync会删除传输被打断后已经传输的文件；，应与--append(指定文件接着上次传输中断的地方接着传)/--append-verify(与append类似，传输完成后会对文件进行一次校验，若校验失败，则重新发送整个文件)配合使用
  
  -e 按规范使用应该在'ssh -p port'外面
  -c rsync默认只检查文件的大小和最后修改日期是否变化，若发生变化，则重新传输；使用这个参数后，通过判断文件内容的校验和，决定是否重新传输
  ```
