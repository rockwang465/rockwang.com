#!/usr/bin/env python
# -*- coding: utf-8 -*-

import yaml
import os
import re

docker_host_10 = "10.5.6.10"
docker_host_loacl = "10.5.6.68:5000"

charts_groups = ['addons_charts', 'component_charts', 'console_charts', 'devops_charts', 'guard_charts',
                 'nebula_charts']


# get old charts information from version.yaml
def get_old_charts():
    data = open('./versions.yaml', 'r')
    old_charts = yaml.full_load(data.read())
    # print(old_charts)
    return old_charts


# get current machine all charts to the list data
def get_current_machine_charts():
    # result1 = os.popen("kubectl get pods --all-namespaces -o yaml | grep 'image:' | awk '{print $NF}' | sort -rn | uniq | sed 's#%s##g' | sed 's#%s##g' | sed 's#^/##g' | grep ':' | sort -rn | uniq" % (docker_host_10, docker_host_loacl))
    # result1 = os.popen("kubectl get pods --all-namespaces -o yaml | grep 'image:' | awk '{print $NF}' | sort -rn | uniq | sed 's#%s##g' | sed 's#%s##g' | sed 's#^/##g' | grep ':' | sort -rn | uniq | grep -v '583ca717-2019050520'" % (docker_host_10, docker_host_loacl))
    result1 = os.popen("helm list --col-width 200 | sed 1d | awk '{print $9}' ")
    result2 = result1.readlines()
    # 取服务名 : helm list --col-width 200 | sed 1d | awk '{print $9}' | sed -r 's/([a-zA-Z-]*)-[v|0-9]?.*/\1/'
    # 取版本号 : helm list --col-width 200 | sed 1d | awk '{print $9}' | sed -r 's/[a-zA-Z-]*-([v|0-9]?.*)/\1/'
    # 注意: engine-video在上面用可能会受影响，结果可能只有engine了，因为以-v做切割的

    # convert charts string data to a list data
    machine_charts = {}
    for i in result2:
        chart1 = i.replace("\n", "").split(",")
        chart2 = chart1[0].split(":")[0]

        chart_name = re.findall(r"(.+?)-[0-9]|v[0-9]", chart2)

        # when chart_name == []:
        if chart_name == []:
            break
        # when '-v1' in chart_name[0]:
        elif chart_name[0]:
            if '-v1' in chart_name[0]:
                chart_name = re.findall(r"(.+?)-v[0-9]", chart2)
                # print(chart_name)

        # when chart_name == ['']:
        if chart_name == ['']:
            chart_name = re.findall(r"(.+?)-v[0-9]|[0-9]", chart2)
            # print("\n chart2 --> %s" % chart2)
            chart_name = chart_name[0]
            # print("chart_name --> %s" % chart_name)
            chart_version = re.split(r'%s-' % chart_name, chart2)[-1]
            # print("chart_version----> %s" % chart_version)

            machine_charts[chart_name] = chart_version
        # else is correct value:
        else:
            # print("\n chart2 --> %s" % chart2)
            chart_name = chart_name[0]
            # print("chart_name --> %s" % chart_name)
            chart_version = re.split(r'%s-' % chart_name, chart2)[-1]
            # print("chart_version----> %s" % chart_version)
            machine_charts[chart_name] = chart_version
    # print(machine_charts)
    return machine_charts


# compare old and new charts version , and save new charts version ,and output to new yaml file
def compare_charts(old_charts, new_charts):
    # old_charts example : {'devops_charts': [{'version': '6.5.4-master-a9e66f5', 'namespace': 'logging', 'name': 'elasticsearch'}, ...}
    # new_charts example : {'senseguard-oauth2': '1.2.0-v1.2.0-001-cbe1b9', 'senseguard-bulk-tool': '1.2.0-v1.2.0-001-46db20',...}
    new_charts_version_dict = old_charts

    # loop the charts groups,for example: 'addons_charts', 'component_charts', 'console_charts', 'devops_charts' ...
    for group in charts_groups:

        for index, charts in enumerate(old_charts[group]):
            # print(charts)
            if new_charts.get(charts["name"]):
                # print(charts["name"])
                if new_charts_version_dict[group][index]["version"] == new_charts.get(charts["name"]):
                    print("[Info] : [%s] chart version is right" % charts["name"])
                    # print(new_charts.get(charts["name"]))
                    # print(new_charts_version_dict[group][index]["version"])
                else:
                    print("[Warning] : [%s] chart version not right, the old version is : %s , the new version is : %s" % (charts["name"], new_charts_version_dict[group][index]["version"], new_charts.get(charts["name"])))
                    new_charts_version_dict[group][index]["version"] = new_charts.get(charts["name"])
            else:
                print("[Error] : The current machine was not found this chart name : %s" % charts["name"])
                # print(charts)
                # print(new_charts)

    f = open('./new_versions.yaml', 'w')
    yaml.dump(old_charts, f)


old_charts = get_old_charts()
new_charts = get_current_machine_charts()
compare_charts(old_charts, new_charts)
