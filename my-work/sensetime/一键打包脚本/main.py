#!/usr/bin/env python
# -*- coding: utf-8 -*-
import argparse
from get_version import *
from connect_10 import *


port = 22
env_10_ip = '10.5.6.10'
env_10_username = 'root'
env_10_passwd = 'Sensetime12345'
repository_domain = 'www.sensetime.com'  # docker images domain name(以后docker镜像可能会定义一个域名开头的，而非IP开头的)
# docker_repo_port = '5000'  # 在标准环境中的docker仓库端口

json_file = 'versions.json'
remote_file = '/data/packages/sensenebula/pack/versions.json'
remote_json_file = '/data/packages/sensenebula/pack/versions.json'
script_main_file = 'pack_main.py'
remote_main_file = '/data/packages/sensenebula/pack/pack_main.py'
script_server_file = 'pack_charts_images.py'
remote_server_file = '/data/packages/sensenebula/pack/pack_charts_images.py'
script_base_file = 'pack_base.py'
remote_base_file = '/data/packages/sensenebula/pack/pack_base.py'


def parse_args():
    parser = argparse.ArgumentParser(description='Get charts and images version, and pack it')
    parser.add_argument("env_ip", help="standard environment ip address")
    parser.add_argument("--env_username", help="standard environment username", default="root")
    parser.add_argument("--env_passwd", help="standard environment password", default="Nebula123$%^")
    parser.add_argument("--env_port", help="standard environment port", default="22")
    parser.add_argument("--version", help="pack release version, for example: v1.2.0", default="v1.2.0")
    return parser.parse_args()


if __name__ == "__main__":
    # 0. 传参操作
    args = parse_args()

    # 1. 获取 charts 和 images 版本信息到 versions.json 文件中
    charts = get_charts_version()
    charts.get_helm_charts(args)
    charts.convert_helm_charts()

    images = get_images_version()
    images.get_kube_config(args)
    images.get_kube_images(repository_domain)
    images.convert_json()

    merge_data = merge_charts_images()
    merge_data.merge_to_versions(charts.new_charts, images.new_images)

    # 2. 发送文件到 10.5.6.10 服务器上
    scp = scp_files()
    scp.ssh_scp_put(env_10_ip, port, env_10_username, env_10_passwd, json_file, remote_json_file)  # 临时注释
    scp.ssh_scp_put(env_10_ip, port, env_10_username, env_10_passwd, script_main_file, remote_main_file)
    scp.ssh_scp_put(env_10_ip, port, env_10_username, env_10_passwd, script_server_file, remote_server_file)
    scp.ssh_scp_put(env_10_ip, port, env_10_username, env_10_passwd, script_base_file, remote_base_file)
    scp.exec_10_script(env_10_ip, port, env_10_username, env_10_passwd, remote_main_file, args)
