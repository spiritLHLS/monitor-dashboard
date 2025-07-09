# 展示图

https://github.com/spiritLHLS/monitor-dashboard/blob/master/%E5%B1%95%E7%A4%BA.md

# agent

https://github.com/spiritLHLS/monitor-agent

# 环境

go: 1.22.7

mysql: 8.0

redis: 7.4.0

nginx: 1.24.0

## 编译

修补环境

```
apt-get install xdg-utils -y
```

server 下执行

```macos
GOARCH=amd64 GOOS=linux go build
```

```cmd
set GOARCH=amd64
set GOOS=linux
go build
```

或

```powershell
$env:GOARCH="amd64"
$env:GOOS="linux"
go build
```

生成 server 文件

web 下执行

```
npm run build
```

生成 dist 文件夹

## 部署

请在[mysqld]部分下添加以下行：

```
bind-address=0.0.0.0
expire_logs_days=3
max_binlog_size=512M
```

强行启用IPV4监听和允许SSH登录MYSQL，且设置日志保存时长和大小

执行

```
mysql -u root -p
```

登录后执行

```
#使用mysql
use mysql
#更新root用户权限，“%”指的是所有地址都可以访问
update user set Host='%' where User='root';
flush privileges; 
```

然后执行

```
quit
```

退出登录

nginx 配置

```
location /api {
        rewrite ^/api/(.*)$ /$1 break;  #重写
        proxy_pass http://127.0.0.1:8888;
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
}
```

## 后台挂起执行

```bash
nano /etc/systemd/system/ecs-spider.service
```

```ini
[Unit]
Description=ECS Spider Service
After=network.target

[Service]
Type=simple
ExecStart=/ecs-spiders/server
WorkingDirectory=/ecs-spiders
Restart=always
User=root
Group=root
Environment=PATH=/usr/bin:/usr/local/bin
Environment=NODE_ENV=production

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl daemon-reload
sudo systemctl start ecs-spider.service
```

```
sudo systemctl status ecs-spider.service
```

```
sudo systemctl enable ecs-spider.service
```

## 重启

```
sudo systemctl stop ecs-spider.service
```

```
sudo systemctl restart ecs-spider.service
```
