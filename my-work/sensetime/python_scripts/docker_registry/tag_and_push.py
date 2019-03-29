#!/usr/bin/env python
# -*- coding: utf-8 -*-
import docker
import os
import argparse


def read_image_list_from_file(image_list_file):
    print("#####    Read Image List from {}   ######".format(image_list_file))
    image_data = []
    with open(image_list_file, mode='r+') as f:
        for line in f:
            image_data.append(line.split())

    return image_data


# load images to local storage
def load_docker_image_to_local(client, image_list, image_dir):
    print("#####   Load Image Into Local Storage   ######")
    for name, image_name in image_list:
        print("Start to load image {}".format(name))
        name += ".tgz"
        file_path = os.path.join(image_dir, name)
        with open(file_path, "r") as fd:
            client.images.load(fd)
        print("Load image {} successfully".format(name))
        print


# docker tag images from old repository to new repository & push it into new repository
def tag_and_push_docker_image(client, old, new):
    print("#####   Retag Image to New Registry Address and Push It   ######")
    images = client.images.list()
    for image in images:
        for tag in image.tags:
            if tag.startswith(old):
                new_tag = tag.replace(old, new)
                print("Retag Image from {} to {}".format(tag, new_tag))
                image.tag(new_tag)
                resp = client.images.push(new_tag)
                print(resp)
                print("Push Image {} successfully".format(new_tag))
                print


def parse_args():
    parser = argparse.ArgumentParser(description='Import docker image from docker tgz file, and then retag and push it '
                                                 'to new _registry')
    parser.add_argument("old_registry", help="Address of old docker registry server")
    parser.add_argument("new_registry", help="Address of new docker registry server")
    parser.add_argument("image_list_file", help="The path of file which contain image list")
    parser.add_argument("--image_dir",
                        help="Directory of image package(default: /tmp/sensenebula-guard-1.0/docker_packages)",
                        default="/tmp/sensenebula-guard-1.0/docker_packages")
    parser.add_argument("--docker_url", help="Docker sock or url location.(default: unix://var/run/docker.sock)",
                        default="unix://var/run/docker.sock")
    return parser.parse_args()


def main():
    args = parse_args()
    client = docker.DockerClient(base_url=args.docker_url, version="auto", timeout=60)
    image_list = read_image_list_from_file(args.image_list_file)
    try:
        load_docker_image_to_local(client, image_list, args.image_dir)
    except Exception as e:
        print("Error occurred when load image:", e)
    try:
        tag_and_push_docker_image(client, args.old_registry, args.new_registry)
    except Exception as e:
        print("Error occurred when push image:", e)


if __name__ == '__main__':
    main()

