## ngx_http_ssl_module简介
为https提供支持
## ngx_http_ssl_module参数解释
1. ssl on|off;
2. ssl_certificate file; #当前虚拟主机使用PEM格式的证书文件
3. ssl_certificate_key file; #当前虚拟主机上与其证书匹配的私钥文件
4. ssl_protocols [SSLv2] [SSLv3] [TLSv1] [TLSv1.1] [TLSv1.2]; #支持ssl协议版本，默认为后三个
5. ssl_session_cache off|none| [builtin[:size]] | [shared:name:size];
      1. builtin[:size]：使用OpenSSL内建的缓存，此缓存为每个worker进程私有
      2. shared:name:size：在各worker之间使用一个共享的缓存
6. ssl_session_timeout time;客户端一侧的连接可以复用ssl_session_cache中缓存的ssl参数的有效时长
## 配置示例：
```shell
server {
      listen 443 ssl;
      servername www.xxx.com;
      root /opt/nginx/html;
      ssl on;
      ssl_certificate /opt/nginx/ssl/nginx.crt;
      ssl_certificate_key /opt/nginx/ssl/nginx.key;
      ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
      ssl_session_cache shared:sslcache:20m;
}
```
