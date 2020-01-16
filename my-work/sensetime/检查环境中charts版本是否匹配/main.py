#!/usr/bin/env python
# encoding: utf-8

# +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# creater: Rock Wang                                                  +
# creation time: 2020-01-15                                           +
# description: version compare + fetch patch packages + update patch  +
# +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

from version_compare import *
import os

compare_helm_file = "compare_charts_version.txt"


class base_info:
    def user_input(self):
        print("+" * 30)
        print("1. version comparison")
        print("2. fetch the patch packages")
        print("3. update patch packages")
        print("+" * 30)
        self.user_select = input("Please enter 1|2|3 :")
        print("\n")

    def USAGE(self):
        print("Error: please enter: 1|2|3")
        sys.exit(1)


if __name__ == "__main__":
    base = base_info()
    base.user_input()

    get_charts = get_charts_version()

    if base.user_select == 1:  # 如果是版本比对
        get_charts.helm_version()
        compare_charts = compare_charts_version(compare_helm_file)
        compare_charts.get_version_file()
        compare_charts.compare_helm(get_charts.env_charts_info)
        print("\nInfo: Comparison completed")
    elif base.user_select == 2:  # 如果是拉取补丁包
        print("拉取补丁包功能")
    elif base.user_select == 3:  # 更新补丁包
        print("更新补丁包功能")
    else:
        base.USAGE()
