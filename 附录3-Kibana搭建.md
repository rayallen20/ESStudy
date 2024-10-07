# 附录3-Kibana搭建

## PART1. 安装Kibana

### 1.1 添加ES的GPG密钥和APT源

以下操作3个节点都做:

```
root@es-node-1:~# wget -qO - https://artifacts.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add -
Warning: apt-key is deprecated. Manage keyring files in trusted.gpg.d instead (see apt-key(8)).
OK
```

```
root@es-node-1:~# echo "deb https://artifacts.elastic.co/packages/8.x/apt stable main" | sudo tee -a /etc/apt/sources.list.d/elastic-8.x.list
deb https://artifacts.elastic.co/packages/8.x/apt stable main
```

### 1.2 更新包列表并安装Kibana

以下操作3个节点都做:

```
root@es-node-1:~# apt update
```

```
root@es-node-1:~# apt install kibana
```

## PART2. 配置Kibana

以下操作3个节点都做:

```
root@es-node-1:~# vim /etc/kibana/kibana.yml
```

- 修改`server.port`的值为: `5601`
- 修改`server.host`的值为: "本机IP"
- 修改`elasticsearch.hosts`的值为: `["https://192.168.1.195:9200", "https://192.168.1.196:9200", "https://192.168.1.197:9200"]`

```
root@es-node-1:~# cat /etc/kibana/kibana.yml |grep server.port
server.port: 5601
root@es-node-1:~# cat /etc/kibana/kibana.yml |grep server.host
server.host: "192.168.1.195"
root@es-node-1:~# cat /etc/kibana/kibana.yml |grep elasticsearch.hosts
elasticsearch.hosts: ["http://192.168.1.195:9200", "http://192.168.1.196:9200", "http://192.168.1.197:9200"]
```

## PART3. 启动Kibana

以下操作3个节点都做:

```
root@es-node-1:~# systemctl start kibana.service 
root@es-node-1:~# systemctl enable kibana
Created symlink /etc/systemd/system/multi-user.target.wants/kibana.service → /lib/systemd/system/kibana.service.
```

