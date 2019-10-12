#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
from pack_charts_images import *
from pack_base import *

base_source_path = '/data/packages/sensenebula/src/'
base_dir = [{'name': 'license-ca', 'type': 'dir'},
            {'name': 'tools', 'type': 'dir'},
            {'name': 'yum-data', 'type': 'dir'},
            {'name': 'init-new.sh', 'type': 'file'}]
git_infra_addr = 'git@gitlab.sz.sensetime.com:galaxias/infra-ansible.git'
git_tools_addr = 'git@gitlab.sz.sensetime.com:galaxias/tools.git'
ansible_dir_name = 'infra-ansible'


def parse_args():
    parser = argparse.ArgumentParser(description='Get charts and images version, and pack it')
    parser.add_argument("env_ip", help="standard environment ip address")
    parser.add_argument("--version", help="pack release version, for example: v1.2.0", default="v1.2.0")
    parser.add_argument("--infra_branch", help="git clone infra-ansible branch, for example: dev, qudao, SPG", default="dev")
    parser.add_argument("--tools_branch", help="git clone tools branch, for example: dev, SPG", default="dev")
    parser.add_argument("--work_dir", help="pack work directory")
    return parser.parse_args()


# 0. 传参
args = parse_args()

# 1. 打包 charts 和images
dire = define_dir()
dire.path_name(args)
dire.create_dir()

registry = run_registry()
registry.run_docker_registry(dire)

version = get_version()
version.get_version_data(args)

p_img = pack_images()
p_img.docker_operator(version.images_version, dire, registry)
registry.del_docker_registry()

p_helm = pack_charts()
p_helm.helm_operator(version.charts_version, dire)

# 2. 打包 base
c = create_dir()
c.create_path(dire.base_pack_path)

base_files = take_base_files()
base_files.copy_dir(base_source_path, dire.base_pack_path, base_dir)

take_ansbile = take_ansbile_file()
take_ansbile.git_clone_infra(dire.current_release_path, git_infra_addr, args.infra_branch)
take_ansbile.git_clone_tools(dire.current_release_path, git_tools_addr, args.tools_branch)
# take_ansbile.pack_ansible(args, dire, ansible_dir_name, dire.base_pack_path) # 暂时不用的功能，打包ansible脚本成tgz格式

# 3. 打包 release
p_release = pack_release()
p_release.copy_versions(args.work_dir, json_file, dire.current_release_path)
p_release.dump_images_charts(json_file, dire.current_release_path)
p_release.pack_all(release_path, dire.release_package_name)
