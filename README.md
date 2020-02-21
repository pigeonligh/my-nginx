# My Nginx

一个简单的可以远程修改的 Nginx，推荐运行在 Docker 中，基于 nginx:alpine 环境与 Golang 语言开发。

## 功能

- 简单的 HTTP/HTTPS 转发
- SSL 证书管理

## 安装

需要具有 Docker 环境。

先从 Docker Hub 拉取该项目的镜像：

```
docker pull pigeonligh/my-nginx
```

运行容器并映射端口：

```
docker run -d --restart=always -p 80:80 -p 443:443 -p 12345:8080 --name="my-nginx" -e MANAGE_TOKEN="HereIsPassword" pigeonligh/my-nginx
```

其中 `MANAGE_TOKEN` 为登录口令，`12345` 可以改为其他端口，该端口用于管理 Nginx 配置。

然后在浏览器访问 IP:12345 就能进行 Nginx 配置。

### 配置持久化

使用 Docker 的 Volume 功能引导 `/etc/nginx` 路径，即可将配置持久化，使用方法：

```
docker volume create my-nginx-volume
docker run -d --restart=always -p 80:80 -p 443:443 -p 12345:8080 \
    --name="my-nginx" \
    -v my-nginx-volume:/etc/nginx \
    -e MANAGE_TOKEN="HereIsPassword" pigeonligh/my-nginx
```

这样重新创建容器之后就可以保留配置了。

### 修改密码

暂不支持修改密码，可以删除容器并重新创建，因此推荐进行配置持久化，否则将会丢失数据。

删除容器：

```
docker stop my-nginx
docker rm my-nginx
```

## 转发配置

转发 HTTP 时，添加一个 HTTP 配置，填写域名，并将 Locations 中的 `127.0.0.1:8000` 改为需要转发到的位置。

例如，同样使用 Docker 部署一个其他的应用，容器在子网中的 IP 为 `172.17.0.100`，则将 `127.0.0.1:8000` 改为 `172.17.0.100` 即可，前后的字符不需要改变。

如果要从容器中访问主机的地址，则使用 `192.168.65.2`。

### 转发 HTTPS

只要将选项选为 HTTPS，并根据需要修改 SSLProtocols 和 SSLCiphers 即可。添加 HTTPS 需要在证书管理中配置相应的证书，否则应用配置将会失败！

### 默认转发

在域名中填写 `_` 即默认转发。（若要为 HTTPS 添加默认转发，也要添加域名为 `_` 的证书。）

### 不启用配置

若在配置中不选择启用，并不意味着不会进行转发，而是对于给请求返回 `403 Forbidden`。

## 证书管理

填写域名并选择 `crt` 和 `key` 文件，然后保存即可。系统会自动根据域名来匹配转发规则。

## 应用配置

所有的修改需要应用才会生效！所有的修改需要应用才会生效！所有的修改需要应用才会生效！

## 备注

本来还想做通用的端口转发，突然发现在容器里面要端口转发的话比较麻烦，监听更多端口还需要重新创建容器，想想应该也没什么用，就算了。

## MIT License

Copyright (c) 2020 pigeonligh