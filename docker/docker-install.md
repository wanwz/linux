## 阿里云yum安装
### centos7
    1. yum install -y yum-utils device-mapper-persistent-data lvm2
    2. yum-config-manager --add-repo https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
    3. sed -i 's+download.docker.com+mirrors.aliyun.com/docker-ce+' /etc/yum.repos.d/docker-ce.repo
    4. yum makecache fast && yum -y install docker-ce docker-ce-cli containerd.io
    5. service docker start
    6. docker version #安装校验
    
    
