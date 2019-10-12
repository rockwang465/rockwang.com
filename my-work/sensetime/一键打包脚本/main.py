#!/usr/bin/env python
# -*- coding: utf-8 -*-
import argparse
from get_version import *
from connect_10 import *
import datetime
import time

port = 22
env_10_ip = '10.5.6.10'
env_10_username = 'root'
env_10_passwd = 'Sensetime12345'
repository_domain = 'registry.sensenebula.io:5000'  # docker images domain name(以后docker镜像可能会定义一个域名开头的，而非IP开头的)
# docker_repo_port = '5000'  # 在标准环境中的docker仓库端口

pack_dir = '/data/packages/sensenebula/pack/'
json_file = 'versions.json'
script_main_file = 'pack_main.py'
script_server_file = 'pack_charts_images.py'
script_base_file = 'pack_base.py'


def parse_args():
    parser = argparse.ArgumentParser(description='Get charts and images version, and pack it')
    parser.add_argument("env_ip", help="standard environment ip address")
    parser.add_argument("--env_username", help="standard environment username", default="root")
    parser.add_argument("--env_passwd", help="standard environment password", default="Nebula123$%^")
    parser.add_argument("--env_port", help="standard environment port", default="22")
    parser.add_argument("--version", help="pack release version, for example: v1.2.0", default="v1.2.0")
    parser.add_argument("--infra_branch", help="git clone infra-ansible branch, for example: dev, qudao, SPG",
                        default="dev")
    parser.add_argument("--tools_branch", help="git clone tools branch, for example: dev, SPG", default="dev")
    return parser.parse_args()


class define_path():
    def path_name(self, args):
        self.curr_time = datetime.datetime.now().strftime('%Y%m%d%H%M%S')
        self.tail_ip = args.env_ip.split('.')[-1]  # ip的最后一位，如63
        self.dir_name = 'pack' + self.tail_ip + '_' + args.version + '_' + self.curr_time  # pack63_v1.2.0_20190929162510
        self.work_dir = pack_dir + self.dir_name
        self.remote_json_file = self.work_dir + '/' + json_file
        self.remote_main_file = self.work_dir + '/' + script_main_file
        self.remote_server_file = self.work_dir + '/' + script_server_file
        self.remote_base_file = self.work_dir + '/' + script_base_file


if __name__ == "__main__":
    # 0. 传参操作
    args = parse_args()

    # 1. 目录定义
    path = define_path()
    path.path_name(args)

    # # 2. 获取 charts 和 images 版本信息到 versions.json 文件中
    charts = get_charts_version()  # Rock测试: 第一次执行生成versions.json；第二次修改versions.json，并把此处第2步全部注释。
    charts.get_helm_charts(args)
    charts.convert_helm_charts()
    charts.convert_common_charts()

    images = get_images_version()
    images.get_kube_config(args)
    images.get_kube_images(repository_domain)
    images.convert_json()

    merge_data = merge_charts_images()
    merge_data.merge_to_versions(charts.new_charts, images.new_images)

    # 3. 发送文件到 10.5.6.10 服务器上
    scp = scp_files()
    scp.ssh_scp_put(env_10_ip, port, env_10_username, env_10_passwd, json_file, path.remote_json_file)  # 临时注释
    scp.ssh_scp_put(env_10_ip, port, env_10_username, env_10_passwd, script_main_file, path.remote_main_file)
    scp.ssh_scp_put(env_10_ip, port, env_10_username, env_10_passwd, script_server_file, path.remote_server_file)
    scp.ssh_scp_put(env_10_ip, port, env_10_username, env_10_passwd, script_base_file, path.remote_base_file)
    scp.exec_10_script(env_10_ip, port, env_10_username, env_10_passwd, path.remote_main_file, args, path.work_dir)

    print("Info : Pack is over")
