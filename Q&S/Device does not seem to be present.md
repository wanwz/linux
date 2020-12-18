## 问题描述
1.虚拟机添加一块网卡，重启服务器后，不显示ip信息
2.journalctl -xe 提示"Device does not seem to be present"

## 解决办法
1.在网卡配置文件中检查'NAME'和'DEVICE'值是否一致，不一致则修改一致
2.是否缺少'NAME'或'DEVICE'字段，缺少则添加
3.重启网卡

## 问题原因
1.ifcfg-ens*缺少'DEVICE'键值导致该问题
