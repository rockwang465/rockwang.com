#!/usr/bin/env python
# -*- coding: utf-8 -*-

import yaml
import os

docker_host_10 = "10.5.6.10"
docker_host_loacl = "10.5.6.68:5000"


# get old images
def get_old_images():
    data = open('./images.yaml', 'r')
    old_images = yaml.full_load(data.read())
    # print(old_images)
    return old_images


def get_current_machine_images():
    # result1 = os.popen("kubectl get pods --all-namespaces -o yaml | grep 'image:' | awk '{print $NF}' | sort -rn | uniq | sed 's#%s##g' | sed 's#%s##g' | sed 's#^/##g' | grep ':' | sort -rn | uniq" % (docker_host_10, docker_host_loacl))
    result1 = os.popen(
        "kubectl get pods --all-namespaces -o yaml | grep 'image:' | awk '{print $NF}' | sort -rn | uniq | sed 's#%s##g' | sed 's#%s##g' | sed 's#^/##g' | grep ':' | sort -rn | uniq | grep -v '583ca717-2019050520'" % (
            docker_host_10, docker_host_loacl))
    result2 = result1.readlines()

    # convert images string data to a list data
    machine_images = {}
    for i in result2:
        tmp1 = i.replace("\n", "").split(",")
        # print(tmp1)
        tmp2 = tmp1[0].split(":")

        if machine_images.get(tmp2[1]):
            print("[Error] : already have this image data")
            print("[Error] : old data : %s:%s" % (tmp2[0], machine_images.get(tmp2[1])))
            print("[Error] : new data : %s:%s" % (tmp2[0], tmp2[1]))
        else:
            machine_images[tmp2[0]] = tmp2[1]
    # print(machine_images)
    return machine_images


def compare_images(old_images, new_images):
    # old_images example : [{'tag': '1.2.6', 'repository': 'kubernetes/coredns'}, {'tag': 'v0.10.0-amd64', 'repository': 'coreos/flannel'}...]
    # new_images example : {'kubernetes/defaultbackend': '1.4', 'mysql/mysql-server': '8.0.16'}
    new_images_version_dict = {'all_images': []}

    for index, i in enumerate(old_images['all_images']):
        images_name = i.get("repository")
        old_images_version = i.get("tag")
        new_images_version = new_images.get(images_name)

        # when new_images_version value is empty, then use docker images
        if not new_images_version:
            # print(images_name)
            result1 = os.popen("docker images | grep %s | awk '{print $2}' | sort -rn | uniq | wc -l" % images_name)
            result2 = result1.readlines()[0].replace("\n", "")
            if result2 == "1":
                result3 = os.popen("docker images | grep %s | awk '{print $2}' | sort -rn | uniq" % images_name)
                result4 = result3.readlines()[0].replace("\n", "")
                old_images['all_images'][index]['tag'] = result4
                print(
                    "[Warning] :  [%s] image name not found in [kubectl get pods --all-namespaces -o yaml], Now use the image from [docker images] , the version is : %s" % (
                        images_name, result4))
            else:
                print(
                    '[Error] : The [kubernetes get pods --all-namespaces -o yaml] no %s images ,but [docker images] has multiple, please check' % images_name)

        # when old_images_version == new_images_version
        elif old_images_version == new_images_version:
            print("[Info] : [%s] image version is right" % images_name)
            # print(index)

        # when old_images_version != new_images_version
        else:
            print("[Warning] :  [%s] image version not right, the old version is : %s , the new version is : %s" % (
                images_name, old_images_version, new_images_version))
            # use index to set new version
            old_images['all_images'][index]['tag'] = new_images_version
        # os.system("sleep 1")
    # print(old_images)

    f = open('./new_images.yaml', 'w')
    yaml.dump(old_images, f)


old_images = get_old_images()
new_images = get_current_machine_images()
compare_images(old_images, new_images)
