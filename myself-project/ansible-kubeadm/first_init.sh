#!/bin/sh

echo -e "\n1. Install ansible "
yum install ansible -y >/dev/null 2>&1
ansible-playbook -h >/dev/null 2>&1
if [ $? -eq 0 ];then
    echo -e "ansible install successfully"
else
    echo -e "Error : install failuer" 
fi
