#!/usr/bin/env python
# encoding: utf-8

# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# creater: Rock Wang                                             +
# creation time: 2019-11-19                                      +
# description: Memory and CPU resources for optimizing services  +
# ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

import os
import sys
import time
import argparse
import socket
from get_k8s_info import *
from render_templates import *

# 所有优化的服务名，如果您不需要对应服务的优化，请将对应列表值
optimization_server_name = {
    'component': ['cassandra', 'kafka'],
    'logging': ['elasticsearch'],
    'monitoring': ['prometheus-operator'],
    'nebula': ['engine-timespace-feature-db']
}

packages_path = '/opt/optimization'


def parse_args():
    parser = argparse.ArgumentParser(description='Memory and CPU resources for optimizing services')
    return parser.parse_args()


# 0. Welcome
class base_info:
    def Welcome(self):
        print("+" * 45)
        print("+ Welcome use source optimization scripts ! +")
        print("+" * 45)

    def get_local_ip(self):
        myname = socket.getfqdn(socket.gethostname())
        self.local_ip = socket.gethostbyname(myname)
        print("\n")
        print("Warning: Please confirm that the IP of master1 is: [%s] ,execute script in 5 seconds !" % self.local_ip)
        time.sleep(1)
        print("Start execution ...")
        time.sleep(1)


# 5. 获取之前保存字典中的override文件路径，及各服务的values.yaml的路径
class get_modify_file:
    # 修改/opt/optimization目录下override文件中request的memory及cpu的大小
    def modify_override_yaml(self, override_file):
        for key_name in override_file:  # key_name: component
            # print(override_file)
            # print(key_name)
            for name in override_file.get(
                    key_name):  # override_file.get(key_name): {'elasticsearch': '/opt/optimization/logging/elsaticsearch.values.yml'}
                server_name = name
                override_yaml_path = override_file.get(key_name).get(name)
                modify_values.modify_override_args(server_name, override_yaml_path)  # 传入服务名，文件路径

    # 修改values.yaml中的cpu memory
    def modify_values_yaml(self, charts_version):
        # 解压所有fetch下来的charts包
        res = os.system("cd %s && for i in `ls *.tgz` ; do tar xf $i ; done" % packages_path)
        if res != 0:
            print("Error : failure to [cd %s && ls *.tgz | xargs tar xf]" % packages_path)
            sys.exit(1)
        for ns in charts_version:
            for server_name in charts_version.get(ns):
                values_yaml_path = packages_path + "/" + server_name + "/values.yaml"
                modify_values.modify_values_args(server_name, values_yaml_path)


# 6. 资源优化后，开始更新服务
class update_optimization_service:
    def upgrade_service(self, charts_version):
        for ns in charts_version:
            for server_name in charts_version.get(ns):
                tmp_override_new_file = "/tmp/" + server_name + ".values.yaml"
                os.chdir("%s/%s" % (packages_path, server_name))
                os.system("helm upgrade -i %s-%s --namespace=%s -f %s . >/dev/null 2>&1" % (
                    server_name, ns, ns, tmp_override_new_file))
                print("Info : updated [%s] service, please check [kubectl get pods -n %s | grep %s]" % (
                    server_name, ns, server_name))
                time.sleep(2)


if __name__ == '__main__':
    # 0. 欢迎语
    args = parse_args()
    base = base_info()
    base.Welcome()
    local_ip = base.get_local_ip()

    # 1.执行get_k8s_info.py脚本，获取nodes信息，用于模板渲染
    get_images_info = get_kube_info()
    get_images_info.get_kube_config(base.local_ip)
    get_images_info.get_nodes_info()

    # 2. 执行get_k8s_info.py脚本，获取需要优化服务的版本信息、拉取对应版本的包
    get_packages = get_charts_packages(optimization_server_name, packages_path)
    get_packages.get_helm_version()  # A.获取需要优化服务的版本信息
    get_packages.fetch_helm_packages()  # B.拉取对应版本的包
    # get_packages.get_override_name()  # C.找到/tmp/ 目录下的override文件 -- 此函数弃用

    # 3.执行render_templates.py脚本，使用jinja2将templates下的模板文件渲染后放入/opt/optimization/override_yaml/<namespace>下
    render = template_render(get_images_info.all_nodes, packages_path)
    render.init_args()
    render.get_template_file(get_packages.charts_version)

    # 4.执行render_templates.py脚本，定义修改需要优化服务的包中values.yaml及override文件中的request.memory request.cpu的大小，及configmap的jvm大小的函数方法
    modify_values = modify_request_values(render.override_file)

    # 5. 获取之前保存字典中的override文件路径，及各服务的values.yaml的路径
    get_file = get_modify_file()
    get_file.modify_override_yaml(render.override_file)  # 修改/opt/optimization/<namespace>目录下的override文件
    get_file.modify_values_yaml(get_packages.charts_version)  # 修改服务中values.yaml文件

    # 6. 资源优化后，开始更新服务
    update_service = update_optimization_service()
    update_service.upgrade_service(get_packages.charts_version)
