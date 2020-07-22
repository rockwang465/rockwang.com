#!/usr/bin/env python
# -*- coding: utf-8 -*-
import os
import requests
import json
import re
import sys

gpu_list = []
SERVER = "http://127.0.0.1:8000"
TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIn0.leezzwBP0ZugGY0RgqkPgQk5zVmEj8l1NP7nPX5yJHo"
license_quotas_info = {"tfd": "engine-timespace-face-feature-db-worker-nodes",
                       "stfd": "engine-timespace-ped-feature-db-worker-nodes",
                       "ips": "engine-image-face-process-service-nodes"}


def _get_instances(gpu_nu, server_name):
    url = "%s/v1/instances" % SERVER
    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer %s" % TOKEN
    }
    resp = requests.get(url, headers=headers)

    results = []
    if server_name == 'tfd':
        if resp.status_code == 200:
            items = resp.json().get("instances", [])
            if not items:
                _print_error("instance list empty!!!", resp.text)
            results1 = [item for item in items if item.get("name") in ["engine-timespace-feature-db-nebula"]]
            results2 = [item for item in items if item.get("name") in ["engine-struct-timespace-feature-db-nebula"]]
            if get_license_permission("tfd"):
                for item in results1:
                    cfg = item.get('config')
                    # print(cfg)
                    if cfg == '{}\n':
                        _print_error(
                            "not found engine-image-process-service-nebula server override.yaml configuration .", cfg)
                    else:
                        item['config'] = re.sub(r'(NVIDIA_VISIBLE_DEVICES\s*value:\s*)"\d*"',
                                                r'\g<1>"{}"'.format(gpu_nu), cfg)
            if get_license_permission("stfd"):
                for item in results2:
                    cfg = item.get('config')
                    # print(cfg)
                    if cfg == '{}\n':
                        _print_error(
                            "not found engine-image-process-service-nebula server override.yaml configuration .", cfg)
                    else:
                        item['config'] = re.sub(r'(NVIDIA_VISIBLE_DEVICES\s*value:\s*)"\d*"',
                                                r'\g<1>"{}"'.format(gpu_nu), cfg)
        else:
            _print_error("get instance list failed!!!", resp.text)
        results.append(results1[0])
        results.append(results2[0])
    elif server_name == 'ips':
        if resp.status_code == 200:
            items = resp.json().get("instances", [])
            if not items:
                _print_error("instance list empty!!!", resp.text)
            results = [item for item in items if item.get("name") in ["engine-image-process-service-nebula"]]
            for item in results:
                cfg = item.get('config')

                if cfg == '{}\n':
                    _print_error("not found engine-image-process-service-nebula server override.yaml configuration .",
                                 cfg)
                else:
                    item['config'] = re.sub(r'(NVIDIA_VISIBLE_DEVICES\s*value:\s*)"\d*"', r'\g<1>"{}"'.format(gpu_nu),
                                            cfg)
        else:
            _print_error("get instance list failed!!!", resp.text)
    return results


def _print_error(title, message):
    print("#################-ERROR-#################\n")
    print("%s\n%s\n" % (title, message))
    print("#################-ERROR-#################\n")
    sys.exit(1)


def _deploy_instance(instance):
    url = "%s/v1/instances" % SERVER
    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer %s" % TOKEN
    }
    payload = {
        "instance": instance,
        "recreate_pods": True
    }
    resp = requests.post(url, headers=headers, data=json.dumps(payload))
    if resp.status_code != 200:
        _print_error("deploy instance failed!!!", resp.text)


def upgrade_instance():
    tag = False
    gpus = check_gpu()
    gpu_nu = gpu_without_vps(gpus)  # 找到适合非vps服务占用的显卡(即tfd 和 struct-tfd 和 ips 3个服务使用的显卡)
    ips_exist = 0
    gpu_search_exist = 0  # tfd or stfd
    for item in gpus:
        if item.get('gpu') == gpu_nu:  # 如果是合适安装tfd ips struct-tfd的gpu，则pass
            if item.get('ips'):
                ips_exist += 1
            if item.get('tfd'):
                gpu_search_exist += 1
        else:  # 检查在不合适的gpu上是否有不合适的服务占用，有则修改配置重新安装服务
            if item.get('ips') != 0 and get_license_permission('ips'):
                ips_exist += 1
                print("开始重新部署ips")
                instances = _get_instances(gpu_nu, 'ips')
                for instance in instances:
                    _deploy_instance(instance)
                tag = True
            if item.get('tfd') != 0:
                ips_exist += 1
                print("开始重新部署tfd及stfd")
                instances = _get_instances(gpu_nu, 'tfd')
                for instance in instances:
                    _deploy_instance(instance)
                tag = True

    if ips_exist == 0:
        reupdate_instance(gpu_nu, 'ips')
        tag = True
    if gpu_search_exist == 0:
        reupdate_instance(gpu_nu, 'tfd')
        tag = True

    if tag:
        print('vps has been adjusted')
    else:
        print('there is not gpu should be adjustment')


# 找到适合非vps服务占用的显卡(即tfd 和 struct-tfd 和 ips 3个服务使用的显卡)
def gpu_without_vps(gpus):
    for item in gpus:
        if item.get('vps') != 1:
            return item.get('gpu')


# 判断当前服务是否有license授权
def get_license_permission(server_name):
    server_nodes_name = license_quotas_info[server_name]
    quotas_wc_cmd = "license_client status | grep %s | wc -l" % server_nodes_name
    quotas_num_cmd = "license_client status | grep %s | awk -F} '{print $(NF-1)}'  |awk '{print $NF}'" % server_nodes_name
    quotas_wc = os.popen(quotas_wc_cmd).read().split()[0]
    quotas_num = os.popen(quotas_num_cmd).read().split()[0]
    if int(quotas_wc) == 1 and int(quotas_num) > 1:  # 说明有授权
        return True  # license已授权
    else:
        return False


# 如果服务授权了，但没有占显卡(可能服务pending了)
def reupdate_instance(gpu_nu, server_name):
    if get_license_permission(server_name):
        instances = _get_instances(gpu_nu, server_name)
        for instance in instances:
            _deploy_instance(instance)


def check_gpu():
    gpu_nu = os.popen('nvidia-smi -L|wc -l')  # 4张卡，数字为4
    nu = int(gpu_nu.read().split()[0])  # 4
    for num in range(nu):
        # 单张卡id上的信息判断
        tfd_nu = os.popen('nvidia-smi -i %d | grep -w %d | grep -w "C" | grep "search-worker" | wc -l' % (num, num))
        tfd = int(tfd_nu.read().split()[0])
        vps_nu = os.popen(
            'nvidia-smi -i %d | grep -w %d | grep -w "C"| grep "video-process-service-worker" | wc -l' % (num, num))
        vps = int(vps_nu.read().split()[0])
        ips_nu = os.popen(
            'nvidia-smi -i %d | grep -w %d | grep -w "C"| grep "engine-image-process-service" | wc -l' % (num, num))
        ips = int(ips_nu.read().split()[0])

        gpu_list.append(
            {
                'gpu': num,
                'tfd': tfd,
                'vps': vps,
                'ips': ips
            }
        )
    # print(gpu_list)  # [{'gpu': 0, 'ips': 0, 'vps': 1, 'tfd': 0}, {'gpu': 1, 'ips': 1, 'vps': 0, 'tfd': 2}, {'gpu': 2, 'ips': 0, 'vps': 1, 'tfd': 0}, {'gpu': 3, 'ips': 0, 'vps': 1, 'tfd': 0}]
    return gpu_list


if __name__ == '__main__':
    upgrade_instance()
