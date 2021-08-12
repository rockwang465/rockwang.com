#!/usr/bin/env python
# -*- coding:utf-8 -*-
import requests
import json
import sys
import argparse

# pangoo_url = "http://172.20.7.199:30001"
# pangoo_url = "http://172.20.10.174:8888"

login_path = "/base/login"
get_build_path = "/pipelines/%s/jenkins/%s"
send_msg_path = "/builds/%s/results/%s"

# user = {"username": "jenkins", "password": "jenkins123"}
app_name = "zookeeper"
build_number = 218


# input arguments
def parse_args():
    parser = argparse.ArgumentParser(description='Send notify to ding talk')
    parser.add_argument("--username", help="login the pangoo username", default="jenkins")
    parser.add_argument("--password", help="login the pangoo password", default="jenkins123")
    parser.add_argument("--pangoo_url", help="pangoo platform address", default="http://172.20.10.174:8888")
    return parser.parse_args()


def login(arg):
    login_url = arg.pangoo_url + login_path
    user = {"username": arg.username, "password": arg.password}
    req = requests.post(url=login_url, json=user)

    req_data = json.loads(req.text)
    token = req_data.get("data").get("token")
    if token == "":
        print("Error: can not get token")
        sys.exit(1)
    return token


def get_build(arg):
    token = login(arg)
    get_build_url = arg.pangoo_url + get_build_path % (app_name, build_number)
    headers = {"x-token": token}
    req = requests.get(url=get_build_url, headers=headers)
    req_data = json.loads(req.text)

    build_id = req_data.get("data").get("id")
    if not build_id or build_id == 0:
        print("build id:", build_id)
        print("Error: get build info fail")
        sys.exit(1)
    if build_id & build_id != 0:
        return build_id
    else:
        print("Error: the build info is wrong")
        sys.exit(1)


def send_ding_talk_msg(arg):
    token = login(arg)
    build_id = get_build(arg)
    send_msg_url = arg.pangoo_url + send_msg_path % (build_id, build_number)

    headers = {"x-token": token}
    req = requests.get(url=send_msg_url, headers=headers)
    req_data = json.loads(req.text)
    print(req_data)

    if req_data.get("code") != 200:
        print("Error: send ding talk message fail!")
        sys.exit(1)
    else:
        print("Info: send ding talk message success!")


if __name__ == '__main__':
    args = parse_args()
    send_ding_talk_msg(args)
