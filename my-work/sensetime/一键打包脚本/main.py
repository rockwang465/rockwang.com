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

json_file = 'versions.json'
remote_file = '/data/packages/sensenebula/pack/versions.json'
remote_json_file = '/data/packages/sensenebula/pack/versions.json'
script_main_file = 'pack_main.py'
remote_main_file = '/data/packages/sensenebula/pack/pack_main.py'
script_server_file = 'pack_charts_images.py'
remote_server_file = '/data/packages/sensenebula/pack/pack_charts_images.py'
script_base_file = 'pack_base.py'
remote_base_file = '/data/packages/sensenebula/pack/pack_base.py'

# versions_pack_file = './versions_pack.json'
# remote_pack_file = '/data/packages/sensenebula/pack/versions_pack.json'

if __name__ == "__main__":
    # 1. get charts and images version to versions.json file
    charts = get_charts_version()
    charts_info = charts.get_helm_charts(standard_env_ip, port, standard_env_username, standard_env_passwd)
    charts.convert_helm_charts()

    images = get_images_version()
    images.get_kube_config(standard_env_ip, port, standard_env_username, standard_env_passwd)
    images.get_kube_images(repository_domain)
    images.convert_json()

    merge_data = merge_charts_images()
    merge_data.merge_to_versions(charts.new_charts, images.new_images)
    # merge_data.all_version 是最终的json数据
    # merge_data.merge_to_pack(charts.new_charts, images.new_images)

    # 2. send files to 10.5.6.10
    scp = scp_files()
    scp.ssh_scp_put(env_10_ip, port, env_10_username, env_10_passwd, json_file, remote_json_file)  # 临时注释
    scp.ssh_scp_put(env_10_ip, port, env_10_username, env_10_passwd, script_main_file, remote_main_file)
    scp.ssh_scp_put(env_10_ip, port, env_10_username, env_10_passwd, script_server_file, remote_server_file)
    scp.ssh_scp_put(env_10_ip, port, env_10_username, env_10_passwd, script_base_file, remote_base_file)
    scp.exec_10_script(env_10_ip, port, env_10_username, env_10_passwd, remote_main_file)

# 剩余工作及待优化的点:
# 1 剩余工作
# 1.1 打包后 如何插入到数据库

# 2. 优化方面
# 2.1 多次命名 versions.json文件变量，可以使用传参或者jinja2
# 2.2 scp文件的remote file改为路径方式
