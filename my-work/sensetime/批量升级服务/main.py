#!/usr/bin/env python
# -*- coding: utf-8 -*-

import os,sys

# 思路:
# 1.需要一个yaml文件，已经写好的charts服务名、charts版本、名称空间
# 2.需要对应的override.yaml文件，建议提前配好templates模板文件，可以从deploy中拿
# 3.脚本功能
#    a. 转换yaml文件
#    b. 转换模板文件
#    c. 拉取charts包
#    d. 安装服务
#    e. 对报错的服务确认是否未镜像问题。 如果时镜像拉取问题，则解决此问题。
#    f. 检测服务，并提示是否完成部署。

