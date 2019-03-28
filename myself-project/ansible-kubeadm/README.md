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













########################
modify network name to eth0
1.
```
cd /etc/sysconfig/network-scripts/
mv ifcfg-ens33 ifcfg-eth0
vi ifcfg-eth
name=eth0
```

2. 
```
vi /etc/default/grub
biosdevname=0 net.ifnames=0
grub2-mkconfig -o /boot/grub2/grub.cfg
echo 'nameserver 8.8.8.8' >>/etc/resolv.conf
reboot  #reset machine
```
