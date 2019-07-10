## Welcome to visit us
![Image text](https://github.com/rockwang465/rockwang.com/blob/master/picture/lyf.jpg)





# documents
### 1. 二进制部署Kubernetes v1.13.4 HA可选
```
https://zhangguanzhang.github.io/2019/03/03/kubernetes-1-13-4/
```

### 2. ansible一键部署高可用Kubernetes
```
https://github.com/zhangguanzhang/Kubernetes-ansible
```

### 3. k8s高可用涉及到ip填写的相关配置和一些坑
+ 这里讲到了k8s负载均衡机制
```
https://zhangguanzhang.github.io/2019/03/11/k8s-ha/
```
```
apiVersion: kubeadm.k8s.io/v1beta1
kind: ClusterConfiguration
kubernetesVersion: stable
apiServer:
  certSANs:
  - "LOAD_BALANCER_DNS"
controlPlaneEndpoint: "LOAD_BALANCER_DNS:LOAD_BALANCER_PORT"
```
+ `controlPlaneEndpoint`:应匹配负载均衡器的地址或DNS和端口
+ 而这个`controlPlaneEndpoint`实际上最终会取ip(注意不带端口)写到kube-apiserver的选项`--advertise-address`作为值。
+ 默认情况下`--advertise-address`不配置将会和`--bind-address`一样。它的作用就是宣告,在etcd启动后kube-apiserver初次起来后会创建一个svc名叫kubernetes,而这个svc的endpoints就是选项`--advertise-address`的ip,port则是apiserver的`--secure-port`。假设用户配置的`--secure-port`为6443,所以一般云上SLB的话那这个宣告可以填写LB的ip,然后默认的kubernetes的endpoints是`<SLB_IP>:6443`。
+ Rock : 具体可以看公司的配置 `roles/kubernetes/templates/kubeadm-init.yaml`
+ 模板文件
```
# cat roles/kubernetes/templates/kubeadm-init.yaml
apiVersion: kubeadm.k8s.io/v1beta1
kind: InitConfiguration
localAPIEndpoint:
  advertiseAddress: "{{ ansible_host | default(ansible_default_ipv4.address) }}"
  bindPort: 6443
nodeRegistration:
{% if kube_override_hostname|default('') %}
  name: {{ kube_override_hostname }}
{% endif %}
{% if kube_master_noschedule and inventory_hostname in groups['kube-master'] and inventory_hostname not in groups['kube-node'] %}
  taints:
  - effect: NoSchedule
    key: node-role.kubernetes.io/master
{% endif %}
  kubeletExtraArgs:
    cgroup-driver: "{{ docker_cgroup_driver }}"
---
apiVersion: kubeadm.k8s.io/v1beta1
kind: ClusterConfiguration
# etcd:
#   external:
#       endpoints:
# {% for endpoint in etcd_access_addresses.split(',') %}
#       - {{ endpoint }}
# {% endfor %}
#       caFile: {{ etcd_cert_dir }}/ca.pem
#       certFile: {{ etcd_cert_dir }}/node-{{ inventory_hostname }}.pem
#       keyFile: {{ etcd_cert_dir }}/node-{{ inventory_hostname }}-key.pem
networking:
  dnsDomain: {{ dns_domain }}
  serviceSubnet: {{ kube_service_subnet }}
  podSubnet: {{ kube_pod_subnet }}
kubernetesVersion: "{{ kube_version }}"
{% if kube_api_loadbalancer is defined %}
controlPlaneEndpoint: {{ kube_api_loadbalancer.domain }}:6443
{% else %}
controlPlaneEndpoint: {{ ansible_host | default(ansible_default_ipv4.address) }}:6443
{% endif %}
clusterName: {{ cluster_name }}
certificatesDir: {{ kube_cert_dir }}
imageRepository: {{ kube_image_repo }}
useHyperKubeImage: false
apiServer:
  certSANs:
{% if kube_api_loadbalancer is defined %}
  - {{ kube_api_loadbalancer.domain }}
  - {{ kube_api_loadbalancer.ip }}
{% endif %}
{% for item in groups['kube-master'] %}
  - {{ item }}
  - {{ hostvars[item]['access_ip'] | default(hostvars[item]['ip'] | default(hostvars[item]['ansible_host'])) }}
{% endfor %}
---
apiVersion: kubeproxy.config.k8s.io/v1alpha1
kind: KubeProxyConfiguration
ipvs:
  excludeCIDRs:
{% for item in kube_proxy_exclude_cidrs %}
  - {{ item }}
{% endfor %}
  minSyncPeriod: {{ kube_proxy_min_sync_period }}
  scheduler: {{ kube_proxy_scheduler }}
  syncPeriod: {{ kube_proxy_sync_period }}
mode: {{ kube_proxy_mode }}
```
+ 模板生成后的文件
``` 
# cat /etc/kubernetes/kubeadm-init.yaml

apiVersion: kubeadm.k8s.io/v1beta1
kind: InitConfiguration
localAPIEndpoint:
  advertiseAddress: "10.5.6.63"
  bindPort: 6443
nodeRegistration:
  kubeletExtraArgs:
    cgroup-driver: "systemd"
---
apiVersion: kubeadm.k8s.io/v1beta1
kind: ClusterConfiguration
# etcd:
#   external:
#       endpoints:
# #       -
# #       caFile: /etc/kubernetes/pki/etcd/ca.pem
#       certFile: /etc/kubernetes/pki/etcd/node-nebula-test-63.pem
#       keyFile: /etc/kubernetes/pki/etcd/node-nebula-test-63-key.pem
networking:
  dnsDomain: cluster.local
  serviceSubnet: 10.96.0.0/16
  podSubnet: 10.244.0.0/16
kubernetesVersion: "v1.13.2"
controlPlaneEndpoint: kube-api-lb.cluster.local:6443
clusterName: nebula-cluster
certificatesDir: /etc/kubernetes/pki
imageRepository: 10.5.6.10/kubernetes
useHyperKubeImage: false
apiServer:
  certSANs:
  - kube-api-lb.cluster.local
  - 10.5.6.201
  - nebula-test-63
  - 10.5.6.63
  - nebula-test-64
  - 10.5.6.64
  - nebula-test-65
  - 10.5.6.65
---
apiVersion: kubeproxy.config.k8s.io/v1alpha1
kind: KubeProxyConfiguration
ipvs:
  excludeCIDRs:
  - 10.5.6.100/24
  minSyncPeriod: 0s
  scheduler: rr
  syncPeriod: 30s
mode: ipvs
```


