#!/usr/bin/env python
# -*- coding: utf-8 -*-

import os
import sys
import yaml


class create_dir:
    def create_path(self, base_pack_path):
        # print(base_pack_path)
        if not os.path.isdir(base_pack_path):
            os.mkdir(base_pack_path)
            print("Info : [%s] directory created successful" % base_pack_path)
        else:
            print("Error : [%s] path already exists " % base_pack_path)
            sys.exit(1)


# 拿到基础文件到 base 目录中
class take_base_files:
    # source_path 为 /data/packages/sensenebula/src/
    # base_pack_path 为 /data/packages/sensenebula/releases/SenseNebula-G-v1.2.0+20190812103917/base,
    # base_dir 为  [{'name': 'license-ca', 'type': 'dir'}...]
    def copy_dir(self, source_path, base_pack_path, base_dir):
        for k in base_dir:
            desc_name = source_path + k["name"]
            if k["type"] == "dir":
                res = os.system("cp -r %s %s" % (desc_name, base_pack_path))
            elif k["type"] == "file":
                res = os.system("cp %s %s" % (desc_name, base_pack_path))
            else:
                print("Error : [%s] file [%s] type error" % (k["name"], k["type"]))
                sys.exit(1)
            if res != 0:
                print("Error : [%s] file copy failed" % k)
                sys.exit(1)
            print("Info : [%s] directory copy successful" % desc_name)


# 拿到 infra-ansible 脚本
class take_ansbile_file:
    # git 拉取 infra-ansible 代码
    def git_clone(self, base_pack_path, ansible_git_addr, git_branch):
        os.chdir(base_pack_path)
        res = os.system("git clone -b %s %s -q" % (git_branch, ansible_git_addr))
        if res != 0:
            print("Error : Failure to git clone")

    # 打包ansible代码包
    def pack_ansible(self, ansible_dir_name, base_pack_path):
        res1 = os.system("tar -zcf %s.tgz %s" % (ansible_dir_name, ansible_dir_name))
        if res1 != 0:
            print("Error : Failure to pack %s")
            sys.exit(1)
        # 判断删除的文件路径不为/ 目录
        if base_pack_path == "/":
            print("Error : [%s] path error" % base_pack_path)
            sys.exit(1)
        else:
            res2 = os.system("rm -rf %s/%s" % (base_pack_path, ansible_dir_name))


# 整个release目录打包
class pack_release:
    # 拷贝versions.json文件到 SenseNebula-G-xxx 中
    def copy_versions(self, work_dir, json_file, current_release_path):
        res = os.system("cp %s/%s %s" % (work_dir, json_file, current_release_path))
        if res != 0:
            print("Error : Copy operator [cp %s/%s %s] failure" % (work_dir, json_file, current_release_path))

    # 转换versions.json为 images.yaml 和 versions.yaml
    def dump_images_charts(self, json_file, current_release_path):
        os.chdir(current_release_path)
        f = open(json_file, 'r')
        data = yaml.full_load(f.read())
        images_dic = data.get('images')
        charts_dic = data.get('charts')

        images_file = open('./images.yaml', 'w')
        yaml.dump(images_dic, images_file)
        charts_file = open('./versions.yaml', 'w')
        yaml.dump(charts_dic, charts_file)
        f.close()

    # 打包 SenseNebula-G-xxx 为 tgz 包
    def pack_all(self, release_path, release_package_name):
        os.chdir(release_path)
        res1 = os.system("tar -zcf %s.tgz %s" % (release_package_name, release_package_name))
        if res1 != 0:
            print("Error : Failure to pack %s.tgz" % release_package_name)
            sys.exit(1)
        res2 = os.system("md5sum %s.tgz > %s.tgz.md5" % (release_package_name, release_package_name))
        if res2 != 0:
            print("Error : Failure to make md5")
