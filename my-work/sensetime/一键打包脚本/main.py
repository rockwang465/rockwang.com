#!/usr/bin/env python
# -*- coding: utf-8 -*-

from get_version import *
from connect_10 import *

standard_env_ip = '10.5.6.66'
standard_env_username = 'root'
standard_env_passwd = 'Nebula123$%^'
port = 22
env_10_ip = '10.5.6.10'
env_10_username = 'root'
env_10_passwd = 'Sensetime12345'
repository_domain = 'www.sensetime.com'  # docker images domain name(以后docker镜像可能会定义一个域名开头的，而非IP开头的)
# docker_repo_port = '5000'  # 在标准环境中的docker仓库端口

json_file = './versions.json'
remote_json_file = '/tmp/pack/versions.json'
script_file = './pack_charts_images.py'
remote_script_file = '/tmp/pack/pack_charts_images.py'

if __name__ == "__main__":
    # 1. get charts and images version to versions.json file
    charts = get_charts_version()
    charts_info = charts.get_helm_charts(standard_env_ip, port, standard_env_username, standard_env_passwd)
    charts.convert_helm_charts()

    images = get_images_version()
    images.get_kube_config(standard_env_ip, port, standard_env_username, standard_env_passwd)
    images.get_kube_images(repository_domain)
    images.convert_json()

    merge_data = merge_charts_images(charts.new_charts, images.new_images)
    merge_data.merge()
    # merge_data.all_data 是最终的json数据

    # 2. to standard evn pull charts and images packages
    pack = scp_files()
    pack.ssh_scp_put(env_10_ip, port, env_10_username, env_10_passwd, json_file, remote_json_file)
    pack.ssh_scp_put(env_10_ip, port, env_10_username, env_10_passwd, script_file, remote_script_file)
    pack.exec_10_script(env_10_ip, port, env_10_username, env_10_passwd, remote_script_file)