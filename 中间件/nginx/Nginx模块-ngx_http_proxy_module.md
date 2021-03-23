## ngx_http_proxy_module介绍
1. 作为反向代理，将请求转发到指定服务器

### 后端服务器nginx获取真实ip地址
```shell
1. 对代理服务器nginx.conf进行配置
]# vim /path/to/conf/nginx.conf
...
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    access_log  logs/access.log  main;
...
    location / {
        proxy_pass http://ip[:port];
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header REMOTE-HOST $remote_addr;
        proxy_set-header X-Forwarded-For $proxy_add-x-forwarded_for;
    }
```
> 参考文章
> 1. [阿里云SLB后NGINX、TOMCAT获取真实IP](https://www.jianshu.com/p/77d2fd957ab8)
