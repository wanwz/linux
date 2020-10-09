## nginx-源代码编译部署

## 脚本

```shell
]# vim /scripts/nginx_install.sh
#!/usr/bin/env bash
# Author: wanwz
# Date: 2020-08-03
# Comment: 通过脚本自动安装Nginx
set -euo pipefail

touch /opt/nginx_install.log
ERROR_LOG=/opt/nginx_install.log

echo "******preinstall******"
yum install gcc gcc-c++ pcre pcre-devel zlib zlib-devel openssl openssl-devel -y &>${ERROR_LOG}
if [ $? == 0 ];then
    echo "依赖包已安装完成！"
else
		echo "安装失败，请检查错误日志！"
fi
mkdir -p /opt/nginx && cd /opt
set -x
wget http://nginx.org/download/nginx-1.18.0.tar.gz &>${ERROR_LOG} && tar xf nginx-1.18.0.tar.gz &>${ERROR_LOG}
set +x
echo "******preinstall is ok*****"

echo "******install nginx******"
cd nginx-1.18.0.tar.gz
set -x
./configure --prefix=/opt/nginx --with-http_stub_status_module && make && make install
set +x

if [ $? == 0 ];then
		echo "******Nginx已安装完成！******"
else
		echo "------Nginx安装失败！------"
fi
]# bash -n nginx_install.sh #检查语法是否有错误
]# bash nginx_install.sh #无误则执行脚本
```

