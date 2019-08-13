#!/usr/bin/env python
# -*- coding: utf-8 -*-

from pack_charts_images import *
from pack_base import *

base_source_path = '/data/packages/sensenebula/src/'
base_dir = [{'name': 'license-ca', 'type': 'dir'},
            {'name': 'tools', 'type': 'dir'},
            {'name': 'yum-data', 'type': 'dir'}]
ansible_git_addr = 'git@gitlab.sz.sensetime.com:galaxias/infra-ansible.git'
git_branch = 'dev'
ansible_dir_name = 'infra-ansible'

# if __name__ == 'main':
# 1. pack_charts_images
dire = define_dir()
dire.path_name()
dire.create_dir()

registry = run_registry()
registry.run_docker_registry(dire)

version = get_version()
version.get_version_data()

p_img = pack_images()
p_img.docker_operator(version.images_version, dire, registry)
# p_img.docker_operator(version.k8s_images_version)
registry.del_docker_registry()

p_helm = pack_charts()
p_helm.helm_operator(version.charts_version, dire)

# 2. pack_base
c = create_dir()
c.create_path(dire.base_pack_path)

base_files = take_base_files()
base_files.copy_dir(base_source_path, dire.base_pack_path, base_dir)

take_ansbile = take_ansbile_file()
take_ansbile.git_clone(dire.base_pack_path, ansible_git_addr, git_branch)
take_ansbile.pack_ansible(ansible_dir_name, dire.base_pack_path)

# 3. pack_release
p_release = pack_release()
p_release.copy_versions(work_dir, json_file, dire.current_release_path)
p_release.pack_all(release_path, dire.release_package_name)
