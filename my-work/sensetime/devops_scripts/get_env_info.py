#!/usr/bin/env python
# -*- coding: utf-8 -*-

# ---------------------------- #
# version | 1.0.0              #
# ---------------------------- #

import platform
import psutil
import os
import argparse
import subprocess
import re
import sys
import yaml

engine_namespace = 'engine'
addons_charts = ['local-volume-provisioner', 'kubernetes-dashboard', 'nginx-ingress']
guard_charts = ['gateway', 'aurora', 'coral', 'device-manager-service', 'auth']
component_charts = ['kafka', 'zookeeper', 'cassandra', 'minio', 'infra-object-storage-gateway', 'seaweedfs', 'redis',
                    'mysql', 'emqx', 'haproxy', 'infra-model-manager', 'nacos', 'etcd', 'elasticsearch-curator']
devops_charts = ['logstash', 'kibana', 'filebeat', 'prometheus-operator', 'apm-server',
                 'jaeger-operator', 'infra-console-service', 'infra-frontend-service']
all_charts = {'addons_charts': [], 'devops_charts': [], 'component_charts': [], 'guard_charts': [],
              'engine_charts': []}
charts_version_name = 'charts_version.yaml'
vps_worker_name = ['engine-video-crowd-max-process-service-worker', 'engine-video-es-pach-process-service-worker',
                   'engine-video-es-process-service-worker']
model_object_type = ['crowd', 'face', 'pach']


# input arguments
def parse_args():
    parser = argparse.ArgumentParser(description='Get the cluster info')
    parser.add_argument("host_ip", help="the cluster environment ip address")
    return parser.parse_args()


def print_info():
    print("+---------------------------------------+")
    print("+ support version: E-V1.x.0 version     +")
    print("+---------------------------------------+\n")


class GetBaseInfo:
    def __init__(self):
        self.processor_count = 0
        self.cpu_model = ""
        self.cpu_core_id = {}
        self.cpu_core_num = 0
        self.disk_name_cmd = "df -h | grep '/dev/sd' | awk '{print $NF}'"

    # get linux operator system information
    @staticmethod
    def get_os_info():
        os_version_info = platform.linux_distribution()  # ('CentOS Linux', '8.2.2004', 'Core')
        os_version = os_version_info[0] + os_version_info[1]
        os_bits_info = platform.architecture()  # ('64bit', '')
        os_bits = os_bits_info[0]
        print("+------------------------system info------------------------+")
        print("os type: [%s]" % platform.system())  # Linux
        print("os kernel release: [%s]" % platform.release())  # 4.18.0-193.6.3.el8_2.x86_64
        print("os machine version: [%s]" % platform.machine())  # x86_64
        print("os version: [%s]" % os_version)  # CentOS Linux8.2.2004
        print("os bits version: [%s]" % os_bits)  # 64bit
        print("os host name: [%s]" % platform.node())  # nebula-ce4869
        print("+------------------------system info------------------------+\n")

    # get cpu information
    def get_cpu_info(self):
        """
        Return the information in /proc/cpuinfo
        as a dictionary in the following format:
        cpu_info['proc0']={...}
        cpu_info['proc1']={...}
        """
        file_cpu_info = open("/proc/cpuinfo")
        try:
            for line in file_cpu_info:
                if line.find("processor") == 0:
                    self.processor_count += 1
                elif line.find("model name") == 0:
                    if self.cpu_model == "":
                        self.cpu_model = line.split(":")[1].strip()
                elif line.find("physical id") == 0:
                    physical_id = line.split(":")[1].strip()
                    self.cpu_core_id[physical_id] = True
                    self.cpu_core_num = len(self.cpu_core_id)
            print("+------------------------cpu    info------------------------+")
            print("cpu processor: [%s]" % self.processor_count)
            print("cpu model: [%s]" % self.cpu_model)
            print("cpu physical id num: [%s]" % self.cpu_core_num)
            print("+------------------------cpu    info------------------------+\n")
        finally:
            file_cpu_info.close()

    # get memory information
    @staticmethod
    def get_mem_info():
        pc_mem = psutil.virtual_memory()
        div_gb_factor = (1024.0 ** 3)
        print("+------------------------memory info------------------------+")
        print("memory total: [%.2fGB]" % float(pc_mem.total / div_gb_factor))
        print("memory available: [%.2fGB]" % float(pc_mem.available / div_gb_factor))
        print("memory used: [%.2fGB]" % float(pc_mem.used / div_gb_factor))
        print("memory used percent: [%.2f%%]" % float(pc_mem.percent))
        print("memory free: [%.2fGB]" % float(pc_mem.free / div_gb_factor))
        print("+------------------------memory info------------------------+\n")

    def get_disk_info(self):
        res = subprocess.Popen(self.disk_name_cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
        stdout, stderr = res.communicate()
        if res.returncode != 0:
            print("Error: get disk name failed, %s" % stderr)
            sys.exit(1)
        disk_name = stdout.decode('utf-8')
        print("+------------------------disk   info------------------------+")
        for name in disk_name.split("\n"):
            if name != "":
                disk = os.statvfs(name)
                percent = (disk.f_blocks - disk.f_bfree) * 100 / (disk.f_blocks - disk.f_bfree + disk.f_bavail)
                print("disk used percent: [%.2f%%], disk mount name: [%s]" % (float(percent), name))
        print("+------------------------disk   info------------------------+\n")


# get helm charts version
class ChartsVersion:
    def __init__(self):
        self.helm_list = {"charts": []}
        self.machine_charts = {"charts": []}
        self.all_charts = {}

    # get helm charts version in standard environment
    def get_helm_charts(self):
        cut_line = '{print $8" "$9" "$NF}'  # cut charts and namespace
        cmd_str = "helm list --col-width 200 | sed 1d | awk '%s' | sort -rn | uniq" % cut_line
        res = subprocess.Popen(cmd_str, shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
        stdout, stderr = res.communicate()
        if res.returncode != 0:
            print("Error: get helm info failed, %s" % stderr)
            sys.exit(1)
        bytes_helm_info = stdout  # type is bytes
        # helm_info = str(bytes_helm_info, encoding='utf-8')
        helm_info = bytes_helm_info.decode('utf-8')

        for i in helm_info.split("\n"):
            if i == "":
                pass
            else:
                info = i.encode('utf-8').split()
                if info[0] == "FAILED":
                    print("Error : Helm sever status is FAILED, server info is [%s]" % info[1])
                    # sys.exit(1)
                else:
                    self.helm_list["charts"].append({
                        "info": info[1].decode('utf-8'),
                        "namespace": info[-1].decode('utf-8')
                    })
        # print(self.helm_list)

    # convert standard environment helm charts version to dict
    def convert_helm_charts(self):
        for i in self.helm_list["charts"]:
            chart_name = re.findall(r"(.+?)-([0-9]|v[0-9])", i["info"])[0][0]

            chart_version = re.split(r'^%s-' % chart_name, i["info"])[-1]
            self.machine_charts["charts"].append({
                "name": chart_name,
                "version": chart_version,
                "namespace": i["namespace"]
            })
        # print(self.machine_charts)

    # convert to common charts
    def convert_all_charts(self):
        for k in self.machine_charts["charts"]:
            name = k.get("name")
            ns = k.get("namespace")
            if name in addons_charts:
                all_charts["addons_charts"].append(k)
            elif name in devops_charts:
                all_charts["devops_charts"].append(k)
            elif name in component_charts:
                all_charts["component_charts"].append(k)
            else:
                if ns == engine_namespace:
                    all_charts["engine_charts"].append(k)
                elif ns == "logging":
                    all_charts["devops_charts"].append(k)
                elif ns == "component":
                    all_charts["component_charts"].append(k)
                elif "senseguard" in name or "aurora" in name or name in guard_charts:
                    all_charts["guard_charts"].append(k)
                else:
                    print("Error: Have a error key [%s]" % k)
                    sys.exit(1)
        self.all_charts = all_charts
        # print(self.all_charts)

    def save_charts_yaml(self):
        with open(charts_version_name, 'w') as fp:
            yaml.dump(self.all_charts, fp)
        print("Info : Successful get of charts version to [%s] \n" % charts_version_name)


# get sdk and algo version
class GetEngineInfo:
    @staticmethod
    def sdk_version():
        for object_type in model_object_type:
            # get pod name
            pod_cmd_str = "kubectl get pods -n %s | grep -e 'video.*worker' | grep Running | grep %s | awk '{print $1}'" % (
                engine_namespace, object_type)
            # print(pod_cmd_str)
            pod_res = subprocess.Popen(pod_cmd_str, shell=True, stdout=subprocess.PIPE,
                                       stderr=subprocess.STDOUT)
            stdout, stderr = pod_res.communicate()
            if pod_res.returncode != 0:
                print("Error: get pod failed, %s" % stderr)
                sys.exit(1)
            pod_name = stdout.decode('utf-8').split("\n")[0]
            # print(pod_name)
            if pod_name == "":
                print("Warning: not found %s object type" % object_type)
                continue

            # get sdk version
            sdk_object_type = ""
            if object_type == "pach":
                sdk_object_type = "structural"

            if sdk_object_type != "":
                sdk_cmd_str = "kubectl exec -it -n %s %s -- ls libs/kestrel -hl | grep %s | sed 's/.*\(lib.*so[.0-9]\{3\}\)/\1/' | tail -1" % (
                    engine_namespace, pod_name, sdk_object_type)
            else:
                sdk_cmd_str = "kubectl exec -it -n %s %s -- ls libs/kestrel -hl | grep %s | sed 's/.*\(lib.*so[.0-9]\{3\}\)/\1/' | tail -1" % (
                    engine_namespace, pod_name, object_type)
            # print(sdk_cmd_str)
            sdk_res = subprocess.Popen(sdk_cmd_str, shell=True, stdout=subprocess.PIPE,
                                       stderr=subprocess.STDOUT)
            stdout, stderr = sdk_res.communicate()
            if sdk_res.returncode != 0:
                print("Error: get sdk failed, %s" % stderr)
                sys.exit(1)
            sdk_version = stdout.decode('utf-8')
            print("+------------------------engine info------------------------+")
            print("object type: [%s]" % object_type)
            print("sdk version: %s" % sdk_version)

            # get model
            #  pipeline_config=`kubectl exec -it -n ${namespace} ${pod} -- ls -hl /config | grep ${ot} | grep pipeline | awk '{print $9}'`
            #  models=`kubectl exec -it -n ${namespace} ${pod} -- cat /config/${pipeline_config} | grep "\"model\"" | sed 's/.*\(KM_.*model\).*/\1/'`
            # a. get pipeline config name
            config_cmd_str = "kubectl exec -it -n %s %s -- ls -hl /config | grep %s | grep pipeline | awk '{print $9}'" % (
                engine_namespace, pod_name, object_type)
            # print(config_cmd_str)
            config_res = subprocess.Popen(config_cmd_str, shell=True, stdout=subprocess.PIPE,
                                          stderr=subprocess.STDOUT)
            stdout, stderr = config_res.communicate()
            if config_res.returncode != 0:
                print("Error: get pipeline config failed, %s" % stderr)
                sys.exit(1)
            config_name = stdout.decode('utf-8').split("\n")[0]

            # b. get model name by pipeline config name
            # shell command: kubectl exec -it -n engine <pod_name> -- cat /config/%s | grep '\"model\"' | sed 's/.*\(KM_.*model\).*/\1/'
            model_cmd_str = "kubectl exec -it -n %s %s -- cat /config/%s | grep '\"model\"' | sed 's/.*\(KM_.*model\).*/\\1/'" % (
                engine_namespace, pod_name, config_name)
            # print(model_cmd_str)
            model_res = subprocess.Popen(model_cmd_str, shell=True, stdout=subprocess.PIPE,
                                         stderr=subprocess.STDOUT)
            stdout, stderr = model_res.communicate()
            if model_res.returncode != 0:
                print("Error: get model failed, %s" % stderr)
                sys.exit(1)
            model_name = stdout.decode('utf-8')
            print("model version: ")
            print(model_name)
            print("+------------------------engine info------------------------+\n")

    @staticmethod
    def algo_version():
        vps_cm_str = "kubectl describe configmap -n %s engine-video-process-service-config  | grep 'com\.sensetime' | grep ref | awk '{print $NF}'" % (
            engine_namespace)
        res = subprocess.Popen(vps_cm_str, shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
        stdout, stderr = res.communicate()
        if res.returncode != 0:
            print("Error: get configmap failed, %s" % stderr)
            sys.exit(1)
        configmap_info = stdout.decode('utf-8')
        if configmap_info != "":
            print("+------------------------algo   info------------------------+")
            for info in configmap_info.split("\n"):
                if info != "":
                    print(info)
            print("+------------------------algo   info------------------------+\n")


if __name__ == '__main__':
    # 0. input arguments and print info
    args = parse_args()
    print_info()

    # 1. get system info
    base_info = GetBaseInfo()
    base_info.get_os_info()
    base_info.get_cpu_info()
    base_info.get_mem_info()
    base_info.get_disk_info()

    # 2. get helm charts version
    chart = ChartsVersion()
    chart.get_helm_charts()
    chart.convert_helm_charts()
    chart.convert_all_charts()
    chart.save_charts_yaml()

    # 3.get sdk and algo version
    engine = GetEngineInfo()
    engine.sdk_version()
    engine.algo_version()
