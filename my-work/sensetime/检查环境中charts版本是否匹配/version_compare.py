#!/usr/bin/env python
# encoding: utf-8

import subprocess
import os
import sys
import re

namespace = ["kube-system", "component", "logging", "monitoring", "nebula", "default"]


class get_charts_version():
    def __init__(self):
        self.env_charts_info = {}  # 当前环境的charts信息 {"nebula": [{"servername":"aurora","version":"v1.0.0-master-xxxx"}]}
        self.namespace = namespace

    def helm_version(self):
        awk_cmd = '{print $8" "$9" "$NF}'  # DEPLOYED access-control-process-1.2.0-master-new-db-c124195 nebula
        cmdstr = "helm list --col-width 200 | sed 1d | awk '%s'" % awk_cmd
        res = subprocess.Popen(cmdstr, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        if res.stderr.read():
            print("Error: failure to execute [%s]" % cmdstr)
            sys.exit(1)
        else:
            helm_list1 = res.stdout.read()
            if helm_list1:
                helm_list2 = helm_list1.split("\n")
            else:
                print("Error: helm list is empty")
                sys.exit(1)

        for i in helm_list2:
            chart_info = i.split()
            if chart_info:
                ns = chart_info[-1]
                if ns in self.namespace:
                    if not self.env_charts_info.get(ns):
                        self.env_charts_info[ns] = []

                    server_info = chart_info[1]
                    if chart_info[0] == "DEPLOYED":
                        pass
                    elif chart_info[0] == "FAILED":
                        print("Warning: [%s] helm status is FAILED" % server_info)
                    else:
                        print("Error: [%s] status is [%s], please check" % (server_info, chart_info[0]))
                        sys.exit(1)
                    server_name = re.findall(r"(.+?)-([0-9]|v[0-9])", server_info)[0][0]
                    server_version = server_info.split("%s-" % server_name)[-1]
                    chart_dict = {"version": server_version, "server_name": server_name}
                    self.env_charts_info[ns].append(chart_dict)
                else:
                    print("Error: [%s] namespace not in [%s]" % (ns, self.namespace))
        # print("\n")
        # print("Info : evn_charts_info : %s" % self.env_charts_info)


class compare_charts_version:
    def __init__(self, compare_helm_file):
        self.compare_helm_file = compare_helm_file

    def get_version_file(self):
        with open(self.compare_helm_file, 'r') as fp:
            self.compare_file_data = fp.read().split("\n")

    def compare_helm(self, env_charts_info):
        print("\n")
        for chart in self.compare_file_data:
            if chart:
                file_chart_info = chart.split()
                # print("chart_info: %s" % file_chart_info)
                ns = file_chart_info[0]
                server_name = file_chart_info[1]
                server_version = file_chart_info[-1]

                n = 0
                for machine_chart in env_charts_info.get(ns):
                    if server_name == "infra-object-storage-gateway":
                        server_name = "osg"
                    elif server_name == "gpu-searcher":
                        server_name = "engine-timespace-feature-db"
                    elif server_name == "redis":
                        server_name = "redisoperator"
                    elif server_name == "prometheus":
                        server_name = "prometheus-operator"
                    if machine_chart.get("server_name") == server_name:  # 服务名相同
                        n += 1
                        if machine_chart.get("version") == server_version:
                            continue
                            # print("Info: [%s-%s] == [%s-%s]" % (
                            #     server_name, server_version, server_name, machine_chart.get("version")))
                        else:
                            print("Error: [%s] service environment version [%s], [%s] file service version [%s]" % (
                                server_name, machine_chart.get("version"), self.compare_helm_file, server_version))
                if n == 0:
                    print("Error: [%s] service not found, please check" % server_name)
                    sys.exit(1)
