## 关于服务器安全-ssh

### 修改默认ssh端口
  ```
  ]# echo "Port xxxx" >> /etc/ssh/sshd_config
  ```

### 禁止root远程登陆
  ```
  ]# echo "PermitRootLogin no" >> /etc/ssh/sshd_config
  ```

### 如果服务器除固定ip登录外，其余均不允许登陆
  ```
  ]# vim /etc/hosts.deny
  ...
  sshd:ALL
  
  ]# vim /etc/hosts.allow
  ...
  sshd:192.168.xx.xx:allow
  sshd:192.168.xx.xx:allow
  ...
  ```
