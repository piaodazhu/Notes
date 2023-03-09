## k8s特性：
- 高可用，自动灾难恢复
- 灰度更新
- 可回滚历史版本
- 伸缩扩展
- 负载均衡
- 完善的生态

## 概念
pod: 一个pod包含一个容器组，容器组内共享同一个数据卷，每个pod有自己的虚拟IP

node: 一个node包含多个pod，主节点会考虑调度pod到哪个worker node运行

control plain: 由api server，scheduler，etcd，controller manager，cloud controller manager组成

## 配置
工作负载分类: Deployment StatefulSet DaemonSet Job&CronJob

service: 生命周期不跟pod绑定，提供负载均衡能力，对集群提供外部访问端口，服务内部的pods可以通过名字直接访问服务。type中ClusterIP是仅内部访问，NodePort提供带有负载均衡的外部访问，LoadBalance需要负载均衡器支持生成新IP不常用（见Ingress），Headless不会分配IP。

statefulset: 无状态应用可以随意扩充副本，但是有状态的如redis，mysql，不能随意扩充副本。statefulset会固定每个pod的名字。

数据持久化需要添加数据卷

configmap可以把配置的变量作为环境变量传入容器

Helm包管理器可以为集群安装第三方软件，可以从artifacthub找到

命名空间kubens便于管理

Ingress，需要负载均衡插件(如METALLB)

## 部署之minikube
用于在单机学习k8s
**大坑**: 不要设置http_proxy和https_proxy, 运行`minikube start --image-mirror-country='cn' --image-repository='registry.cn-hangzhou.aliyuncs.com/google_containers'`

## 裸机部署
全是坑。直接看`https://www.bilibili.com/video/BV1Tg411P7EB`

## 操作
总体查看

```sh
kubectl get nodes
kubectl get pods (-o wide可看到pod在节点上如何分配)
kubectl get deployment
```

pod查看
```sh
kubectl get all
kubectl describe pod podname
kubectl describe service podname
kubectl logs podname
kubectl exec -it podname -- bash
```

部署
```sh
kubectl apply -f xxx.yaml
kubectl scale depolyment deployname --replicas=N
kubectl port-forward pod-name localport:podport
kubectl delete deployment deployname
kubectl rollout restart deployment deployname
kubectl pause deployment deployname
kubectl resume deployment deployname
```

版本管理
```sh
kubectl rollout history deployment deployname
kubectl rollout undo deployment deployname (--to-revision=x)
```

服务管理
```sh
kubectl apply -f xxxservice.yaml
```

## redis集群部署实践
### 实践环境
物理机一台，vmware运行三台Ubuntu Server 22.04.1 LTS虚拟机，一台作为master，其余作为node1和node2，安装docker环境。

### 基本步骤
1. 准备k8s环境，master节点上安装cni，观察到所有节点READY，master节点上安装helm包管理器。
2. 创建足够的pv。
3. 参考[AritifactHub/redis](https://artifacthub.io/packages/helm/bitnami/redis)部署redis。
4. 集群内测试 & 集群外测试。

### 关键点之pv
pv就是Persistent Volume，持久卷信息。pvc是Persistent Volume Claim持久卷声明。helm部署redis的时候，会在集群中创建pvc。可以把它理解为服务中的每个pods对存储卷的具体需求。而pv则是能满足这些需求的存储卷。k8s中的数据持久化分3层抽象即pvc-pv-sc。sc即Storage Class，目前还没有用到。

pvc示例：
```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: TestPVC
spec:
  accessModes: ["ReadWriteOnce"]
  storageClassName: "local-storage"
  resources:
    requests:
      storage: 2Gi
```
表示TestPVC至少需要2GB的持久化存储空间，那么pv容量大于2GB就能满足需求。注意有几个pod就需要几个pv，比如redis集群指定1master3replicas，就需要创建4个pv。

pv示例：
```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: TestPV
spec:
  capacity:
    storage: 2Gi
  accessModes:
    - ReadWriteOnce       # 卷可以被一个节点以读写方式挂载
  persistentVolumeReclaimPolicy: Retain
  ###storageClassName: local-storage
  local:
    path: /tmp/pv
  nodeAffinity:
    required:
      # 通过 hostname 限定在某个节点创建存储卷
      nodeSelectorTerms:
        - matchExpressions:
            - key: kubernetes.io/hostname
              operator: In
              values:
                - node2
```
表示TestPV具有2GB的存储空间，还需要注意的是，pv必须设置节点亲和性，比如这里的node2，表示这块pv是放在node2上的，另外spec.local给的这个path，必须是存在于node2本地文件目录中的目录。persistentVolumeReclaimPolicy可以是Retain或者Delete，这里Retain表示删除了pvc之后，这个pv会保留，要回收利用他，需要手动删掉pv的Claim信息，详见后文。

下面是pv有关的踩坑经历。

首先通过helm安装的redis，目前为止正常。然后运行`kubectl get pods`看，发现redis新建的pods直接就是PENDING状态。后来	`kubectl describe pods redis-xxx-master`之类的查看，报错信息大致是找不到足够的PersistentVolume。用`kubectl get pv`一看，什么都没有；用`kubectl get pvs`一看，有两个redis相关的pvs状态是PENDING，这时就要建立一些pv。去网上找了一个pv文件，这里面的节点亲和性的values要改成自己集群内部的某个节点名字，path也要改成自己的，确保目录存在，然后`kubectl apply -f pv1.yaml`。这时不报错，用`kubectl get pv`发现pv创建成功，过了一会发现pv变成了BOUND状态，而其中一个pvs状态恢复了正常。说明当前一个pv可以满足一个pvs。按照相同的办法，取不同的名字、节点、路径，建立多个pv，等pv被pvs领走之后，服务恢复正常。经过测试redis部署成功。
