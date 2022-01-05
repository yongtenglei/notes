# [consul](https://github.com/hashicorp/consul)

## 搭建 docker consul 服务

```docker
version: "3"

services:

  consul1:
    image: consul
    container_name: node1
    command: agent -server -bootstrap-expect=3 -node=node1 -bind=0.0.0.0 -client=0.0.0.0 -datacenter=dc1

  consul2:
    image: consul
    container_name: node2
    command: agent -server -retry-join=node1  -node=node2 -bind=0.0.0.0 -client=0.0.0.0 -datacenter=dc1
    depends_on:
      - consul1

  consul3:
    image: consul
    container_name: node3
    command: agent -server -retry-join=node1  -node=node3 -bind=0.0.0.0 -client=0.0.0.0 -datacenter=dc1
    depends_on:
      - consul1

  consul4:
    image: consul
    container_name: node4
    command: agent  -retry-join=node1  -node=node4 -bind=0.0.0.0 -client=0.0.0.0 -datacenter=dc1 -ui
    depends_on:
      - consul2
      - consul3
    ports:
      - "8500:8500"

```

启动服务:
`docker-compose up`

查看服务状态(web ui):
`localhost:8500`

## 注册与反注册服务

### 注册服务 ()

<div align=center><img src="https://tvax2.sinaimg.cn/large/006cK6rNly1gxtfawg0o0j30h00j0jy0.jpg">

</div>

<div align=center><img src="https://tva4.sinaimg.cn/large/006cK6rNly1gxtf4bwq35j31eu0hm0xe.jpg">
</div>

<div align=center><img src="https://tvax4.sinaimg.cn/large/006cK6rNly1gxtf6fbyc1j30h806egmu.jpg">
</div>

但不提供健康检查

>

### 反注册服务

<div align=center><img src="https://tva4.sinaimg.cn/large/006cK6rNly1gxtfb8wy49j30hc0gcdl0.jpg">

</div>

<div align=center><img src="https://tva3.sinaimg.cn/large/006cK6rNly1gxtf982a4qj31es0futbu.jpg">

</div>
