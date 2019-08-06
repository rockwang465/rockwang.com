#!/usr/bin/env python
# -*- coding: utf-8 -*-
import paramiko
import re
from kubernetes import client, config
import sys
import json

# standard_env_ip = '10.5.6.66'
# port = 22
# username = 'root'
# passwd = 'Nebula123$%^'
# docker_repo_port = '5000'  # 在标准环境中的docker仓库端口
# repository_domain = 'www.sensetime.com'  # docker images domain name(以后docker镜像可能会定义一个域名开头的，而非IP开头的)
kube_config_path = '/root/.kube/config'
kube_config_file = './kubeconfig'
new_json_file = './versions.json'
all_namespace = ['component', 'nebula', 'default', 'logging', 'monitoring']  # 未加: galaxias helm kube-public kube-system
lack_images = [{'repository': 'elasticsearch/busybox', 'tag': 'latest'},
               {'repository': 'gitlabci/golang', 'tag': '1.9-cuda-gcc49-1'},
               {'repository': 'component/mc', 'tag': 'RELEASE.2019-02-13T19-48-27Z'},
               {'repository': 'external_storage/local-volume-provisioner', 'tag': 'v2.3.0'}]
ssh = paramiko.SSHClient()


# ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
# ssh.connect(standard_env_ip, port, username, passwd)


class get_charts_version:
    # get helm charts version in standard environment
    def get_helm_charts(self, standard_env_ip, port, username, passwd):
        # ssh = paramiko.SSHClient()
        cut_line = '{print $9" "$NF}'  # cut charts and namespace
        try:
            ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
            ssh.connect(standard_env_ip, port, username, passwd)
            stdin, stdout, stderr = ssh.exec_command(
                "helm list --col-width 200 | sed 1d | awk '%s' | sort -rn | uniq" % cut_line)
            # print(stdout.read().decode())
            get_helm_info = stdout.read().decode()
            # print(get_helm_info)
            ssh.close()
            self.helm_list = {"charts": []}
            for i in get_helm_info.split("\n"):
                if i == "":
                    pass
                else:
                    # self.helm_list.append(i.encode('utf-8'))
                    info = i.encode('utf-8').split()
                    # self.helm_info["charts"].append(1)
                    self.helm_list["charts"].append({
                        "info": info[0],
                        "namespace": info[-1]
                    })
        except Exception as ex:
            print("Error : Failure to get helm charts version : %s" % ex)
            sys.exit(1)  # 获取异常，后面的都不用执行了

    # convert standard environment helm charts version to dict
    def convert_helm_charts(self):
        # print(self.helm_list)
        machine_charts = {"charts": []}
        for i in self.helm_list["charts"]:
            chart_name = re.findall(r"(.+?)-[0-9]|v[0-9]", i["info"])

            # when chart_name == []:
            if chart_name == []:
                continue
            # when '-v1' in chart_name[0]:
            elif chart_name[0]:
                if '-v1' in chart_name[0]:
                    chart_name = re.findall(r"(.+?)-v[0-9]", i["info"])

            # when chart_name == ['']:
            if chart_name == ['']:
                chart_name = re.findall(r"(.+?)-v[0-9]|[0-9]", i["info"])
                chart_name = chart_name[0]
                chart_version = re.split(r'%s-' % chart_name, i["info"])[-1]
                # machine_charts[chart_name] = chart_version
                machine_charts["charts"].append({
                    "name": chart_name,
                    "version": chart_version,
                    "namespace": i["namespace"]
                })
            # else is correct value:
            else:
                chart_name = chart_name[0]
                chart_version = re.split(r'%s-' % chart_name, i["info"])[-1]
                # machine_charts[chart_name] = chart_version
                machine_charts["charts"].append({
                    "name": chart_name,
                    "version": chart_version,
                    "namespace": i["namespace"]
                })
        self.new_charts = machine_charts
        # print(self.new_charts)  # 备用勿删


# get docker image version in standard environment(通过k8s api获取docker image版本号)
class get_images_version:
    def get_kube_config(self, standard_env_ip, port, username, passwd):
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

    def get_kube_images(self, repository_domain):
        config.load_kube_config(kube_config_file)
        v1 = client.CoreV1Api()

        self.images_dict = {}
        # images_dict = {'sensenebula-guard-std/senseguard-td-result-consume': '1.2.0-v1.2.0-b8d7c0222222222222222'}
        for ns in all_namespace:
            reps = v1.list_namespaced_pod(ns)
            for item in reps.items:
                for cont in item.spec.containers:
                    tmp_img1 = cont.image  # tmp_img1 = '10.5.6.10/sensenebula-guard-std/senseguard-timezone-management:1.2.0-v1.2.0-eedcba'
                    # print('tmp_img1-source value:', tmp_img1)
                    tmp_host = tmp_img1.split("/")[:1][0]  # tmp_host = '10.5.6.10'
                    if tmp_host == repository_domain or tmp_host.replace(".", "").replace(":",
                                                                                          "").isdigit():  # 判断host为域名，或者host去掉 .点 和 :冒号 后为数字(即为IP)，【表示取的镜像没错】
                        # 删除host/ 后，用:分割，确认长度为 2或者1
                        tmp_tag1 = tmp_img1.replace("%s/" % tmp_host,
                                                    "")  # 删除掉host/ ,tmp_tag1 = 'sensenebula-guard-std/senseguard-timezone-management:1.2.0-v1.2.0-eedcba'
                        tmp_tag2 = tmp_tag1.split(
                            ":")  # tmp_tag2 = ['sensenebula-guard-std/senseguard-timezone-management','1.2.0-v1.2.0-eedcba']

                        if len(tmp_tag2) == 2:  # 当长度正好为2，说明正常
                            self.name = tmp_tag2[0]
                            self.tag = tmp_tag2[-1]
                            self.judge_images()
                        elif len(tmp_tag2) == 1:  # 说明没有版本号
                            self.name = tmp_tag2[0]
                            self.tag = "latest"
                            self.judge_images()
                        else:  # 服务名 + version = 2，仅服务名 = 1，如果不为2，也不为1，则报错
                            print("Error : images [%s] length abnormal ,pleast check" % tmp_img1)
                    else:  # 当镜像的头不为ip，也不为域名
                        tmp_tag1 = tmp_img1  # tmp_tag1 = 'jaegertracing/spark-dependencies'
                        tmp_tag2 = tmp_tag1.split(":")

                        if len(tmp_tag2) == 2:  # 当长度正好为2，说明正常
                            self.name = tmp_tag2[0]
                            self.tag = tmp_tag2[-1]
                            if not self.images_dict.get(self.name):
                                print("Warning : tag head not is domain or ip, That image is [%s]" % tmp_img1)
                            self.judge_images()
                        elif len(tmp_tag2) == 1:  # 说明没有版本号
                            self.name = tmp_tag2[0]
                            self.tag = "latest"
                            self.judge_images()
                        else:  # 服务名 + version = 2，仅服务名 = 1，如果不为2，也不为1，则报错
                            print("Error : images [%s] length abnormal ,pleast check" % tmp_img1)
        # print(self.images_dict)

    def judge_images(self):
        if self.images_dict.get(self.name) and self.images_dict.get(self.name) == self.tag:  # 存在此key,且value相同
            pass
        elif self.images_dict.get(self.name) and self.images_dict.get(self.name) != self.tag:  # 存在此key,但value不同
            print("Error : image [%s] has multiple version, old version [%s], new version [%s]" % (
                self.images_dict.get(self.name), self.tag))
        else:  # 说明正确，赋值
            self.images_dict[self.name] = self.tag

    def convert_json(self):
        machine_images = {"images": []}
        for k, v in self.images_dict.items():
            machine_images["images"].append({
                "repository": k,
                "tag": v
            })
        machine_images["images"].append(lack_images)
        self.new_images = machine_images
        # print(self.new_images)


class merge_charts_images:
    def __init__(self, charts, images):
        self.charts = charts
        self.images = images

    def merge(self):
        if not self.charts["charts"]:
            print("Error: charts dict is empty")
            sys.exit(1)
        elif not self.images["images"]:
            print("Error: images dict is empty")
            sys.exit(1)
        else:
            self.charts.update(self.images)
            self.all_data = self.charts
        # print(self.all_data)
        with open(new_json_file, 'w') as fp:
            json.dump(self.all_data, fp)
        print("Info : Successful get of charts and images version to [%s] \n" % new_json_file)
