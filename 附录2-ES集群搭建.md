# 附录2-ES集群搭建

## PART1. 服务器环境

4核16G 80G硬盘虚拟机3台

|主机名|IP|
|:-:|:-:|
|es-node-1|192.168.1.195|
|es-node-2|192.168.1.196|
|es-node-3|192.168.1.197|

## PART2. 安装JAVA

3个节点都要做:

```
root@es-node-1:~# apt install openjdk-17-jdk -y
```

## PART3. 安装ES

### 3.1 添加ES的GPG密钥

3个节点都要做:

```
root@es-node-1:~# wget -qO - https://artifacts.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add -
```

### 3.2 添加ES源

3个节点都要做:

```
root@es-node-1:~# sudo sh -c 'echo "deb https://artifacts.elastic.co/packages/8.x/apt stable main" > /etc/apt/sources.list.d/elastic-8.x.list'
```

### 3.3 更新并安装ES

3个节点都要做:

```
root@es-node-1:~# apt update
```

```
root@es-node-1:~# apt install elasticsearch -y
```

## PART4. 配置ES

### 4.1 配置节点的存储路径

#### a. 创建数据存储目录

3个节点都要做:

```
root@es-node-1:~# mkdir -p /mnt/new_partition/elasticsearch
root@es-node-1:~# chown -R elasticsearch:elasticsearch /mnt/new_partition/elasticsearch
root@es-node-1:~# chmod -R 0755 /mnt/new_partition/elasticsearch
```

#### b. 编辑ES配置文件

3个节点都要做:

```
root@es-node-1:~# vim /etc/elasticsearch/elasticsearch.yml 
```

将其中的`path.data`项的值修改为`/mnt/new_partition/elasticsearch`

```
root@es-node-1:~# cat /etc/elasticsearch/elasticsearch.yml|grep path.data
path.data: /mnt/new_partition/elasticsearch
```

## PART5. 配置ES集群

### 5.1 设置集群名称

3个节点都要做:

```
root@es-node-1:~# vim /etc/elasticsearch/elasticsearch.yml 
root@es-node-1:~# cat /etc/elasticsearch/elasticsearch.yml|grep cluster.name
cluster.name: my-es-cluster
```

将其中的`cluster.name`项的值修改为`my-es-cluster`

### 5.2 设置节点角色

#### a. 设置节点名称

这一步分别将`node.name`项的值修改为各自的主机名,**要求各个节点之间该配置项的值唯一**

```
root@es-node-1:~# vim /etc/elasticsearch/elasticsearch.yml 
root@es-node-1:~# cat /etc/elasticsearch/elasticsearch.yml|grep node.name
node.name: es-node-1
```

```
root@es-node-2:~# vim /etc/elasticsearch/elasticsearch.yml
root@es-node-2:~# cat /etc/elasticsearch/elasticsearch.yml|grep node.name
node.name: es-node-2
```

```
root@es-node-3:~# vim /etc/elasticsearch/elasticsearch.yml 
root@es-node-3:~# cat /etc/elasticsearch/elasticsearch.yml|grep node.name
node.name: es-node-3
```

### 5.3 设置节点网络

这一步各个节点指定自己的IP地址即可:

```
root@es-node-1:~# vim /etc/elasticsearch/elasticsearch.yml 
root@es-node-1:~# cat /etc/elasticsearch/elasticsearch.yml|grep network.host
network.host: 192.168.1.195
```

```
root@es-node-2:~# vim /etc/elasticsearch/elasticsearch.yml
root@es-node-2:~# cat /etc/elasticsearch/elasticsearch.yml|grep network.host
network.host: 192.168.1.196
```

```
root@es-node-3:~# vim /etc/elasticsearch/elasticsearch.yml 
root@es-node-3:~# cat /etc/elasticsearch/elasticsearch.yml|grep network.host
network.host: 192.168.1.197
```

### 5.4 配置跨节点通信

#### a. 指定集群中的其他节点

这一步3个节点将`discovery.seed_hosts`配置项的值设置为相同的即可:

```
root@es-node-1:~# vim /etc/elasticsearch/elasticsearch.yml 
root@es-node-1:~# cat /etc/elasticsearch/elasticsearch.yml|grep discovery.seed_hosts
discovery.seed_hosts: ["192.168.1.195", "192.168.1.196", "192.168.1.197"]
```

#### b. 定义集群初始化时的主节点列表

注意:在上述过程中,我定义了2个主节点和1个数据节点,因此只有2个主节点的`node.name`项的值,会出现在`cluster.initial_master_nodes`中,该配置项用于定义在集群初始化时,可以参与主节点选举的节点列表

这一步3个节点将`cluster.initial_master_nodes`配置项的值设置为相同的即可:

```
root@es-node-1:~# cat /etc/elasticsearch/elasticsearch.yml|grep cluster.initial_master_nodes
cluster.initial_master_nodes: ["es-node-1", "es-node-2"]
# cluster.initial_master_nodes: ["es-node-1"]
```

这里注意,还有一处`cluster.initial_master_nodes`,是用于配置SSL和安全设置的,暂时将其注释即可

```
root@es-node-2:~# cat /etc/elasticsearch/elasticsearch.yml|grep cluster.initial_master_nodes
cluster.initial_master_nodes: ["es-node-1", "es-node-2"]
# cluster.initial_master_nodes: ["es-node-2"]
```

```
root@es-node-3:~# cat /etc/elasticsearch/elasticsearch.yml|grep cluster.initial_master_nodes
cluster.initial_master_nodes: ["es-node-1", "es-node-2"]
# cluster.initial_master_nodes: ["es-node-3"]
```

### 5.5 其他工作

将所有`xpack`相关的配置项全部注释

然后在配置文件末尾添加:

```
xpack.security.enabled: false
xpack.security.http.ssl.enabled: false
# xpack.security.http.ssl.keystore.path: certs/http.p12
xpack.security.transport.ssl.enabled: false
# xpack.security.transport.ssl.verification_mode: certificate
# xpack.security.transport.ssl.keystore.path: certs/transport.p12
# xpack.security.transport.ssl.truststore.path: certs/transport.p12
# xpack.security.transport.ssl.keystore.secure_password: your_keystore_password
```

注: 这里取消注释并将bool类型的选项置为true,即可开启SSL

但是这里我无法解密http.p12,因为不知道密码.所以就先关了.后续可能需要使用自签发的证书来解决这个问题

## PART6. 启动ES服务

3个节点均启动服务并设置开机自启动即可:

```
root@es-node-3:~# systemctl start elasticsearch.service
root@es-node-3:~# systemctl enable elasticsearch.service 
Created symlink /etc/systemd/system/multi-user.target.wants/elasticsearch.service → /lib/systemd/system/elasticsearch.service.
```

## PART7. 查看集群信息

```
root@es-node-1:~# curl -X GET "localhost:9200/_cluster/health?pretty"
{
  "cluster_name" : "my-es-cluster",
  "status" : "green",
  "timed_out" : false,
  "number_of_nodes" : 3,
  "number_of_data_nodes" : 3,
  "active_primary_shards" : 0,
  "active_shards" : 0,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 0,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 100.0
}
```