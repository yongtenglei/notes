# k8s for beginner

====================================================

## 环境准备

### 必备

docker: `sudo pacman -S docker`

kubectl: `yay -S kubectl` for control kubernets

### 一定程度可选(推荐)

kind: `yay -S kind` for create cluster

### 可选

kubectx: `yay -S kubectx` change cluster faster

======================================================

## 使用 kind 创建一个集群

```yaml
# kind-cluster.yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
name: my-cluster
nodes:
  - role: control-plane
  - role: worker
  - role: worker
  - role: worker
```

以上文件指定 kind 为集群, version 以及 集群名称 name.

4 个节点, 一个为 control-plane, 三个 worker.

### 使用配置文件创建集群

`kind clusetr create --config=kind-cluster.yaml`

### 查看创立是否成功

`docker container ls`

### 查看目前所在集群

- 使用 kubectl

  `kubectl config get-contexts`
  <++>

- 使用 kubectx

  `kubectx` 查询当前所在交互集群
  <++>

  `kuberctx new_cluster` 更换交互集群

- 查看现有的 pods

  `kbubectl get pod` 第一次执行会显示没有资源

  `kbubectl get pod -A` 显示所有可用资源

======================================================

## 创建并发一个 docker 应用 such as `greeter`

### 创建一个 go 应用

```go
package main

import (
	"fmt"
	"strconv"
	"time"
)

func Greeting() {
	index := 0
	for {
		index++
		fmt.Println("hello k8s" + strconv.Itoa(index))
		time.Sleep(time.Second)
	}
}

func main() {
	Greeting()
}

```

### 创建 docker 镜像并推送到 docker hub

### Dockerfile

```Dockerfile
# Dockerfile
FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /greeter

CMD ["/greeter"]
```

### 制作镜像

`docker build . -t docker_hub_id/imagename[:version]`

eg.
`docker build . -t tinklerey/greeter[:v1.0.0]` 版本号可选 省略后默认为 latest

### 查看镜像

`docker image ls`

### 推送镜像

`docker push tinklerey/greeter`

======================================================

## 将 docker 应用部署到 kubernets cluster 上

### 配置 pod

```yaml
# greeter.yaml
apiVersion: v1
kind: Pod
metadata:
  name: greeter
  labels:
    plan: dev
  # namespace:
spec:
  containers:
    - name: greeter
      image: tinklerey/greeter
      resources:
        limits:
          memory: "200Mi"
          cpu: "500m"
        #requests:
        #memory: "100Mi"
      #command: ["stress"]
      #args: ["--vm", "1", "--vm-bytes", "150M", "--vm-hang", "1"]
```

`More config here` [More config](https://kubernetes.io/zh/docs/tasks/configure-pod-container/assign-memory-resource/)

### 应用 pod 配置文件进行配置

`kubectl apply -f greeter.yaml`

### 查看 pod 的状态

`kubectl get po` 查看 pods 的状态信息 (po / pod / pos are the same thing)

`kubectl descrebe po greeter` 查看更完整信息, 配置信息等

### 查看 pod 配置

`kubectl get pod greeter -oyaml` 查看 gretter yaml 配置

`kubectl get pod greeter -owide` 查看 gretter wide 形式

### 查看 pod 的 log 信息, follow pod

`kubectl logs greater [-n namespace]` 查看当前及以往的 log, -n flag 指定 namespace 的名称

`kubectl logs greater -f` -f flag 实时跟随 log 信息

======================================================

## 删除 pod

`kubectl delete -f greeter.yaml`

======================================================

[视频推荐](https://www.bilibili.com/video/BV1Xq4y1G7cF)
