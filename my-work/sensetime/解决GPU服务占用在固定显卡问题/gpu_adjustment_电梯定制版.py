#!/usr/bin/env python
# -*- coding: utf-8 -*-
import os
import requests
import json
import re
import sys
import time
import subprocess

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


# 1.对应版本: SenseNebula-E-v1.0.0(2020-08-18)
# 2.脚本总逻辑:
#   a. 必须基于正确逻辑的license授权后，进行以下操作。
#      示例: 非正规授权，如2张显卡，跑2个物体，2个空梯，1个人群，那肯定不行，因为这必须要3张显卡才能做到，那么就说明license授权不对。
#   b. 物体(es-pach-process)和空梯(es-empty-process)一定是共享的，不存在独占情况。
#   c. 物体和空梯是安装一一对应的顺序固定显卡的。即物体podA和空梯podA在显卡0上。
#   d. 其他算法服务，必须跑在非物体、非空梯的显卡上。且其他算法都是独占，不存在共享。
# 3.细化逻辑:
#   a. 先查询各个license的授权，将nodes值放入 license_vps_quotas_info 中("nodes": int)。
#   b. 从sophon上获取vps的override文件，将空梯和物体的replicas的值改为 license_vps_quotas_info 中的nodes值。
#   c. 判断vps的物体和空梯部署的卡上是否有其他服务。
#      c1. 如果有其他服务，调用函数[看是否有可用显卡了](查询卡id，查询物体和空梯哪个值大，拿卡id总数 - 物体/空梯最大值 ，如果大于等于1，拿到可用的显卡id)，
#          将可用的显卡id，固定到其他服务上。
#      c2. 如果有其他服务，发现没有可用的显卡id，则将error日志打印到/opt/gpu_adjustment.err中(日志要有时间显示)
#      c3. 如果一切正常，则将info日志打印到/opt/gpu_adjustment.log中(日志要有时间显示)


# 0.定义quotas
# license_vps_quotas_info = {
#     "veps": "engine-video-es-process-service-nodes",  # 摔倒
#     "vpps": "engine-video-pach-process-service-nodes",  # 结构化
#     "vcops": "engine-video-crowd-oversea-process-service-nodes",  # 人群
#     "veeps": "engine-video-es-empty-process-service-nodes",  # 空梯
#     "vepps": "engine-video-es-pach-process-service-nodes",  # 物体
# }

# 1.定义vps列表
vps_service_name = [
    'engine-video-es-process-service-worker',
    'engine-video-pach-process-service-worker'
    'engine-video-crowd-max-process-service-worker',
    'engine-video-es-empty-process-service-worker',
    'engine-video-es-pach-process-service-worker']

# vps_service_info = []  # 记录vps每个服务的名称、pid
vps_service_info = [
    {  # 摔倒
        'name': 'engine-video-es-process-service-worker',
        "veps": "engine-video-es-process-service-nodes",
        'status': 'alone'
    },
    {  # 结构化
        'name': 'engine-video-pach-process-service-worker',
        "vpps": "engine-video-pach-process-service-nodes",
        'status': 'alone'
    },
    {  # 人群
        'name': 'engine-video-crowd-max-process-service-worker',
        "vcops": "engine-video-crowd-oversea-process-service-nodes",
        'status': 'alone'
    },
    {  # 空梯
        'name': 'engine-video-es-empty-process-service-worker',
        "veeps": "engine-video-es-empty-process-service-nodes",
        'status': 'share'
    },
    {  # 物体
        'name': 'engine-video-es-pach-process-service-worker',
        "vepps": "engine-video-es-pach-process-service-nodes",
        'status': 'share'
    }
]


class vps_info:
    def __init__(self):
        self.vps_service_info = vps_service_info
        self.gpu_num = os.popen('nvidia-smi -L | wc -l').read().split()[0]

    # # 1.判断显卡数是否为4，确认是否继续调优
    # def judge_card_num(self):
    #     self.gpu_num = os.popen('nvidia-smi -L | wc -l').read().split()[0]
    #     if self.gpu_num != 4:
    #         print('Warning: 当前机器非4卡，是否继续进行优化?')
    #         res = raw_input('是否继续优化:yes/no: ')
    #         if res not in ['yes', 'YES', 'y', 'Y']:
    #             sys.exit(0)
    #         print("Warning: 10s后开始继续调优...")
    #         time.sleep(10)

    # 2.获取每个服务对应的进程的pid
    def get_vps_pid(self):
        for name in vps_service_name:
            # get_name_cmd = "ps -ef | grep %s | grep -v grep" % name
            get_pid_cmd = "ps -ef | grep %s | grep -v grep | awk '{print $2}'" % name
            # pid = os.popen(get_pid_cmd).read().split()[0]
            # res = subprocess.Popen(cmdstr, shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
            # stdout, stderr = res.communicate()
            # if res.returncode != 0:
            #   raise InstallError('%s\n%s' % (stderr, stdout))

            res = subprocess.Popen(get_pid_cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
            stdout, stderr = res.communicate()
            if res.returncode != 0:
                print("Error : 命令执行失败")
                sys.exit(1)

            # for info in self.vps_service_info:
            #     if info.get('name') == name:
            #         info['pid'] = pid
        print(vps_service_info)

    # 3.获取每个服务的quotas的nodes值
    def get_quotas_nodes(self):
        print("get quotas nodes")

    # 3.判断传进来的服务名在gpu中的pid是否属于共享、独占的状态
    # def gpu_service_check(self, info):
    def gpu_service_check(self):
        # b.轮询每张显卡

        # c.查看每张显卡中的pid
        # d.如果此pid是对应的服务是独占的，确认是否此卡上还有其他pid
        # e.如果此pid是共享的，确认是否还有另一个pid
        print(info)

    # 4.获取每个服务在gpu中占卡的状态是否正常
    # def judge_service_status(self):
    #     for info in self.vps_service_info:
    #         self.gpu_service_check(info)


if __name__ == '__main__':
    # upgrade_instance()

    vps_ops = vps_info()
    # vps_ops.judge_card_num()  # 判断显卡数是否为4，确认是否继续调优
    vps_ops.get_vps_pid()
    vps_ops.get_quotas_nodes()
    # vps_ops.gpu_service_check()

    # judge_service_status()
