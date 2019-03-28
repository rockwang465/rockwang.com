# Kubernetes install by kubeadm
## 1. Install introduction
* The first step
``` 
/bin/sh init-ansible.sh
```
* The second step
``` 
ansible-playbook init-env.yaml
```
* The third step
``` 
ansible-playbook install-k8s.yaml
```

