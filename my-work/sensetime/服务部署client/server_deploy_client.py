#!/usr/bin/env python
# -*- coding: utf-8 -*-

# import requests
import argparse
import os
import sys
import requests

ns = ['default', 'logging', 'monitoring', 'component', 'nebula', 'guard']


def install_services(args):
    # 先做get试试
    # 请求所有的charts版本
    url = "http://10.5.6.14:30001/v1/charts"
    req = requests.get(url)
    print(req.text)


def check_args(args):
    if args.namespace not in ns:
        print("Error : Please input namespace : %s" % (" ".join(ns)))
        sys.exit(1)
    if not args.workdir:
        print("Error : Not exist workdir, Must be 2 arguments")
        sys.exit(1)
    print("Current input arguments : (namespace:) %s , (workdir:) %s" % (args.namespace, args.workdir))


def parse_arg():
    parse = argparse.ArgumentParser(description="Install services in kubernetes")
    parse.add_argument('install', help='Install services')
    parse.add_argument('-ns', '--namespace', help='Install all services under the specified namespace',
                       default='default')
    parse.add_argument('-wd', '--workdir',
                       help='Install server directory path, the default path "." is the current path', default='.')
    parse.add_argument('-f', '--file',
                       help='Install server directory path, the default path "." is the current path', default='.')
    return parse.parse_args()


def main():
    args = parse_arg()
    check_args(args)
    install_services(args)


if __name__ == '__main__':
    main()
