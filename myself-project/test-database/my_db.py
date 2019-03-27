#!/usr/bin/env python
# -*- coding:utf-8 -*-

import os
import sys

db_file_name = "db_file.txt"
# db_file_path = '/var/mobile/Containers/Data/Application/1A91F6C8-EA17-43F5-A3EA-C42F498E529B/Documents/KeepData/files/db_file.txt'
db_file_path = 'db_file.txt'
notice_format = 'delete/select [feild] from [table_name] where [field] >/=/</like [condition]\n' \
                'add [table_name] ("id","name","age","job","phone")\n' \
                'update [table_name] set [field]="value1" where [feild]="value" '
table_name = "db_info"


# def opera_add():
#     pass
# def opera_delete():
#     pass
# def opera_update():
#     pass
# def opera_select():
#     pass


# 下面4个opera_xxx函数为 增删改查函数。
def opera_add(sql, db_data_list):  # add db_info ("6","wyc","22","AS","13812345678")
    print("add : ", sql)
    if "(" not in sql or ")" not in sql:
        print("Error2 : syntax is error, please input : \n %s" % notice_format)
    else:
        add_head_sql, add_tail_sql = sql.split("(")
        # print(add_head_sql, add_tail_sql)
        # print(len(add_head_sql.split()))
        # print(add_head_sql.split()[1])
        if len(add_head_sql.split()) == 2 and add_head_sql.split()[1] == table_name:
            replace_value1 = add_tail_sql.replace(")", "")
            replace_value2 = replace_value1.replace('"', '')
            replace_value3 = replace_value2.split(",")
            if len(replace_value3) != len(db_data_list):
                # print(replace_value3)
                print("Error : add field amount is error")
            else:   # 这地方有报错，有时间处理
                for i in len(replace_value3):
                    print(db_data_list[i])
                    print(replace_value3[i])
                    # db_data_list[i].append(replace_value3[i])
            print(db_data_list)
        else:
            print("Error3 : syntax is error, please input : \n %s" % notice_format)


def opera_delete(sql, db_data_list):
    print("delete", sql)
    # if "from" not in sql


def opera_update(sql, db_data_list):
    print("update", sql)


def opera_select(sql, db_data_list):
    print("select", sql)


# 此函数为用户输入的内容，进行判断，并且分发给对应的函数(增删改查)。
def get_input(db_data_list):
    # notice_format = "add/delete/update/select [field] from [table_name] where [field] >/=/</like [condition]"
    first_method = {
        'add': opera_add,
        'delete': opera_delete,
        'update': opera_update,
        'select': opera_select
    }

    while True:
        sql = input("Please input sql >>>")
        if not sql: continue  # 注意，必须在input下面一行才可以。
        sql_head = sql.split()[0]
        if sql_head not in ["add", "delete", "update", "select"] or len(sql.split()) < 3:
            print("Error1 : syntax is error, please input :\n %s" % notice_format)
        else:
            for method in first_method:
                if sql_head in method:
                    # print(sql_head, method)
                    first_method[method](sql, db_data_list)


# 打开db_file.txt文件，并转换为数据列表：db_data_list
def get_db_file():
    # get db_file.txt data
    if os.path.exists(db_file_path):
        f1 = open(db_file_path, 'r+')
        f2 = f1.readlines()
        db_list1 = []
        for line in f2:
            db_list1.append(line.split(','))
        # print(db_list1)
    else:
        print("Error : not found %s file" % db_file_name)
        sys.exit(1)

    # transfer data to id,name,age,phone,job,data
    db_data_list = []
    db_info = ["id", "name", "age", "job", "phone"]
    for i in db_info:
        i = []
        db_data_list.append(i)
    # print(db_data_list)

    for k in db_list1:
        for index in range(len(k)):
            # print(k[index].strip())
            # print(db_data_list[index])
            db_data_list[index].append(k[index].strip())
            # print(db_data_list)
    # print(db_data_list)
    return db_data_list


def main():
    db_data_list = get_db_file()
    get_input(db_data_list)


if __name__ == '__main__':
    main()
