#!/usr/bin/env python
# -*- coding: utf-8 -*-

# from jinja2 import Environment, PackageLoader
# from jinja2 import Environment, FileSystemLoader
import jinja2
import os
from class_services_deploy_client import *


def get_k8s_node():
    result1 = os.popen("cat /etc/hosts | grep '.cluster.local' | grep -v kube-api | awk '{print $2}'")
    result2 = result1.readlines()
    k8s_node = []
    for i in result2:
        node = i.replace("\n", "")
        k8s_node.append(node)
    return k8s_node


def render(path, master_ip, slave_ip, deploy_mode, k8s_node):
    files1 = os.popen('ls templates/%s' % path)
    files2 = files1.readlines()
    for file in files2:
        # print(files2)
        file_name = file.replace("\n", "")
        TemplateLoader = jinja2.FileSystemLoader(os.path.abspath('templates/%s' % path))
        TemplateEnv = jinja2.Environment(loader=TemplateLoader)
        template = TemplateEnv.get_template(file_name)
        config = template.render(master_ip=master_ip, slave_ip=slave_ip, docker_registry=master_ip,
                                 deploy_mode=deploy_mode)
        f = open('./result/' + file_name, 'w')
        f.write(config)
        f.close()

        # with open('./result/%s' % file_name) as fp:
        #     fp.write(config)


k8s_node = get_k8s_node()

for i in temp_dir:
    render(i, master_ip, slave_ip, deploy_mode, k8s_node)
