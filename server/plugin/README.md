# 设置consul

ufw设置8000-8499 8501-9000 8500 的tcp/udp放行

如何安装：https://developer.hashicorp.com/consul/install?product_intent=consul

创建 Systemd 服务文件

```sh
sudo nano /etc/systemd/system/consul.service
```

然后填入以下内容：

```ini
[Unit]
Description=Consul Agent
After=network.target

[Service]
ExecStart=/usr/bin/consul agent -server -bootstrap-expect=1 -data-dir=/tmp/consul -ui \
  -bind=103.251.166.79 \
  -client=0.0.0.0 \
  -advertise=103.251.166.79 \
  -config-dir=/etc/consul.d/
Restart=always
User=root

[Install]
WantedBy=multi-user.target
```

重新加载 Systemd 配置

```sh
sudo systemctl daemon-reload
```

设置 Consul 开机自启

```sh
sudo systemctl enable consul
```

启动 Consul

```sh
sudo systemctl start consul
```

查看状态

```sh
sudo systemctl status consul
```

**停止**
```sh
sudo systemctl stop consul
```
  
**重启**

```sh
sudo systemctl restart consul
```

配置Token控制

```shell
echo -e '\nacl {\n  enabled = true\n  default_policy = "deny"\n  down_policy = "extend-cache"\n  enable_token_persistence = true\n}' | sudo tee -a /etc/consul.d/consul.hcl
```

```shell
consul acl bootstrap
```

记下SecretID为 ```689cc548-775e-1b08-d23a-d34c3562f90c```用于```http://ipv4:8500/ui```的登录和控制

```
AccessorID:       1e2b478f-10a1-4a80-83cc-61875f0a0bd5
SecretID:         689cc548-775e-1b08-d23a-d34c3562f90c
Description:      Bootstrap Token (Global Management)
Local:            false
Create Time:      2025-02-05 09:50:41.584760222 +0100 CET
Policies:
   00000000-0000-0000-0000-000000000001 - global-management
```
