# 附录4-ik分词器安装

以下操作3个节点都要做:

## PART1. 查找

`elasticsearch-plugin`是一个用于管理Elasticsearch插件的命令行工具.该工具主要用途有:

- 安装插件
- 移除插件
- 列出已安装的插件

```
root@es-node-1:~# find / -name elasticsearch-plugin
/usr/share/elasticsearch/bin/elasticsearch-plugin
```

## PART2. 查看当前ES版本

```
root@es-node-1:/usr/share/elasticsearch/bin# curl -X GET "http://localhost:9200" 
```

```JSON
{
  "name" : "es-node-1",
  "cluster_name" : "my-es-cluster",
  "cluster_uuid" : "M9cOUa5MRUOWKY6LS0aslg",
  "version" : {
    "number" : "8.15.2",
    "build_flavor" : "default",
    "build_type" : "deb",
    "build_hash" : "98adf7bf6bb69b66ab95b761c9e5aadb0bb059a3",
    "build_date" : "2024-09-19T10:06:03.564235954Z",
    "build_snapshot" : false,
    "lucene_version" : "9.11.1",
    "minimum_wire_compatibility_version" : "7.17.0",
    "minimum_index_compatibility_version" : "7.0.0"
  },
  "tagline" : "You Know, for Search"
}
```

这里查看版本,是因为ik分词器版本要求和ES的版本完全相同

## PART3. 安装ik分词器

```
root@es-node-1:~# cd /usr/share/elasticsearch/bin
root@es-node-1:/usr/share/elasticsearch/bin# ./elasticsearch-plugin install https://get.infini.cloud/elasticsearch/analysis-ik/8.15.2
....
Continue with installation? [y/N]y
-> Installed analysis-ik
-> Please restart Elasticsearch to activate any plugins installed
```

## PART4. 重启ES并查看已安装的插件

```
root@es-node-1:/usr/share/elasticsearch/bin# systemctl restart elasticsearch.service 
root@es-node-1:/usr/share/elasticsearch/bin# ./elasticsearch-plugin list
analysis-ik
```