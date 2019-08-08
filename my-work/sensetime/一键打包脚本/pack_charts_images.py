#!/usr/bin/env python
# -*- coding: utf-8 -*-
import os
import sys
import datetime
import json
import time

# 大概步骤描述
#
# 1.1 使用docker起一个docker仓库，挂载本地目录
# 1.2 在docker仓库里，pull,tag,push到仓库里
# 1.3 打包docker仓库目录
# 2.1 拉取charts包
# 3   打包docker目录、charts目录、基础包

standard_env_ip = '10.5.6.66'  # 后期jinja2优化
mapping_host_port = 8001  # 主机访问registry端口
mapping_docker_port = 5000  # 容器内部的registry端口
release_name = 'SenseNebula-G'
release_version = 'v1.2.0'
release_path = '/data/packages/sensenebula/releases/'
mount_registry_path = '/var/lib/registry'  # 默认registry的容器目录
registry_image = '10.5.6.10/docker.io/registry:2'

work_dir = '/data/packages/sensenebula/pack'  # 后期jinja2优化
versions_pack_file = './versions_pack.json'  # 后期jinja2优化
env_10_ip = '10.5.6.10'  # 后期jinja2优化
env_10_charts_port = '8080'


class define_dir:
    def __init__(self):
        self.mapping_host_port = mapping_host_port
        self.tail_ip = standard_env_ip.split(".")[-1]
        self.registry_name = "registry" + self.tail_ip  # docker容器命名，在结尾加上标准环境ip的尾部(如10.5.6.66则为 registry66)
        self.now_time = datetime.datetime.now().strftime('%Y%m%d%H%M%S')
        self.release_package_name = release_name + "-" + release_version + "+" + self.now_time
        self.mount_10_path = release_path + self.release_package_name + "/" + "images"  # docker registry挂载在10.5.6.10上的路径
        self.charts_pack_path = release_path + self.release_package_name + "/" + "charts"  # 下载charts包路径

    def create_dir(self):
        if os.path.isdir(self.mount_10_path):
            print("Error : Already exists [%s] directory" % self.mount_10_path)
            sys.exit(1)
        else:
            os.makedirs(self.mount_10_path)
            print("Info : Successful create [%s] of pack directory" % self.mount_10_path)
            print("\n\n")


class run_registry(define_dir):
    # run a docker registry
    def run_docker_registry(self):
        print("Info : Start running docker registry")
        # 判断 mapping_host_port 对应的端口是否已存在
        res_port = int(os.popen("ss -lntup |grep %s |wc -l" % self.mapping_host_port).read().replace("\n", ""))

        # 判断是否存在同名字的容器名字
        res_name = int(os.popen(
            "docker ps | grep %s | grep %s | wc -l" % (self.registry_name, self.mapping_host_port)).read().replace("\n",
                                                                                                                   ""))

        if res_port == 0 and res_name == 0:
            res = os.popen("docker run -d -p %s:%s --restart=always --name %s -v %s:%s %s" % (
                self.mapping_host_port, mapping_docker_port, self.registry_name, self.mount_10_path,
                mount_registry_path, registry_image))
            self.cont_id = res.read()
            if not self.cont_id:
                print("Error : Failure to run docker registry")
                sys.exit(1)
        else:
            print("Error : Already exists [%s] port or [%s] container name, Please check " % (
                self.mapping_host_port, self.registry_name))
            sys.exit(1)
        print("Info : Successful running registry")
        print("\n\n")
        # 示例 : os.system("docker run -d -p 8001:5000 --restart=always --name registry -v /data/packages/sensenebula/releases/SenseNebula_G-v1.2.0+20190807104744/images:/var/lib/registry 10.5.6.10/docker.io/registry:2")

    def del_docker_registry(self):
        res = os.system("docker rm -f %s" % self.cont_id)
        if res == 0:
            print("Info : Successful to stop docker registry ")
        else:
            print("Error : Failure to stop docker registry , id : [%s]" % self.cont_id)


class get_version:
    # 拿到images版本信息
    def get_version_data(self):
        os.chdir(work_dir)
        with open(versions_pack_file, 'r') as fp:
            data = json.loads(fp.read())
        if data:
            pass
        else:
            print("Error : get %s file data failure" % versions_pack_file)
            sys.exit(1)
        self.images_version = data.get("images")
        self.k8s_images_version = data.get("k8s_images")
        self.charts_version = data.get("charts")


class pack_images:
    def docker_operator(self, images_version):
        print("Info : Start packing images")
        # {u'tag': u'5.5.4', u'repository': u'elasticsearch/curator'},
        for i in images_version:
            print(i.get("tag"), i.get("repository"))
            tag = i.get("tag")
            repo = i.get("repository")

            pull_image = "%s/%s:%s" % (env_10_ip, repo, tag)
            # print("Info : docker pull %s" % pull_image)
            res1 = os.system("docker pull %s" % pull_image)
            if res1 != 0:
                print("Error : docker pull [%s] failure" % pull_image)
                sys.exit(1)

            tag_push_image = "%s:%s/%s:%s" % (env_10_ip, mapping_host_port, repo, tag)
            # print("Info : docker tag %s %s" % (pull_image, tag_push_image))
            res2 = os.system("docker tag %s %s" % (pull_image, tag_push_image))
            if res2 != 0:
                print("Error : docker tag [%s] [%s] failure" % (pull_image, tag_push_image))
                sys.exit(1)

            # print("Info : docker push %s" % tag_push_image)
            res3 = os.system("docker push %s" % tag_push_image)
            if res3 != 0:
                print("Error : docker push [%s] failure" % tag_push_image)
                sys.exit(1)
            time.sleep(1)
        print("Info : Successful pack to [%s] directory" % registry.release_package_name)
        print("\n\n")


class pack_charts(define_dir):
    def helm_operator(self, charts_version):
        print("Info : Start packing charts")
        if os.path.isdir(self.charts_pack_path):
            print("Error : Already exists [%s] directory" % self.charts_pack_path)
            sys.exit(1)
        else:
            os.makedirs(self.charts_pack_path)
            print("Info : Successful create [%s] of pack directory" % self.charts_pack_path)
            print("\n\n")

        os.chdir(self.charts_pack_path)
        for i in charts_version:
            c_name = i.get("name")
            c_version = i.get("version")
            c_namespace = i.get("namespace")
            # print(c_name, c_version, c_namespace)
            # helm fetch http://10.5.6.10:8080/charts/elasticsearch-curator-5.5.4-master-ac9bae1.tgz
            complete_server_name = c_name + "-" + c_version
            res_fetch = os.system(
                "helm fetch http://%s:%s/charts/%s.tgz" % (env_10_ip, env_10_charts_port, complete_server_name))
            res_md5sum = os.system("md5sum %s.tgz > %s.tgz.md5" % (complete_server_name, complete_server_name))
            if res_fetch != 0:
                print("Error : Failure to [helm fetch http://%s:%s/charts/%s.tgz]" % (
                    env_10_ip, env_10_charts_port, complete_server_name))
                sys.exit(1)
            elif res_md5sum != 0:
                print("Error : Failure to [md5sum %s.tgz > %s.tgz.md5]" % (complete_server_name, complete_server_name))
                sys.exit(1)
        print("Info : Successful pack to [%s] directory" % self.charts_pack_path)
        print("\n\n")


# if __name__ == 'main':
dir = define_dir()
dir.create_dir()

registry = run_registry()
registry.run_docker_registry()

version = get_version()
version.get_version_data()

p_img = pack_images()
p_img.docker_operator(version.images_version)
p_img.docker_operator(version.k8s_images_version)
registry.del_docker_registry()

p_helm = pack_charts()
p_helm.helm_operator(version.charts_version)
