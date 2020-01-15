#!/usr/bin/env python
# -*- coding: utf-8 -*-
import paramiko
import re
from kubernetes import client, config
import sys
import json
import copy

kube_config_path = '/root/.kube/config'
kube_config_file = 'kubeconfig'
local_domain = '127.0.0.1'
json_file = 'versions.json'
all_namespace = ['component', 'nebula', 'default', 'logging', 'monitoring']  # 未加: galaxias helm kube-public kube-system
lack_images = [{'repository': 'elasticsearch/busybox', 'tag': 'latest'},
               {'repository': 'gitlabci/golang', 'tag': '1.9-cuda-gcc49-1'},
               {'repository': 'component/mc', 'tag': 'RELEASE.2019-02-13T19-48-27Z'},
               {'repository': 'external_storage/local-volume-provisioner', 'tag': 'v2.3.0'}]
# kubernetes base images
k8s_images = [{'repository': 'nvidia/k8s-device-plugin', 'tag': '1.10'},
              {'repository': 'kubernetes/nginx-ingress-controller', 'tag': '0.21.0'},
              {'repository': 'kubernetes/kube-scheduler', 'tag': 'v1.13.2'},
              {'repository': 'kubernetes/kubernetes-dashboard-amd64', 'tag': 'v1.10.1'},
              {'repository': 'kubernetes/kube-proxy', 'tag': 'v1.13.2'},
              {'repository': 'kubernetes/kube-controller-manager', 'tag': 'v1.13.2'},
              {'repository': 'kubernetes/kube-apiserver', 'tag': 'v1.13.2'},
              {'repository': 'kubernetes/etcd', 'tag': '3.2.24'},
              {'repository': 'kubernetes/defaultbackend', 'tag': '1.4'},
              {'repository': 'kubernetes/coredns', 'tag': '1.2.6'},
              {'repository': 'external_storage/local-volume-provisioner', 'tag': 'v2.3.0'},
              {'repository': 'coreos/flannel', 'tag': 'v0.10.0-amd64'},
              {'repository': 'kubernetes/tiller', 'tag': 'v2.13.1'},
              {'repository': 'kubernetes/pause', 'tag': '3.1'}]

# custom version.yaml request template
addons_charts = ['local-volume-provisioner', 'kubernetes-dashboard', 'nginx-ingress']
component_charts = ['kafka', 'zookeeper', 'cassandra', 'minio', 'osg', 'seaweedfs', 'redisoperator', 'mysql', 'emqx']
devops_charts = ['elasticsearch', 'elasticsearch-curator', 'filebeat', 'prometheus-operator', 'jaeger-operator']
common_charts = {'addons_charts': [], 'devops_charts': [], 'component_charts': [], 'guard_charts': [],
                 'nebula_charts': []}

ssh = paramiko.SSHClient()


class get_charts_version:
    # get helm charts version in standard environment
    def get_helm_charts(self, args):
        # ssh = paramiko.SSHClient()
        cut_line = '{print $8" "$9" "$NF}'  # cut charts and namespace
        try:
            ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
            ssh.connect(args.env_ip, args.env_port, args.env_username, args.env_passwd)
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
                    info = i.encode('utf-8').split()
                    if info[0] == "FAILED":
                        print("Error : Helm sever status is FAILED, server info is [%s]" % info[1])
                        sys.exit(1)
                    else:
                        self.helm_list["charts"].append({
                            "info": info[1],
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
            chart_name = re.findall(r"(.+?)-([0-9]|v[0-9])", i["info"])[0][0]

            chart_version = re.split(r'%s-' % chart_name, i["info"])[-1]
            machine_charts["charts"].append({
                "name": chart_name,
                "version": chart_version,
                "namespace": i["namespace"]
            })
        self.machine_charts = machine_charts
        # print(self.new_charts)  # 备用勿删

    # convert to common charts
    def convert_common_charts(self):
        for k in self.machine_charts["charts"]:
            name = k.get("name")
            ns = k.get("namespace")
            if name in addons_charts:
                common_charts["addons_charts"].append(k)
            elif name in devops_charts:
                common_charts["devops_charts"].append(k)
            elif name in component_charts:
                common_charts["component_charts"].append(k)
            else:
                if ns == "nebula":
                    common_charts["nebula_charts"].append(k)
                elif "senseguard" in name or "aurora" in name:
                    common_charts["guard_charts"].append(k)
                else:
                    print("Error: Have a error key [%s]" % k)
                    sys.exit(1)
        self.new_charts = common_charts
        # print(self.new_charts)
        # print("\n")


# get docker image version in standard environment(通过k8s api获取docker image版本号)
class get_images_version:
    def get_kube_config(self, args):
        try:
            ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
            ssh.connect(args.env_ip, args.env_port, args.env_username, args.env_passwd)
            stdin, stdout, stderr = ssh.exec_command("cat %s" % kube_config_path)
            # print(stdout.read().decode())
            self.kube_config_body = stdout.read().decode()
            ssh.close()

            # 当config中含有127.0.0.1时，替换为环境的ip。
            if local_domain in self.kube_config_body:
                kube_file = self.kube_config_body.replace(local_domain, args.env_ip)
            else:
                kube_file = self.kube_config_body

            with open(kube_config_file, 'w') as fp:
                fp.write(kube_file)
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
                self.name, self.images_dict.get(self.name), self.tag))
            sys.exit(1)
        else:  # 说明正确，赋值
            self.images_dict[self.name] = self.tag

    def convert_json(self):
        machine_images = {"images": []}
        for k, v in self.images_dict.items():
            machine_images["images"].append({
                "repository": k,
                "tag": v
            })

        for i in lack_images:
            machine_images["images"].append(i)
        # machine_images["images"].append(lack_images)
        self.new_images = machine_images
        # print(self.new_images)


class merge_charts_images:
    # 生成的为所有的images+charts的versions.json文件，给于一键部署使用的。
    def merge_to_versions(self, charts, images):
        self.charts1 = copy.deepcopy(charts)
        self.images1 = copy.deepcopy(images)

        for i in k8s_images:
            self.images1['images'].append(i)

        self.charts1.update(self.images1)
        self.all_version = self.charts1
        # print(self.all_version)
        with open(json_file, 'w') as fp:
            json.dump(self.all_version, fp)
        print("Info : Successful get of charts and images version to [%s] \n" % json_file)
