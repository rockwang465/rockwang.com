#!/usr/bin/env python
# -*- coding: utf-8 -*-

# import requests
import argparse
import os
import sys
import requests
import yaml

from devops import *

req_host = "10.5.6.14"
req_port = "8090"
username = "admin"
password = "sensetime"
package_path = "/opt/SenseNebula-G-v1.1.0+20190621134305"
charts_version_path = package_path + "/versions.yaml"

# jinja2 to render define
# docker_registry = "10.5.6.14:5000"
master_ip = "10.5.6.14:5000"
slave_ip = "127.0.0.1"
# standalone or distributed
deploy_mode = "standalone"

all_namespaces = ['default', 'logging', 'monitoring', 'component', 'nebula', 'guard']
headers = {
    "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
}

temp_dir = ['devops', 'component', 'nebula']

# install_ns_services函数是安装用户输入的名称空间下的所有服务
# def install_ns_services(ns, headers):
#     print("[Info] : Install component namespace services")
#     componet_charts_api = ""
#     all_charts = all_services["charts"]
#     # print(all_charts)
#
#     for i in all_charts:
#         # print(i)
#         if i.get(ns):
#             print("This is %s  ------>" % ns)
#             print(i.get(ns))
#
#     body = {
#         "instance": {
#             "name": "cassandra",
#             "namespace": ns,
#             "app_name": "cassandra",
#             "version": "3.11.2-master-35f939d",
#             "config": "values.yaml",
#             "revision": 0,
#             "status": "string",
#             "created": "2019-07-04T03:26:41.273Z",
#             "updated": "2019-07-04T03:26:41.273Z"
#         },
#         "recreate_pods": "true"
#     }
#
#     body2 = {
#         "instance": {
#             "name": "cassandra",
#             "namespace": ns,
#             "app_name": "cassandra",
#             "version": "3.11.2-master-35f939d",
#             "config": "values.yaml",
#             "revision": 0,
#             "status": "true"
#         },
#         "recreate_pods": "true"
#     }
#
#     # 安装服务API接口
#     install_charts_api = "/v1/instances"
#     url = "http://" + req_host + ":" + req_port + install_charts_api
#
#     # 开始调用api执行安装
#     req = requests.post(url, json=body2, headers=headers)
#     print(req.text)
#
#     # 1.判断是否当前服务器有这个服务，并判断是否为最新的
#     #   1.1 如果是最新的，不用安装
#     #   1.2 如果不是最新的，则需要更新


class install_devops():
    def __init__(self):
        print("install devops")
        d1 = devops()
        d1.sum()


class get_jinja2():
    def __init__(self):
        print('python jinja2_convert.py')
        # os.system('python jinja2_convert.py')


# class query_charts_version(ns, headers):
class query_charts_version():
    def __init__(self, ns, headers):
        # 查询服务版本API接口  -----------------【不确定是查本机还是所有版本中的某个版本】
        # name和version可以在查找所有版本中的信息中确认下结果
        name = "cassandra"
        # version = "3.11.2-master-35f939d"
        version = "3.4.10-dev-9266726"
        query_charts_api = "/v1/charts" + "/" + name + "/" + version
        shili = "/v1/instances"
        url = "http://" + req_host + ":" + req_port + query_charts_api
        url2 = "http://" + req_host + ":" + req_port + shili
        # print(url)

        # 开始执行查询操作
        req = requests.get(url, headers=headers)
        req2 = requests.get(url2, headers=headers)
        # req = requests.get(url)
        print("查询单个Cassandra---------")
        print(req.text)
        print("查询实例------------------")
        print(req2.text)


# analysis_args函数是用于分析args并传入到对应函数中执行
class analysis_yaml():
    def __init__(self):
        print("[Info] : Now start to analysis yaml file")
        # 拿到version.yaml文件，使用yaml转为json数据
        data = open(charts_version_path, 'r')
        json_data = yaml.full_load(data.read())
        if not json_data:
            print("[Error] : json数据转换失败")
            sys.exit(1)
        # return json_data
        self.json_data = json_data


# get_charts_info函数是用于获取所有的charts的所有版本信息
# class get_charts_info(args, headers):
class get_charts_info():
    def __init__(self, args, headers):
        # get获取所有的charts版本
        get_all_charts_api = "/v1/charts"
        url = "http://" + req_host + ":" + req_port + get_all_charts_api

        req = requests.get(url, headers=headers)
        # print(req.text)

        # 将版本信息写到文件中
        if req.text:
            # print(req.text)
            f = open('charts_info.txt', 'w')
            f.write(req.text)
            f.close()
        else:
            print("[Info] : No any charts server")
        # print("get charts info")


# check_packages函数是用于检测/opt/目录下的安装文件是否完整
class check_packages():
    def __init__(self):
        if not os.path.exists(package_path):
            print("[Error] : The %s path not found ." % package_path)
            sys.exit(1)
        if not os.path.exists(charts_version_path):
            print("[Error] : The %s file not found ." % charts_version_path)
            sys.exit(1)
        print("[Info] : %s and %s is exists ." % (package_path, charts_version_path))


# login函数是用于API请求时的登录认证
class login():
    def __init__(self):
        # 请求认证地址
        url = "http://" + req_host + ":" + req_port + "/v1/authenticate"
        # 登录账号密码
        body = {
            "username": username,
            "password": password
        }
        r = requests.post(url, json=body, headers=headers)
        # print(r.text)
        token = r.json()["token"]

        # 返回token内容，给headers使用
        # return "Bearer" + " " + token
        new_token = "Bearer" + " " + token
        headers["Authorization"] = new_token
        # return headers
        self.headers = headers


# chekc_args函数是用于简单检查用户在命令行的传参是否合法
# class check_args(args):
class check_args():
    def __init__(self, args):
        if args.namespace not in all_namespaces:
            print("Error : Please input namespace : %s" % (" ".join(all_namespaces)))
            sys.exit(1)
        if not args.workdir:
            print("Error : Not exist workdir, Must be 2 arguments")
            sys.exit(1)
        print("[Info] : Current input arguments : (namespace:) %s , (workdir:) %s" % (args.namespace, args.workdir))


# parse_arg函数用于定义用户在命令行的传参，方便后面调用
class parse_arg():
    def __init__(self):
        parse = argparse.ArgumentParser(description="Install services in kubernetes")
        parse.add_argument('install', help='Install services')
        parse.add_argument('-u', '--upgrade', help='Upgrade services, current has this service')
        parse.add_argument('-i', '--install', help='Install services, current no this service')

        parse.add_argument('-n', '--name', help='Install service name')
        parse.add_argument('-ns', '--namespace', help='Install all services under the specified namespace',
                           default='default')
        parse.add_argument('-wd', '--workdir',
                           help='Install server directory path, the default path "." is the current path', default='.')
        parse.add_argument('-f', '--file',
                           help='Install server directory path, the default path "." is the current path', default='.')

        #  print(parse.parse_args())
        self.parse = parse.parse_args
        #  return parse.parse_args()


class main():
    def __init__(self):
        # 拿到用户传入的参数
        args = parse_arg()
        # print(args.parse())
        # 检查参数是否合法
        check_args(args.parse())
        # 登录认证
        login_headers = login()
        # 检查包的状况
        check_packages()
        # 获取所有的charts的所有版本信息
        get_charts_info(args.parse(), login_headers.headers)
        # 转换yaml文件为json数据versions
        versions = analysis_yaml()
        # 生成每个服务的jinja2的overwide.yaml文件
        get_jinja2()

        # 安装devops服务(logging、monitoring)
        install_devops()


if __name__ == '__main__':
    main()
