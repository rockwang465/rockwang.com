#!/usr/bin/env python
# -*- coding: utf-8 -*-

import subprocess
import os
from kubernetes import client, config
from kubernetes.client import configuration

kube_config_file = '/root/.kube/config'
kube_config_new_file = 'kubeconfig'
local_domain = '127.0.0.1'


# 1. 通过k8s api获取nodes信息
class get_kube_info:
    # 获取/root/.kube/config文件
    def get_kube_config(self, local_ip):
        cmdstr = "cat %s" % kube_config_file
        res = subprocess.Popen(cmdstr, shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
        if res.stderr != 0:
            self.kube_config_body = res.stdout.read()

        # 当config中含有127.0.0.1时，替换为环境的ip。
        if local_domain in self.kube_config_body:
            kube_file = self.kube_config_body.replace(local_domain, local_ip)
        else:
            kube_file = self.kube_config_body

        with open(kube_config_new_file, 'w') as fp:
            fp.write(kube_file)

    # 获取k8s的各节点信息
    def get_nodes_info(self):
        config.load_kube_config(kube_config_new_file)
        configuration.assert_hostname = False
        configuration.verify_ssl = False
        v1 = client.CoreV1Api()

        res = v1.list_node()  # 获取nodes信息
        self.all_nodes = []  # 存放所有node的信息
        for item in res.items:
            node = {}
            # item.status.addresses为所有集群节点的主机名和ip地址的列表，如后面为其中一个列表: [{'address': '10.151.5.44', 'type': 'InternalIP'}, {'address': 'nebula-test-44', 'type': 'Hostname'}]
            for node_info in item.status.addresses:
                if node_info.type == "Hostname":
                    node["hostname"] = node_info.address
                if node_info.type == "InternalIP":
                    node["ip"] = node_info.address
            self.all_nodes.append(node)


# 2. helm list 获取版本号，并写入optimization_server.txt中
class get_charts_packages:
    def __init__(self, charts_version, optimization_server_name, packages_path):
        self.charts_version = charts_version
        self.optimization_server_name = optimization_server_name
        self.packages_path = packages_path

    # 获取helm list 中的版本信息
    def get_helm_version(self):
        for key_ns in self.optimization_server_name:
            for server_name in self.optimization_server_name.get(key_ns):
                res = os.popen("helm list | grep %s | grep -v elasticsearch-curator | awk '{print $9}'" % server_name)
                res_data = res.read().strip()
                if res_data:
                    self.charts_version[key_ns].append(res_data)
                else:
                    print('Error : Not found server name : [%s]' % server_name)
        # print("Info : Here is the chart information :")
        # print(self.charts_version)
        print("\n")

    # 通过上面获取的版本信息，这里进行fetch下载
    def fetch_helm_packages(self):
        for key_ns in self.charts_version:
            for charts_info in self.charts_version.get(key_ns):
                if not os.path.exists(self.packages_path):
                    print("Info : create directory %s" % self.packages_path)
                    os.mkdir(self.packages_path)
                print(
                        "Info : [cd %s && helm fetch http://10.151.3.75:8080/charts/%s.tgz]" % (
                    self.packages_path, charts_info))
                res = os.system(
                    "cd %s && helm fetch http://10.151.3.75:8080/charts/%s.tgz" % (self.packages_path, charts_info))
                if res != 0:
                    print("Error : failure to [helm fetch http://10.151.3.75:8080/charts/%s.tgz]")
                    sys.exit(1)

    # # 获取 /tmp/下各服务的override.values.yaml文件名 -- 老版本用/tmp/下君宇生成的override文件，但容易被系统清理掉，所以弃用了。
    # def get_override_name(self):
    #     for key_name in optimization_server_name:
    #         # print(optimization_server_name.get(key_name))
    #         for server_name in optimization_server_name.get(key_name):
    #             # print(server_name)
    #             res = os.popen('ls /tmp/%s-%s* | head -1' % (server_name, key_name))
    #             server_file = res.read().strip()
    #             if server_file:
    #                 # 添加对应服务在/tmp/目录下的overrid文件名
    #                 self.override_file[key_name][
    #                     server_name] = server_file  # override_file['component']['kafka']='/tmp/kafka-component.1573202714'
    #             else:
    #                 print("Error : Not found [/tmp/%s-%s....] file" % (server_name, key_name))
    #                 sys.exit(1)
    #     # print(self.override_file)
