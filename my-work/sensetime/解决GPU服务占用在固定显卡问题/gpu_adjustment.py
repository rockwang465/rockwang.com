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
            for item in results1:
                cfg = item.get('config')
                item['config'] = re.sub(r'(NVIDIA_VISIBLE_DEVICES\s*value:\s*)"\d*"', r'\g<1>"{}"'.format(gpu_nu), cfg)
            for item in results2:
                cfg = item.get('config')
                item['config'] = re.sub(r'(NVIDIA_VISIBLE_DEVICES\s*value:\s*)"\d*"', r'\g<1>"{}"'.format(gpu_nu), cfg)
        else:
            _print_error("get instance list failed!!!", resp.text)
        results.append(results1[0])
        results.append(results2[0])
    elif server_name == 'ips':
        if resp.status_code == 200:
            items = resp.json().get("instances", [])
            if not items:
                _print_error("instance list empty!!!", resp.text)
            results = [item for item in items if item.get("name") in ["engine-image-ingress-service-nebula"]]
            for item in results:
                cfg = item.get('config')
                item['config'] = re.sub(r'(NVIDIA_VISIBLE_DEVICES\s*value:\s*)"\d*"', r'\g<1>"{}"'.format(gpu_nu), cfg)
        else:
            _print_error("get instance list failed!!!", resp.text)
    return results


def _print_error(title, message):
    print("#################-ERROR-#################\n")
    print("%s\n%s\n" % (title, message))
    print("#################-ERROR-#################\n")


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
    # usable_card = get_usable_gpu(gpus)
    # print("usable gpu card id : %s" % usable_card)
    gpu_nu = gpu_without_vps()  # 找到适合非vps服务占用的显卡(即tfd 和 struct-tfd 和 ips 3个服务使用的显卡)
    # print("gpu_nu value: %s" % gpu_nu)
    for item in gpus:
        if item.get('gpu') == gpu_nu:
            pass
        else:
            print("else")
            if item.get('ips') != 0:
                print("准备重新部署ips")
                # instances = _get_instances(gpu_nu, 'ips')
                # for instance in instances:
                #     _deploy_instance(instance)
                # tag = True
            if item.get('tfd') != 0:
                print("准备重新部署tfd")
                # instances = _get_instances(gpu_nu, 'tfd')
                # for instance in instances:
                #     _deploy_instance(instance)
                # tag = True
    if tag:
        print('vps has been adjusted')
    else:
        print('there is not gpu should be adjustment')


# 找到适合非vps服务占用的显卡(即tfd 和 struct-tfd 和 ips 3个服务使用的显卡)
def gpu_without_vps():
    gpus = check_gpu()
    print(gpus)
    for item in gpus:
        if item.get('vps') != 1:
            return item.get('gpu')


def check_gpu():
    gpu_nu = os.popen('nvidia-smi -L|wc -l')  # 4张卡，数字为4
    nu = int(gpu_nu.read().split()[0])  # 4
    for num in range(nu):
        tfd_nu = os.popen('nvidia-smi -i %d | grep -w %d | grep -w "C"|grep "search-worker"|wc -l' % (num, num))
        tfd = int(tfd_nu.read().split()[0])
        vps_nu = os.popen(
            'nvidia-smi -i %d | grep -w %d | grep -w "C"| grep "video-process-service-worker"|wc -l' % (num, num))
        vps = int(vps_nu.read().split()[0])

        ips_nu = os.popen(
            'nvidia-smi -i %d | grep -w %d | grep -w "C"| grep "engine-image-process-service"|wc -l' % (num, num))
        ips = int(ips_nu.read().split()[0])

        gpu_list.append(
            {
                'gpu': num,
                'tfd': tfd,
                'vps': vps,
                'ips': ips
            }
        )
    # print(gpu_list)  [{'gpu': 0, 'vps': 0, 'tfd': 2}, {'gpu': 1, 'vps': 1, 'tfd': 0}, {'gpu': 2, 'vps': 1, 'tfd': 0}, {'gpu': 3, 'vps': 1, 'tfd': 0}]
    return gpu_list


# # 找到适合非vps服务占用的显卡(即tfd 和 struct-tfd 和 ips 3个服务使用的显卡)
# def get_usable_gpu(gpus):
#     usable_gpu = []
#     for item in gpus:
#         if item.get('vps') == 0:
#             usable_gpu.append(item.get('gpu'))
#
#     if len(usable_gpu) == 1:
#         usable_gpu_id = usable_gpu[0]
#     elif len(usable_gpu) == 0:
#         _print_error("Not found usable gpu card", "")
#         sys.exit(1)
#     else:
#         usable_gpu_id = usable_gpu[0]
#     return usable_gpu_id


if __name__ == '__main__':
    upgrade_instance()
