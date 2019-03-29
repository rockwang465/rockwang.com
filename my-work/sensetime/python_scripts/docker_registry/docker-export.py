#!/usr/bin/env python
# -*- coding: utf-8 -*-

import socket, docker, os

data = []
registry_sh = '10.5.6.10'
workdir = '/tmp/sensenebula-guard-1.0/docker_packages'
docker_url = 'unix://var/run/docker.sock'
docker_version = '1.24'


def get_local_ip():
    try:
        s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        s.connect(('8.8.8.8', 80))
        ip = s.getsockname()[0]
    finally:
        s.close()
    return ip


# get images name to data list
def get_images_name():
    # f = open(file='repo-10.5.6.10', mode='r+', encoding='utf-8')
    f = open('image_info.txt', mode='r+')
    for line in f:
        global data
        data1 = line.split()
        data.append(data1)
    f.close()


# docker pull images from 10.5.6.10 repo  --> old
# def docker_pull():
#     client = docker.from_env()
#     for name,image_name in data:
#         print("正在docker pull %s" %(name))
#         client.images.pull(image_name)


# docker pull images from 10.5.6.10 repo --> new(有验证)
def docker_pull():
    client = docker.DockerClient(base_url=docker_url, version=docker_version)
    for name, image_name in data:
        if image_name.startswith(registry_sh):
            auth_config = {'username': 'admin', 'password': 'Sensetime12345'}
            print("正在docker pull %s" % (name))
            image = client.images.pull(image_name, auth_config=auth_config)
        else:
            # print("Error : This image name is not 10.5.6.10 head : %s" % (image_name))
            print("Error : This image name is not %s head : %s" % (registry_sh, image_name))
            # image2 = client.images.pull(image_name)


# docker save images to /tmp/ direcotry
def docker_save():
    create_dir = os.system("mkdir -p %s" % (workdir))
    os.chdir(workdir)
    save_list = []
    for name, image_name in data:
        print("正在docker save %s.tar to %s" % (name, workdir))
        recode = os.system("docker save -o %s.tar %s" % (name, image_name))
        if recode != 0:
            print("Error : %s image save fail" % (name))


local_host_ip = get_local_ip()
get_images_name()
docker_pull()
docker_save()

