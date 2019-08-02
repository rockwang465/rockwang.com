#!/usr/bin/env python
# -*- coding: utf-8 -*-
import paramiko
import re
import yaml
from kubernetes import client, config
import time

standard_env_ip = '10.5.6.66'
port = 22
username = 'root'
passwd = 'Nebula123$%^'
repository_domain = 'www.sensetime.com'  # docker images domain name(以后docker镜像可能会定义一个域名开头的，而非IP开头的)
kube_config_path = '/root/.kube/config'
kube_config_file = './kubeconfig'
new_version = './new_versions.yaml'
charts_groups = ['addons_charts', 'component_charts', 'console_charts', 'devops_charts', 'guard_charts',
                 'nebula_charts']
all_namespace = ['component', 'nebula', 'default', 'logging', 'monitoring']  # 未加: galaxias helm kube-public kube-system

ssh = paramiko.SSHClient()


# ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
# ssh.connect(standard_env_ip, port, username, passwd)

class get_charts_version:
    def __init__(self):
        self.charts_groups = charts_groups

    # get old charts information from version.yaml
    def get_old_charts(self):
        data = open('./versions.yaml', 'r')
        self.old_charts = yaml.full_load(data.read())

    # get helm charts version in standard environment
    def get_helm_charts(self):
        # ssh = paramiko.SSHClient()
        try:
            ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
            ssh.connect(standard_env_ip, port, username, passwd)
            stdin, stdout, stderr = ssh.exec_command(
                "helm list --col-width 200 | sed 1d | awk '{print $9}' | sort -rn | uniq")
            # print(stdout.read().decode())
            helm_info = stdout.read().decode()
            ssh.close()
            self.helm_list = []
            for i in helm_info.split("\n"):
                if i == "":
                    pass
                else:
                    self.helm_list.append(i.encode('utf-8'))
            # print(self.helm_list)
        except Exception as ex:
            print("Error : Failure to get helm charts version : %s" % ex)

    # convert standard environment helm charts version to dict
    def convert_helm_charts(self):
        # print(self.helm_list)
        machine_charts = {}
        for i in self.helm_list:
            chart_name = re.findall(r"(.+?)-[0-9]|v[0-9]", i)

            # when chart_name == []:
            if chart_name == []:
                continue
            # when '-v1' in chart_name[0]:
            elif chart_name[0]:
                if '-v1' in chart_name[0]:
                    chart_name = re.findall(r"(.+?)-v[0-9]", i)

            # when chart_name == ['']:
            if chart_name == ['']:
                chart_name = re.findall(r"(.+?)-v[0-9]|[0-9]", i)
                chart_name = chart_name[0]
                chart_version = re.split(r'%s-' % chart_name, i)[-1]
                machine_charts[chart_name] = chart_version
            # else is correct value:
            else:
                chart_name = chart_name[0]
                chart_version = re.split(r'%s-' % chart_name, i)[-1]
                machine_charts[chart_name] = chart_version
        self.new_charts = machine_charts
        # print(self.new_charts)  # 备用勿删

    # compare old and new charts version , and save new charts version ,and output to new yaml file
    # (,and print infomation on a display)
    def compare_charts(self):
        # old_charts example : {'devops_charts': [{'version': '6.5.4-master-a9e66f5', 'namespace': 'logging', 'name': 'elasticsearch'}, ...}
        # new_charts example : {'senseguard-oauth2': '1.2.0-v1.2.0-001-cbe1b9', 'senseguard-bulk-tool': '1.2.0-v1.2.0-001-46db20',...}
        new_charts_version_dict = self.old_charts

        # loop the charts groups,for example: 'addons_charts', 'component_charts', 'console_charts', 'devops_charts' ...
        for group in self.charts_groups:
            for index, charts in enumerate(self.old_charts[group]):
                if self.new_charts.get(charts["name"]):
                    # print(charts["name"])
                    if new_charts_version_dict[group][index]["version"] == self.new_charts.get(charts["name"]):
                        # print("[Info] : [%s] chart version is right" % charts["name"])  # 用于提示的，此行保留备用
                        pass
                    else:
                        print(
                            "[Warning] : [%s] chart version not right, the old version is : %s , the new version is : %s" % (
                                charts["name"], new_charts_version_dict[group][index]["version"],
                                self.new_charts.get(charts["name"])))
                        new_charts_version_dict[group][index]["version"] = self.new_charts.get(charts["name"])
                else:
                    print("[Error] : The current machine was not found this chart name : %s" % charts["name"])

        fp = open(new_version, 'w')
        yaml.dump(self.old_charts, fp)


# get docker image version in standard environment(通过k8s api获取docker image版本号)
class get_images_version:
    # def __init__(self):
    #     print("class : get docker images version")

    def get_kube_config(self):
        try:
            ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
            ssh.connect(standard_env_ip, port, username, passwd)
            stdin, stdout, stderr = ssh.exec_command("cat %s" % kube_config_path)
            # print(stdout.read().decode())
            self.kube_config_body = stdout.read().decode()
            ssh.close()
            with open(kube_config_file, 'w') as fp:
                fp.write(self.kube_config_body)
            # return stdout.read().decode()
        except Exception as ex:
            print("Error : Failure to get kube config : %s" % ex)

    def get_kube_images(self):
        config.load_kube_config(kube_config_file)
        v1 = client.CoreV1Api()

        images_dict = {}
        # images_dict = {'sensenebula-guard-std/senseguard-td-result-consume': '1.2.0-v1.2.0-b8d7c0222222222222222'}
        for ns in all_namespace:
            print(ns + "-------------------------------->")
            reps = v1.list_namespaced_pod(ns)
            for item in reps.items:
                for cont in item.spec.containers:
                    tmp_img1 = cont.image  # tmp_img1 = '10.5.6.10/sensenebula-guard-std/senseguard-timezone-management:1.2.0-v1.2.0-eedcba'
                    print('tmp_img1-source value:', tmp_img1)
                    tmp_tag1 = tmp_img1.split("/")[
                           1:]  # tmp_tag1 = ['sensenebula-guard-std', 'senseguard-timezone-management:1.2.0-v1.2.0-eedcba']
                    tmp_host = tmp_img1.split("/")[:1]
                    # print('tmp_tag1-source value 222:', tmp_tag1)
                    tmp_tag2 = tmp_tag1[-1].split(":")  # tmp_tag2 = ['senseguard-timezone-management', '1.2.0-v1.2.0-eedcba']
                    print('tmp_tag2-source value 333333:', tmp_tag2)
                    name = tmp_tag1[0] + "/" + tmp_tag2[0]  # name = sensenebula-guard-std/senseguard-timezone-management
                    tag = tmp_tag2[-1]  # tag = 1.2.0-v1.2.0-eedcba

                    # 如果有数字开头的，(或者后期定的域名)，则为正常
                    if repository_domain == tmp_host[0] or tmp_host[0].replace(".", "").replace(":", "").isdigit():  # 判断tmp_tag1[0]为域名，或者tmp_tag1[0]去掉 .点 和 :冒号 后为数字(即为IP)
                        if images_dict.get(name) and images_dict.get(name) == tag:  # 已存在此key和value
                            # print("Info : Already exist this key[%s] and value[%s]" % (name, tag))
                            pass
                        elif images_dict.get(name) and images_dict.get(name) != tag:  # 有此key，但value值不同，报错
                            print(
                                "Error : Already exist this key[%s] , But the tag are different; old tag:[%s] , new tag:[%s]" % (
                                    name, images_dict.get(name), tag))
                        else:
                            # 如果没有tag，则为 tag=latest
                            # if not images_dict[name]
                            #     images_dict[name] = tag
                            print('maybe has error tag:', tag)
                    else:
                        print('tmp_tag1[0] not is domain name, and not is IP:', tmp_tag1[0])

        print(images_dict)


if __name__ == '__main__':
    charts = get_charts_version()
    charts_info = charts.get_helm_charts()
    charts.convert_helm_charts()
    charts.get_old_charts()
    charts.compare_charts()

    images = get_images_version()
    images.get_kube_config()
    images.get_kube_images()
