## 部署Google Authenticator 动态密码 + SSH 密码双重认证

### 环境
  1. epel源已安装
  2. 服务器时间准确
  3. 关闭selinux

### 安装
  ```
  ]# yum install google-authenticator -y
  ]# google-authenticator #按照提示输入"y"，在输出二维码时，使用谷歌身份验证器扫描记录，没有其它需要就一直"y"下去
  ]# vim /etc/pam.d/sshd
  #%PAM-1.0
  auth       required     pam_sepermit.so
  auth       substack     password-auth
  auth       required     pam_google_authenticator.so nullok #此行为新加的行，nullok
  auth       include      postlogin
  account    required     pam_nologin.so
  ...
  ]# vim /etc/ssh/sshd_config
  ...
  ChallengeResponseAuthentication yes #由"no"修改为"yes"
  
  ]# systemctl restart sshd
  ```

### 使用scrt验证登陆
  1. 连接会话选项把"Keyboard Interactive"放到最前面
  2. ![photo](https://github.com/wanwz/linux/blob/main/image/%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20210323150843.png)
