#!/usr/bin/env python
# -*- coding: utf-8 -*-
import os
import sys
import paramiko

port = 22
username = "root"
passwd = "Nebula123$%^"
es_name = "elasticsearch"
es_ns = "logging"
cut_value = '{print $1" "$2" "$3" "$7}'
es_data_path = "/mnt/locals/elasticsearch/volume0/nodes"

es_error = ["failed to read local state", "failed to read", "file:/usr/share/elasticsearch/data/nodes/0/indices"]
# failed to read local state, exiting...
# org.elasticsearch.ElasticsearchException: java.io.IOException: failed to read [id:54, file:/usr/share/elasticsearch/data/nodes/0/indices/FSSmfd_FQpauB2ZYwdFQ8A/_state/state-54.st]


# 获取当前es非Running状态的服务信息
def get_info():
    recode_line1 = os.popen("kubectl get pods -n %s  -o wide | grep %s | egrep -v 'Running|elasticsearch-curator' | wc -l" % (es_ns, es_name)).read()
    recode_line2 = int(recode_line1.replace("\n", ""))
    recode_value1 = os.popen("kubectl get pods -n %s  -o wide | grep %s | egrep -v 'Running|elasticsearch-curator' | awk '%s'" % (es_ns, es_name, cut_value))

    info = []
    if recode_line2 == 1:
        recode_value2 = recode_value1.read().replace("\n", "")
        tmp2 = recode_value2.split(" ")
        info.append({
            'name': tmp2[0],
            'health': tmp2[1],
            'status': tmp2[2],
            'host': tmp2[3]
        })
        return info
    elif recode_line2 > 1:
        recode_value2 = recode_value1.readlines()
        for i in recode_value2:
            tmp1 = i.replace("\n", "")
            tmp2 = tmp1.split(" ")
            info.append({
                'name': tmp2[0],
                'health': tmp2[1],
                'status': tmp2[2],
                'host': tmp2[3]
            })
        return info
    elif recode_line2 == 0:
        print("Info : elasticsearch 状态正常")
        sys.exit(0)
    else:
        print("Error : 捕捉异常 %s" % recode_line2)


# 获取es的日志
def get_log(info):
    # print(info)
    for i in info:
        name = i.get("name")
        status = i.get("status")
        host = i.get("host")
        log = os.popen("kubectl logs --tail=1000 %s -n %s" % (name, es_ns)).read()
        print("Info : 检查 %s pod服务" % name)
        check_log(log, host, name, status)
        print("\n")


# 检查日志中是否有报错的关键字
def check_log(log, host, name, status):
    # print(log)
    j = 0
    for i in es_error:
        if i in log:
            j += 1
    if j == len(es_error):
        print("Error : 确认有错误")
        del_data(host, name, status)
    else:
        print("Warning : elasticsearch有异常，请检查")


# 删除对应主机host中的/mnt/locals/elasticsearch/volume0/nodes/下的数据
def del_data(host, name, status):
    # 这里传入的host是es中显示的主机名，但paramiko中的host尽量以ip为主，主机名也可以
    ssh = paramiko.SSHClient()
    try:
        ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
        ssh.connect(host, port, username, passwd)

        stdin, stdout, stderr = ssh.exec_command("rm -rf %s" % es_data_path)
        print(stdout.read().decode())
        print("Info : 已成功删除 %s 目录" % es_data_path)
        ssh.close()
        restart_server(name, status)
    except Exception as ex:
        print("Error : 未执行成功 : %s" % ex)


def restart_server(name, status):
    if status == 'Running':
        recode = os.system("kubectl delete pods %s -n %s" % (name, es_ns))
        if recode == 0:
            print("Info : %s 服务已重启" % name)
        else:
            print("Error: %s 重启失败，请检查" % name)


if __name__ == "__main__":
    info = get_info()
    get_log(info)
