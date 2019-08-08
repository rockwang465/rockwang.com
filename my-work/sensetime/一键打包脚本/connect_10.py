#!/usr/bin/env python
# -*- coding: utf-8 -*-
import paramiko
import sys
import os


class scp_files:
    # send files to 10.5.6.10
    def ssh_scp_put(self, ip_10, port, username, passwd, local_file, remote_file):
        """
        :param ip_10: 服务器10.5.6.10 ip地址
        :param port: 端口(22)
        :param username: 用户名
        :param passwd: 用户密码
        :param local_file: 本地文件地址
        :param remote_file: 要上传的文件地址（例：/tmp/test.txt）
        :return:
        """
        try:
            client = paramiko.Transport((ip_10, port))
            client.connect(username=username, password=passwd)
            sftp = paramiko.SFTPClient.from_transport(client)
            print('Info : start upload file [%s] ' % local_file)

            try:
                sftp.put(local_file, remote_file)
            except Exception as ex:
                sftp.mkdir(os.path.split(remote_file)[0])
                sftp.put(local_file, remote_file)
                print("Info : from local : [%s] upload to : [%s]" % (local_file, remote_file))
            # print('Info : upload successful %s ' % datetime.datetime.now())
            print('Info : upload successful')
            client.close()
        except Exception as ex:
            print("Error : Failure to transport file" % ex)
            sys.exit(1)
        print("\n")

    def exec_10_script(self, ip_10, port, username, passwd, exec_script_file):
        try:
            ssh = paramiko.SSHClient()
            ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
            ssh.connect(ip_10, port, username, passwd)
            stdin, stdout, stderr = ssh.exec_command("/usr/bin/python %s" % exec_script_file)
            res, err = stdout.read().decode(), stderr.read().decode()
            if err:
                print("Error : error data", err)
                sys.exit(1)
            else:
                print("Info : res data", res)
            ssh.close()
        except Exception as ex:
            print("Error : Failure to execute script : %s" % ex)
            sys.exit(1)  # 异常则退出
        print("\n")
