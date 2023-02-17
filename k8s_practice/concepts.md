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
工作负载分类: Deployment StatefulSet DaemonSet Job&CronJob
service: 生命周期不跟pod绑定，提供负载均衡能力，对集群提供外部访问端口


## 部署之minikube
用于在单机学习k8s
**大坑**: 不要设置http_proxy和https_proxy, 运行`minikube start --image-mirror-country='cn' --image-repository='registry.cn-hangzhou.aliyuncs.com/google_containers'`

## 裸机部署
全是坑。直接看`https://www.bilibili.com/video/BV1Tg411P7EB`

## 操作
总体查看
kubectl get nodes
kubectl get pods (-o wide可看到pod在节点上如何分配)
kubectl get deployment

pod查看
kubectl get all
kubectl describe pod podname
kubectl logs podname
kubectl exec -it podname -- bash

部署
kubectl apply -f xxx.yaml
kubectl scale depolyment deployname --replicas=N
kubectl port-forward pod-name localport:podport
kubectl delete deployment deployname
kubectl rollout restart deployment deployname
kubectl pause deployment deployname
kubectl resume deployment deployname

版本管理
kubectl rollout history deployment deployname
kubectl rollout undo deployment deployname (--to-revision=x)