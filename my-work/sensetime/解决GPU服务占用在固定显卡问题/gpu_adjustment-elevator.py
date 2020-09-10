#!/usr/bin/env python
# -*- coding: utf-8 -*-
import os
import requests
import json
import yaml
import re
import sys
import time
import subprocess

# 1.对应版本: SenseNebula-E-v1.0.0(2020-08-18)
# 2.脚本总逻辑:
#   a. 必须基于正确逻辑的license授权后，进行以下操作。
#      示例: 非正规授权，如2张显卡，跑2个物体，2个空梯，1个人群，那肯定不行，因为这必须要3张显卡才能做到，那么就说明license授权不对。
#   b. 物体(es-pach-process)和空梯(es-empty-process)一定是共享的，不存在独占情况。
#   c. 物体和空梯是安装一一对应的顺序固定显卡的。即物体podA和空梯podA在显卡0上。
#   d. 其他算法服务，必须跑在非物体、非空梯的显卡上。且其他算法都是独占，不存在共享。
# 3.细化逻辑:
#   a. 先查询各个license的授权，将quotas值放入 vps_service_info 中("quotas": int)。
#   b. 判断vps的物体和空梯部署的卡上是否有其他服务。
#      第一次测试: 其他服务都理解为随机占卡:
#      b1. 循环 空梯和物体卡上是否有其他服务。 有则重装vps的每一个服务。 -- 理论上这一层逻辑就够用了

#      b2. 从0 到 空梯quotas-1 和 从0 到物体quotas-1 的卡上是否两个服务都存在。(循环显卡id，拿到进程id，是否in pid列表)
#         b1.1 如果发现异常，则重装vps的每一个服务。

#      第二次测试: 如果随机占卡经常出问题，需要将其他服务固显卡: 下面逻辑暂时先不管
#      b3. 如果有其他服务，调用函数[看是否有可用显卡了](查询卡id，查询物体和空梯哪个值大，拿卡id总数 - 物体/空梯最大值 ，如果大于等于1，拿到可用的显卡id)，
#          将可用的显卡id，固定到其他服务上。
#      b4. 如果有其他服务，发现没有可用的显卡id，则将error日志打印到/opt/gpu_adjustment.err中(日志要有时间显示)
#      b5. 如果一切正常，则将info日志打印到/opt/gpu_adjustment.log中(日志要有时间显示)
#   c. 从sophon上获取vps的override文件，将空梯和物体的replicas的值改为 license_vps_quotas_info 中的nodes值。
# 4.可能缺少的逻辑
#   a. 拿到可用的显卡，给于quotas>=1，且不为空梯、物体的服务，固定显卡。
#      a0. 当物体/空梯显卡上有其他服务则进行以下操作:
#      a1. 找到可用显卡
#      a2. 找到已授权的非物体/空梯服务
#      a3. 固定非物体/空梯服务的显卡id(字典操作，判断如果没有字段，则增加字段)
#      a4. 将当前服务的quotas值放到replicas上(解决下面b提出的问题，此处modify_vps_override函数的for循环部分已解决了)
#      a5. 更新服务
#   b. 轮询每个服务的replicas是否和quotas的值相等，不相等就更新replicas，并更新vps服务。

# 0.定义quotas
# license_vps_quotas_info = {
#     "veps": "engine-video-es-process-service-nodes",  # 摔倒
#     "vpps": "engine-video-pach-process-service-nodes",  # 结构化
#     "vcops": "engine-video-crowd-oversea-process-service-nodes",  # 人群
#     "veeps": "engine-video-es-empty-process-service-nodes",  # 空梯
#     "vepps": "engine-video-es-pach-process-service-nodes",  # 物体
# }

# vps worker name
vps_service_name = [
    "engine-video-es-process-service-worker",
    "engine-video-pach-process-service-worker"
    "engine-video-crowd-max-process-service-worker",
    "engine-video-es-empty-process-service-worker",
    "engine-video-es-pach-process-service-worker",
    "engine-video-headshoulder-process-service-worker"  # 头肩,电梯不会用到此功能
]

# vps_service_info = []  # 记录vps每个服务的名称、pid
# es_worker
# worker

vps_service_info = [
    {  # 摔倒
        "name": "engine-video-es-process-service-worker",
        "nodes_name": "engine-video-es-process-service-nodes",
        "veps": "engine-video-es-process-service-nodes",
        "status": "alone",
        "worker_name": "es_worker"
    },
    {  # 结构化
        "name": "engine-video-pach-process-service-worker",
        "nodes_name": "engine-video-pach-process-service-nodes",
        "vpps": "engine-video-pach-process-service-nodes",
        "status": "alone",
        "worker_name": "pach_worker"
    },
    {  # 人群
        "name": "engine-video-crowd-max-process-service-worker",
        "nodes_name": "engine-video-crowd-oversea-process-service-nodes",
        "vcops": "engine-video-crowd-oversea-process-service-nodes",
        "status": "alone",
        "worker_name": "crowd_oversea_worker"
    },
    {  # 空梯
        "name": "engine-video-es-empty-process-service-worker",
        "nodes_name": "engine-video-es-empty-process-service-nodes",
        "veeps": "engine-video-es-empty-process-service-nodes",
        "status": "share",
        "worker_name": "es_empty_worker"
    },
    {  # 物体
        "name": "engine-video-es-pach-process-service-worker",
        "nodes_name": "engine-video-es-pach-process-service-nodes",
        "vepps": "engine-video-es-pach-process-service-nodes",
        "status": "share",
        "worker_name": "es_pach_worker"
    },
    {  # 头肩 # 电梯不会用到此功能
        "name": "engine-video-headshoulder-process-service-worker",
        "nodes_name": "engine-video-headshoulder-process-service-nodes",
        "vhps": "engine-video-headshoulder-process-service-nodes",
        "status": "alone",
        "worker_name": "headshoulder_worker"
    },
    # {  # 人脸 # 电梯不会用到此功能
    #     "name": "engine-video-process-service-worker",
    #     "nodes_name": "engine-video-process-service-nodes",
    #     "vpsw": "engine-video-process-service-nodes",
    #     "status": "alone",
    #     "worker_name": "worker"
    # },
]

# override 中的 worker 名称
# override_worker_name = ["es_worker", "pach_worker", "crowd_oversea_worker",
#                "es_empty_worker", "es_pach_worker", "headshoulder_worker"]

SERVER_URL = "http://127.0.0.1:8000"
TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjg0MjAzMzk0LCJpc3MiOiJzb3Bob24ifQ.E0eNvOsYM-jiNeptz9J32Mp5Ad3REwZYd5WMWMh1Rmw"


class vps_info:
    def __init__(self):
        self.vps_service_info = vps_service_info

    # 1.获取每个vps服务对应的系统进程的pid
    def get_vps_pid(self):
        for name in vps_service_name:
            # get_name_cmd = "ps -ef | grep %s | grep -v grep" % name
            get_pid_cmd = "ps -ef | grep %s | grep -v grep | awk '{print $2}'" % name
            # pid = os.popen(get_pid_cmd).read().split()[0]
            res = subprocess.Popen(get_pid_cmd, shell=True, stdin=subprocess.PIPE, stdout=subprocess.PIPE,
                                   stderr=subprocess.PIPE)
            stdout, stderr = res.communicate()
            if stderr:
                print("Error : execute get pid command failed , err = ", stderr)
                sys.exit(1)
            else:
                pid_list = stdout.split("\n")
                if pid_list[-1] == "":
                    pid_list = pid_list[0:-1]

                for info in self.vps_service_info:
                    if info.get('name') == name:
                        info['pid'] = pid_list
        # print(self.vps_service_info)

    # 2.获取每个vps服务的quotas值
    def get_quotas(self):
        for service_info in self.vps_service_info:
            cmdstr = "license_client status | grep %s | awk '{print $4}' | awk -F} '{print $1}'" % service_info[
                "nodes_name"]
            res = subprocess.Popen(cmdstr, shell=True, stdin=subprocess.PIPE, stdout=subprocess.PIPE,
                                   stderr=subprocess.PIPE)
            stdout, stderr = res.communicate()
            if stderr:
                print("Error : execute get license quotas nodes failed , err = ", stderr)
                sys.exit(1)
            else:
                quotas = stdout.replace("\n", "")
                if quotas:
                    quotas_int = int(quotas) - 1  # quotas需要减1
                    service_info["quotas"] = quotas_int
                else:
                    service_info["quotas"] = ""  # 否则放一个空值
        print("\n\nfirst get quotas:-------------->")
        print(self.vps_service_info)


class check_occupy_status:
    def __init__(self):
        self.gpu_num = int(os.popen('nvidia-smi -L | wc -l').read().split()[0])  # 显卡总数

    # 1.获取空梯、物体的quotas
    def take_quotas(self, vps_service_info):
        self.vps_service_info = vps_service_info
        for service_info in self.vps_service_info:
            if service_info["name"] == "engine-video-es-pach-process-service-worker":
                if service_info.get("quotas"):
                    self.vepps_quotas = service_info["quotas"]
                else:
                    print("Info : not found es-pach quotas ")
                    self.vepps_quotas = ""
            elif service_info["name"] == "engine-video-es-empty-process-service-worker":
                if service_info.get("quotas"):
                    self.veeps_quotas = service_info["quotas"]
                else:
                    print("Info : not found es-empty quotas ")
                    self.veeps_quotas = ""

        if self.vepps_quotas == self.veeps_quotas == "":
            print("Info : Unauthorized es-pach-process quotas and es-empty-process quotas")
            print("Now exit , bye bye ...")
            sys.exit(0)
        print(self.veeps_quotas, self.vepps_quotas)

    # 2.拿到空梯/物体 quotas的最大值
    def get_max_quotas(self):
        # 值为空也可用判断
        if self.veeps_quotas >= self.vepps_quotas:
            self.max_quotas = self.veeps_quotas
            # print("veeps_quotas max")
        elif self.veeps_quotas < self.vepps_quotas:
            self.max_quotas = self.veps_quotas
            # print("vepps_quotas max")
        else:
            print("Error : compare failed")
            sys.exit(1)

    # 4.获取每个显卡上的vps的pid
    def nvidia_pid(self):
        self.nvidia_card_pid = []  # 记录每个显卡上的vps的pid
        for i in range(self.gpu_num):
            cmdstr = "nvidia-smi -i %d | grep 'video-process-service-worker' | awk '{print $3}'" % i
            res = subprocess.Popen(cmdstr, shell=True, stdin=subprocess.PIPE, stdout=subprocess.PIPE,
                                   stderr=subprocess.PIPE)
            stdout, stderr = res.communicate()
            if stderr:
                print("Error : execute nvidia-sm -i command failed , err = ", stderr)
                sys.exit(1)
            else:
                vps_pid = stdout.split("\n")
                if vps_pid[-1] == "":
                    vps_pid = vps_pid[0:-1]
                self.nvidia_card_pid.append({"nvidia_vps_pid": vps_pid, "id": i})
        # print(self.nvidia_card_pid)

    # 3.拿到非物体/空梯以外的显卡id,作为可用显卡id
    def get_avaliable_card(self):
        # 以下为老版本的功能代码
        self.gpu_num = int(os.popen('nvidia-smi -L | wc -l').read().split()[0])  # 显卡总数
        self.gpu_id = []
        for id in range(self.gpu_num):
            self.gpu_id.append(id)
        self.available_gpu_id = self.gpu_id[self.max_quotas:]
        print("self.gpu_id: ", self.gpu_id)
        print("available_gpu_id: ", self.available_gpu_id)

    # 5.将空梯/物体quotas值对应的显卡上的pid保存下来(当空梯/物体quotas最大值为3时，则将0、1、2这3张显卡上的pid保存下来，因为这几张卡上是不应该存在非空梯、非物体以外的vps服务)
    def es_pach_empty_nvidia_pid(self):
        self.es_pach_empty_nvidia_pid_list = []
        if int(self.max_quotas):
            for card_info in self.nvidia_card_pid:
                # 如果当前显卡id 小于最大的空梯/物体的quotas值，表示当前判断的显卡上不应该存在非空梯、非物体以外的vps服务。
                # 因为空梯/物体的pod是按照从0开始的顺序，依次固定显卡的。
                if card_info.get("id") < self.max_quotas:
                    for nvidia_pid in card_info.get("nvidia_vps_pid"):
                        self.es_pach_empty_nvidia_pid_list.append(nvidia_pid)
        else:
            print("Info : not authorized es-pach-process quotas and es-empty-process quotas")
            print("Now exit , bye bye ...")
            sys.exit(0)
        # print(self.es_pach_empty_nvidia_pid_list)

    # 6.循环显卡,检查空梯/物体显卡上是否有其他服务,有则更新vps服务
    def loop_nvidia_card(self, opera_vps_override):
        self.ops_override = opera_vps_override
        # self.es_pach_empty_nvidia_pid_list = ['8080', '8107', '8048', '8072', '8094']
        self.remove_list = self.es_pach_empty_nvidia_pid_list[:]
        for nvidia_pid in self.es_pach_empty_nvidia_pid_list:
            for service_info in self.vps_service_info:
                if service_info["name"] == "engine-video-es-pach-process-service-worker" or \
                        service_info["name"] == "engine-video-es-empty-process-service-worker":
                    if nvidia_pid in service_info.get("pid"):
                        # print("remove nvidia pid: [%s]\n" % nvidia_pid)
                        self.remove_list.remove(nvidia_pid)
        print("remove_list", self.remove_list)

        tag = False
        # 当有不该存在的pid时，需要进行重装vps服务
        if len(self.remove_list) != 0:
            print("开始更新vps服务...")
            self.ops_override.deploy_instance()
            tag = True
        if tag:
            time.sleep(1)
            print("已重新安装vps服务...")
        else:
            time.sleep(1)
            print("无需重新安装vps服务...")


# 从sophon中获取vps 的 override文件，并修改vps的override文件，最后更新vps服务
class opera_vps_override:
    def __init__(self, vps_service_info):
        self.vps_service_info = vps_service_info
        self.url = "%s/v1/instances" % SERVER_URL
        self.headers = {
            "Content-Type": "",
            "Authorization": "Bearer %s" % TOKEN
        }

    # 1.获取所有实例的信息
    def get_instances(self):
        resp = requests.get(url=self.url, headers=self.headers)
        self.instances = resp.json().get("instances")

    # 3.更新vps的override的replicas为quotas，固定非物体/空梯服务的显卡id
    def modify_vps_override(self, available_gpu_id):
        self.available_gpu_id = available_gpu_id
        # 保存vps instance
        helm_name = "engine-video-process-service-nebula"
        for instance in self.instances:
            if helm_name == instance.get("name"):
                self.vps_instance = instance  # vps 实例的所有信息

        print("\n\nvps instance -------->", type(self.vps_instance))
        print(self.vps_instance)
        print("\n\n")

        self.vps_config_unicode = self.vps_instance.get("config")  # vps 的 override 文件

        self.vps_config = yaml.load(self.vps_config_unicode, Loader=yaml.FullLoader)  # 转成字典
        print("vps vps_config config -------->", type(self.vps_config))
        print(self.vps_config)

        # 粗略的估算config文件长度是否完整，主要用于提醒
        if len(self.vps_config_unicode) < 20:  # 当override的配置文件长度小于20，可用肯定配置肯定不完整。(这里只是粗略的判断)
            print("Error : vps override is incomplete")
            sys.exit(1)
        elif len(self.vps_config_unicode) < 100:  # 当override的配置文件长度小于20，可用肯定配置不够完整。(这里只是粗略的判断)
            print("Warning : vps override is not enough complete")
        elif len(self.vps_config_unicode) > 300:
            print("Info : vps override is complete , the vps override file length is [%s]" % len(
                self.vps_config_unicode))

        print("available_gpu_id 2-->", self.available_gpu_id)
        self.remove_availabel_gpu_id = self.available_gpu_id[:]
        # 将获取的quotas数值放到override的各worker的replicas中
        for service_info in self.vps_service_info:
            # print("service info worker name: ", service_info.get("worker_name"))
            worker_name = service_info.get("worker_name")
            quotas = service_info.get("quotas")
            if str(quotas).isdigit():  # 如果quotas是数字，表示有值，且已经license授权
                if self.vps_config.get(worker_name):
                    # 修改replicas为quotas值
                    self.vps_config.get(worker_name)["replicas"] = quotas
                    # 如果 worker_name != 空梯/物体，则固定显卡id， resources的limits中gpu设置为null-->问题:如果crowd起的2个，那2个crowd就会在一张卡上了
                    # if worker_name not in ["es_empty_worker", "es_pach_worker"]:
                    #     if len(self.remove_availabel_gpu_id) > 0:
                    #         self.vps_config.get(worker_name)["env"] = [{"name": "NVIDIA_VISIBLE_DEVICES", "value": str(self.remove_availabel_gpu_id[0])}]
                    #         self.remove_availabel_gpu_id.pop(0)
                    #         # 'resources': {'limits': {'nvidia.com/gpu': None}}
                    #         self.vps_config.get(worker_name)["resources"] = {'limits': {'nvidia.com/gpu': None}}
                else:
                    print("Warning: not found [%s] worker name" % worker_name)
                    self.vps_config[worker_name] = {"replicas": quotas}

                    # if worker_name not in ["es_empty_worker", "es_pach_worker"]:
                    #     if len(self.remove_availabel_gpu_id) > 0:
                    #         self.vps_config.get(worker_name)["env"] = [{"name": "NVIDIA_VISIBLE_DEVICES", "value": str(self.remove_availabel_gpu_id[0])}]
                    #         self.remove_availabel_gpu_id.pop(0)
                    #         self.vps_config.get(worker_name)["resources"] = {'limits': {'nvidia.com/gpu': None}}
            else:  # 不为数字说明为空""
                if self.vps_config.get(worker_name):
                    self.vps_config.get(worker_name)["replicas"] = 0
                else:
                    self.vps_config[worker_name] = {"replicas": 0}

        print("剩余可以使用的显卡")
        print(self.remove_availabel_gpu_id)
        print("\n\n")

        self.vps_instance_str = yaml.dump(self.vps_config)
        self.vps_instance["config"] = self.vps_instance_str
        # print("\n\nself.vps_config updated quotas--------ok:")
        # print(self.vps_config)
        print("\n\nself.vps_instance updated quotas--------ok:")
        print(self.vps_instance)

    # 更新非物体/空梯以外服务的配置
    def modify_others_vps(self):
        print("更新非物体/空梯以外服务的配置")
        # 解决非空梯/物体的服务固定到可用显卡
        #   a. 拿到可用的显卡，给于quotas>=1，且不为空梯、物体的服务，固定显卡。
        #   发生条件: 当:物体/空梯显卡上有其他服务则进行以下操作:
        #      a0. 通过传入的服务名，对override中的显卡字段进行修改 -- 可用不用此功能
        #           available_gpu_id
        #      a1. 找到可用显卡（先找到非物体/空梯以外的显卡id，这些就是作为可用显卡id。就算这些卡上有其他服务，但是万一有两个其他服务固定到一张卡上，那就出问题了，所以还不如这里再将显卡id赋值给其他服务上）
        #      a2. 找到已授权的非物体/空梯服务(quotas >= 1)
        #      a3. 固定非物体/空梯服务的显卡id(字典操作，判断如果没有字段，则增加字段)
        #      a4. 将当前服务的quotas值放到replicas上(解决下面b提出的问题，此处modify_vps_override函数的for循环部分已解决了)
        #      a5. 更新服务

    # 部署实例
    # def deploy_instance(self, instance):
    def deploy_instance(self):
        instance = self.vps_instance
        body = {
            "instance": instance,
            "recreate_pods": True
        }
        resp = requests.post(url=self.url, headers=self.headers, data=json.dumps(body))
        if resp.status_code != 200:
            print("\n\n")
            print("Error : deploy instance failed ...", resp.text)
        else:
            print("\n\n")
            print("Info : deploy instance success ...", resp.text)


if __name__ == '__main__':
    # 1. vps pid记录, license中vps服务的quotas记录
    get_vps_info = vps_info()
    get_vps_info.get_vps_pid()  # 获取每个vps服务对应的系统进程的pid
    get_vps_info.get_quotas()  # 获取每个vps服务的quotas值

    # 2. 显卡pid对应服务检查异常
    vps_ops = check_occupy_status()
    vps_ops.take_quotas(get_vps_info.vps_service_info)  # 获取空梯、物体的quotas
    vps_ops.get_max_quotas()  # 拿到空梯/物体 quotas的最大值
    vps_ops.nvidia_pid()  # 获取每个显卡上的vps的pid
    vps_ops.get_avaliable_card()  # 拿到非物体/空梯以外的显卡id,作为可用显卡id --> 这里要改写1
    vps_ops.es_pach_empty_nvidia_pid()  # 拿到空梯及物体显卡id上的所有pid  --> 删除此功能

    # 3. instance及override获取及更新
    ops_override = opera_vps_override(get_vps_info.vps_service_info)
    ops_override.get_instances()  # 获取所有实例的信息
    ops_override.modify_vps_override(vps_ops.available_gpu_id)  # 更新vps的override的replicas为quotas(检测replicas是否等于quotas，不等于则更新override，并更新服务) --> 这里要改写2
    # ops_override.deploy_instance()  # 用于更新vps服务的调用函数

    # 4.更新服务
    vps_ops.loop_nvidia_card(ops_override)  # 循环显卡,检查空梯/物体显卡上是否有其他服务,有则更新vps服务  --> 删除此功能
