#!/usr/bin/env python
# -*- coding: utf-8 -*-

# ********************************************************** #
#  Author        : RockWang                                  #
#  Email         : wangyecheng_vendor@sensetime.com          #
#  Create time   : 2020-03-05                                #
#  Filename      : check_server.py                           #
#  Description   : check server scripts                      #
#  Version       : v2.1.0                                    #
# ********************************************************** #

import socket
import time
import requests
import json
import pymysql
import os
from minio import Minio
from minio.error import ResponseError

ingress_http_port = "30080"
# es_nodeport = "32111"
license_bin = "/usr/local/bin/license_client"
host_domain = "registry.sensenebula.io"
minio_nodeport = "31900"

# ns = ["addons","component","default","helm","ingress","kube-public","kube-system","logging","monitoring","mysql-operator"]
topic = ["senseguard.ac.result.consume", "senseguard.bulk.tool", "senseguard.frontend.comparison",
         "senseguard.guest.management", "senseguard.struct.process.service", "senseguard.sync.device",
         "senseguard.td.result.consume", "stream.features.face_24901", "stream.features.struct", "stream.images.face",
         "sync.stream.features.face_24901"]

bucket = ["GLOBAL", "video_face", "video_panoramic", "keeper_face", "video_pedestrian_cropped",
          "video_pedestrian_panoramic", "video_automobile_cropped", "video_automobile_panoramic",
          "video_cyclist_cropped", "video_cyclist_panoramic", "video_human_powered_vehicle_cropped",
          "video_human_powered_vehicle_panoramic"]

minio_buckets = ["senseguard-face-search", "senseguard-log-export", "senseguard-map-management",
                 "senseguard-records-export", "senseguard-target-export", "senseguard-uums-export",
                 "snapshot-alert-feature-db", "snapshot-struct-timespace-feature-db", "snapshot-timespace-feature-db"]
dbs = ["senseguard", "sys", "uums"]
users = ["GRANT USAGE ON *.* TO `senseguard`@`%`", "GRANT ALL PRIVILEGES ON `senseguard`.* TO `senseguard`@`%`",
         "GRANT ALL PRIVILEGES ON `uums`.* TO `senseguard`@`%`"]


def get_local_ip():
    myname = socket.getfqdn(socket.gethostname())
    myip = socket.gethostbyname(myname)
    return myip


def check_server_usage():
    print("您好: 您当前测试的机器为: %s" % local_host_ip)
    time.sleep(2)


def check_node_state():
    count = 0
    print("\n1. 检查k8s node状态")
    recode = os.system('kubectl get nodes | sed 1d >/dev/null 2>&1')
    if recode == 0:
        state1 = os.popen('kubectl get nodes | sed 1d')
        state2 = state1.readlines()
        if not state2:
            count += 1
            print("Error : 无任何 Node 节点")
        else:
            for line in state2:
                if "NotReady" in line:
                    print("Error : Node 节点 NotReady : %s" % line)
                    count += 1
    else:
        print("Error : 无任何 Node 节点")

    if count == 0:
        print("Node 节点正常")

    # print("\n".rjust(80, '*'))
    time.sleep(2)


def check_pods_state():
    count = 0
    print("\n2. 检查k8s所有pod服务状态")
    recode = os.system('kubectl get pods --all-namespaces | grep -v Running >/dev/null 2>&1')
    if recode == 0:
        state1 = os.popen('kubectl get pods --all-namespaces | grep -v Running | grep -v NAME | grep -v Completed ')
        state2 = state1.readlines()
        for line in state2:
            if "Running" not in line:
                count += 1
                print("Error : 未 Running 的 Pod : %s " % line)
    else:
        count += 1
        print("Error : 无 Pod 服务，请解决")
    if count == 0:
        print("Pod 服务正常")
    # print("\n".ljust(80, '*'))
    time.sleep(2)


def check_license_state():
    print("\n3. 检查加密狗状态")
    recode = os.system('%s status >/dev/null 2>&1' % license_bin)
    if recode == 0:
        state1 = os.popen("%s status | grep 'status is' | awk -F: '{print $2}'" % license_bin)
        state2 = state1.read()
        if "alive" in state2:
            print("加密狗正常")
        else:
            print("Error : 加密狗未激活，请解决")
    else:
        print("Error : 未找到 %s 命令" % license_bin)
    # print("\n".rjust(80, '*'))
    time.sleep(2)


def check_topic_state():
    count = 0
    print("\n4. 检查topic状态")
    recode = os.system(
        "kubectl exec -it -n component kafka-default-0 -- kafka-topics.sh --list --zookeeper zookeeper-default:2181/kafka | grep -v 'consumer_offsets' >/dev/null 2>&1")
    if recode == 0:
        state1 = os.popen(
            "kubectl exec -it -n component kafka-default-0 -- kafka-topics.sh --list --zookeeper zookeeper-default:2181/kafka | grep -v 'consumer_offsets'")
        state2 = state1.read()
        if state2:
            for i in topic:
                if i in state2:
                    pass
                    # print("存在topic : %s" % i)
                else:
                    print("Error : 缺少 topic,请创建 : %s " % i)
                    count += 1
        else:
            count += 1
            print("Error : 无任何topic,请创建")
    else:
        print("Error : 无任何topic,请创建")
    if count == 0:
        print("Topic 正常")
    # print("\n".rjust(80, '*'))
    time.sleep(2)


def check_bucket_state():
    print("\n5. 检查bucket状态")
    try:
        url1 = "http://" + local_host_ip + ':' + ingress_http_port + "/components/osg-default/v1"
        # print(url1)
        req1 = requests.get(url1)
        req_value1 = req1.text
        # print(req_value1)

        if 'buckets' in req_value1:
            url_bucket = []
            for k1 in json.loads(req_value1)["buckets"]:
                if 'name' in k1:
                    url_bucket.append(k1['name'])  # 拿到所有存在的bucket名字
                    # print(url_bucket)
                # else:
                #     print(str("Error : 无任何buckets，请全部创建", encoding='utf-8'))

            if url_bucket:
                count = 0
                for i in bucket:
                    if i not in url_bucket:
                        print("Error : 缺少bucket,请创建 : %s" % i)
                        count += 1
                    # else:
                    #     print("存在bucket : %s" % i)
            else:
                count += 1
                print("Error : 无任何buckets，请全部创建")
            if count == 0:
                print("Bucket 正常")
                time.sleep(1)
        else:
            print("Error : 无任何buckets，请全部创建")
    except  Exception:
        print("Error : 查询失败!")
    # print("\n".rjust(80, '*'))
    time.sleep(2)


def check_minio_bucket():
    print("\n6. 检查minio bucket状态")
    count = 0
    get_buckets_list = []
    minio_domain = host_domain + ":" + minio_nodeport
    minioClient = Minio(minio_domain,
                        access_key='minio',
                        secret_key='minio123',
                        secure=False)
    get_buckets_info = minioClient.list_buckets()
    for bucket_name in get_buckets_info:
        # print(bucket_name.name, bucket.creation_date)
        get_buckets_list.append(bucket_name.name)

    for i in minio_buckets:
        if i not in get_buckets_list:
            print("Error : 缺少 minio bucket,请创建 : %s" % i)
            count += 1
    if count == 0:
        print("minio bucket 正常")
    time.sleep(2)


def check_es_state():
    print("\n7. 检查elasticsearch状态")
    # url2='http://'+ip+':'+es_nodeport+'/_cat/health?pretty'
    recode = os.system(
        "kubectl get svc -n logging | grep elasticsearch-client | awk '{print $5}' | awk -F '[:/]' '{print $2}' >/dev/null 2>&1")
    if recode == 0:
        nodeport1 = os.popen(
            "kubectl get svc -n logging | grep elasticsearch-client | awk '{print $5}' | awk -F '[:/]' '{print $2}'")
        nodeport2 = nodeport1.read()
        if nodeport2:
            try:
                url1 = 'http://' + local_host_ip + ':' + nodeport2.strip() + '/_cluster/health'
                # print(url1)
                req1 = requests.get(url1)
                req_value1 = req1.json()
                if req_value1["status"]:
                    state1 = req_value1["status"]
                    state2 = state1.encode("utf-8")
                    # print(str(state2, encoding="utf-8"))
                    if state2 == "green":
                        print("elasticsearch 状态正常")
                    elif state2 == "yellow":
                        print("Warning : elasticsearch 当前状态为 : %s " % state2)
                    else:
                        print("Error : elasticsearch 状态异常，请及时处理")
                        print("当前状态为 : %s" % state2)
                else:
                    print("Error : elasticsearch 状态异常，请及时处理")
            except  Exception:
                print("Error : 查询失败!")
        else:
            print("Error : 未发现 elasticsearch-client pod")
    else:
        print("Error : 未发现 elasticsearch-client NodePort 端口")
    # print("\n".rjust(80, '*'))
    time.sleep(2)


def check_mysql_state():
    print("\n8. 检查mysql状态")
    # recode = os.system("ls /usr/bin/mysql 2>/dev/null || cp ./package/mysql /usr/bin/mysql && echo 'copy mysql to /usr/bin/mysql' >> /root/check_server.log")
    if os.path.exists("/usr/bin/mysql"):
        try:
            conn = pymysql.connect(host='0.0.0.0', user='root', port=30446, passwd='UVlY88m9suHLsthK', db="mysql",
                                   charset='utf8')
            cur = conn.cursor()  # 获取一个游标
            cur.execute('show databases')
            db_data = cur.fetchall()
            cur.execute('show grants for  senseguard@"%"')
            user_data = cur.fetchall()
            # 注意int类型需要使用str函数转义
            cur.close()  # 关闭游标
            conn.close()  # 释放数据库资源

            # 检查3个databases是否创建
            db_get = []
            for i in db_data:
                db_get.append(i[0])
            db_num = 0
            for i in dbs:
                if i not in db_get:
                    db_num += 1
                    print("Error : 缺少databases,请创建 : %s" % i)
            if db_num == 0:
                print("databases已创建")

            # 检测user是否授权
            # print(user_data)
            users_get = []
            for i in user_data:
                users_get.append(i[0])
            user_num = 0
            for i in users:
                if i not in users_get:
                    user_num += 1
                    print("Error : 缺少users授权，请授权 : %s" % i)
            if user_num == 0:
                print("user已授权")

        except  Exception:
            print("Error : 查询失败!")
    else:
        print("Error : 未找到 /usr/bin/mysql 命令 ")
    # print("\n".rjust(80, '*'))
    time.sleep(2)


if __name__ == '__main__':
    local_host_ip = get_local_ip()
    check_server_usage()
    check_node_state()
    check_pods_state()
    check_license_state()
    check_topic_state()
    check_bucket_state()
    check_minio_bucket()
    check_es_state()
    check_mysql_state()
