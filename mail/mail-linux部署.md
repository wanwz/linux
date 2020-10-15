## mail-linux部署mail

### 安装

```shell
]# yum install mailx
```

### 修改配置文件

```shell
]# vim /etc/mail.rc
...
set from=xxx@163.com
set smtp=smtp.163.com
set smtp-auth-user=xxx@163.com
set smtp-auth-password=授权码
set smtp-auth=login
```

### 测试

```shell
]# echo "这是一封测试邮件！" | mail -s "[TEST]测试" xxx@163.com
```

### 扩展

1. 包含正文

   ```shell
   ]# mail -s "Title" xxx@163.com
   ```

   

2. 不包含正文

   ```shell
   ]# mail -s "Title" xxx@163.com < /root/mail_out.txt
   or
   ]# mail_txt="这是一封测试邮件！"
   ]# echo "$mail_txt" | mail -s "Title" xxx@163.com
   ```

   

3. 包含附件

   ```shell
   ]# mail -s "Title" xxx@163.com -a /root/mail_out.txt
   ```

   
