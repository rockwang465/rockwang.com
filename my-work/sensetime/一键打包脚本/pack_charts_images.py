#!/usr/bin/env python
# -*- coding: utf-8 -*-
import os
import sys
import datetime
import json
import time

# standard_env_ip = '10.5.6.66'  # 后期jinja2优化
mapping_host_port = 8001  # 主机访问registry端口
mapping_docker_port = 5000  # 容器内部的registry端口
release_name = 'SenseNebula-G'
# release_version = 'v1.2.0'
release_path = '/data/packages/sensenebula/releases/'
mount_registry_path = '/var/lib/registry'  # 默认registry的容器目录
registry_image = '10.5.6.10/docker.io/registry:2'

work_dir = '/data/packages/sensenebula/pack'  # 后期jinja2优化
json_file = 'versions.json'  # 后期jinja2优化
env_10_ip = '10.5.6.10'  # 后期jinja2优化
env_10_charts_port = '8080'


# local_ip = '10.5.6.92'


class define_dir:
    def path_name(self, args):
        # self.mapping_host_port = mapping_host_port
        self.tail_ip = args.env_ip.split(".")[-1]
        self.registry_name = "registry" + self.tail_ip  # docker容器命名，在结尾加上标准环境ip的尾部(如10.5.6.66则为 registry66)
        self.now_time = datetime.datetime.now().strftime('%Y%m%d%H%M%S')
        # self.release_package_name = release_name + "-" + release_version + "+" + self.now_time
        self.release_package_name = release_name + "-" + args.version + "+" + self.now_time
        self.current_release_path = release_path + self.release_package_name
        self.mount_10_path = release_path + self.release_package_name + "/" + "images"  # docker registry挂载在10.5.6.10上的路径
        self.charts_pack_path = release_path + self.release_package_name + "/" + "charts"  # 下载charts包路径
        self.base_pack_path = release_path + self.release_package_name + "/" + "base"  # 基础base包路径

    def create_dir(self):
        if os.path.isdir(self.mount_10_path):
            print("Error : [%s] directory already exists " % self.mount_10_path)
            sys.exit(1)
        else:
            os.makedirs(self.mount_10_path)
            # os.makedirs(self.charts_pack_path)
            print("Info : [%s] created successful" % self.mount_10_path)
            print("\n\n")


class run_registry:
    # 运行 docker registry
    def run_docker_registry(self, dire):
        print("Info : Start running docker registry")
        # 判断 mapping_host_port 对应的端口是否已存在
        res_port = int(os.popen("ss -lntup |grep %s |wc -l" % mapping_host_port).read().replace("\n", ""))

        # 判断是否存在同名字的容器名字
        res_name = int(os.popen(
            "docker ps | grep %s | grep %s | wc -l" % (dire.registry_name, mapping_host_port)).read().replace("\n", ""))
        if res_port == 0 and res_name == 0:
            # 示例 : os.popen("docker run -d -p 8001:5000 --restart=always --name registry -v /data/packages/sensenebula/releases/SenseNebula_G-v1.2.0+20190807104744/images:/var/lib/registry 10.5.6.10/docker.io/registry:2")
            res = os.popen("docker run -d -p %s:%s --restart=always --name %s -v %s:%s %s" % (
                mapping_host_port, mapping_docker_port, dire.registry_name, dire.mount_10_path,
                mount_registry_path, registry_image))
            self.cont_id = res.read()  # 拿到容器的id
            if not self.cont_id:
                print("Error : Failure to run docker registry")
                sys.exit(1)
        else:
            print("Error : Already exists [%s] port or [%s] container name, Please check " % (
                mapping_host_port, dire.registry_name))
            sys.exit(1)
        print("Info : Successful running registry")
        print("\n\n")

    # 删除容器
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
        with open(json_file, 'r') as fp:
            data = json.loads(fp.read())
        if data:
            pass
        else:
            print("Error : get %s file data failure" % json_file)
            sys.exit(1)
        self.images_version = data.get("images")
        # self.k8s_images_version = data.get("k8s_images")
        self.charts_version = data.get("charts")


class pack_images:
    # pull tag push 操作
    def docker_operator(self, images_version, dire, registry):
        print("Info : Start packing images")
        # {u'tag': u'5.5.4', u'repository': u'elasticsearch/curator'},
        for i in images_version:
            tag = i.get("tag")
            repo = i.get("repository")

            pull_image = "%s/%s:%s" % (env_10_ip, repo, tag)
            res1 = os.system("docker pull %s >/dev/null" % pull_image)
            time.sleep(1)
            if res1 != 0:
                print("Error : docker pull [%s] failure" % pull_image)
                registry.del_docker_registry()
                sys.exit(1)

            tag_push_image = "%s:%s/%s:%s" % (env_10_ip, mapping_host_port, repo, tag)
            res2 = os.system("docker tag %s %s >/dev/null" % (pull_image, tag_push_image))
            time.sleep(2)
            if res2 != 0:
                print("Error : docker tag [%s] [%s] failure" % (pull_image, tag_push_image))
                registry.del_docker_registry()
                sys.exit(1)

            res3 = os.system("docker push %s >/dev/null" % tag_push_image)
            time.sleep(1)
            if res3 != 0:
                print("Error : docker push [%s] failure" % tag_push_image)
                registry.del_docker_registry()
                sys.exit(1)
        print("Info : [%s] directory pack successful" % dire.release_package_name)
        print("\n")


# 下载 chart 包
class pack_charts:
    def helm_operator(self, charts_version, dire):
        print("Info : Start packing charts")
        if os.path.isdir(dire.charts_pack_path):
            print("Error : [%s] directory already exists " % dire.charts_pack_path)
            sys.exit(1)
        else:
            os.makedirs(dire.charts_pack_path)
            print("Info : [%s] pack directory created successful" % dire.charts_pack_path)
            print("\n")

        os.chdir(dire.charts_pack_path)
        for i in charts_version:
            c_name = i.get("name")
            c_version = i.get("version")
            c_namespace = i.get("namespace")
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
        print("Info : [%s] directory pack successful" % dire.charts_pack_path)
        print("\n")
