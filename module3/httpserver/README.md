# Mac 下构建 Linux 可执行程序
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# 作业
## 构建本地镜像
## 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化
docker build -t kleven2020/httpserver:1.1 .

## 将镜像推送至 docker 官方镜像仓库
docker push kleven2020/httpserver:1.1

## 通过 docker 命令本地启动 httpserver
### 拉去镜像
docker pull kleven2020/httpserver:1.1

### 查看镜像
docker images
|REPOSITORY|TAG|IMAGE ID|CREATED|SIZE|
|----|----|----|----|----|
|kleven2020/httpserver|1.1|5f4dbee51662|20 minutes ago|83.9MB|

### 运行镜像
docker run -d -p 8080:80 --name httpserver 5f4dbee51662
5019b11e44bd0425eeb63660ef41351065b974412d5327b68207a6eebf1f5a60

## 通过 nsenter 进入容器查看 IP 配置

### 1. 
lsns -t net
|NS|TYPE|NPROCS|PID|USER|NETNSID|NSFS|COMMAND|
|----|----|----|----|----|----|----|----|
|4026531992|net|122|1|root|unassigned||/sbin/init nopti|
|4026532257|net|2|1879|root|0|/run/docker/netns/b018ddd6414a| /bin/sh -c ./httpserver|
    
                                       
        

### 2. 
nsenter -t 1879 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
4: eth0@if5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
