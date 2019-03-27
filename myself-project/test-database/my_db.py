#!/usr/bin/env python
# -*- coding:utf-8 -*-

import os
import sys

db_file_name = "db_file.txt"
db_file_path = '/var/mobile/Containers/Data/Application/1A91F6C8-EA17-43F5-A3EA-C42F498E529B/Documents/KeepData/files/db_file.txt'
# db_file_path = 'db_file.txt'
notice_format = 'delete/select [feild] from [table_name] where [field] >/=/</like [condition]\nadd [table_name] ("id","name","age","job","phone")\nupdate [table_name] set [field]="value1" where [feild]="value" '
table_name = "db_info"


# def opera_add():
#     pass
# def opera_delete():
#     pass
# def opera_update():
#     pass
# def opera_select():
#     pass

def opera_add(sql):
    print("add : ", sql)
    if "(" not in sql or ")" not in sql:
        print("Error2 : syntax is error, please input : \n %s" % notice_format)
    else:
        add_head_sql, add_tail_sql = sql.split("(")
        print(add_head_sql, add_tail_sql)
        if add_head_sql[0].split() == 2 and (add_head_sql[0].split())[1] == table_name:
            pass
        else:
            print("Error : syntax error")


def opera_delete(sql):
    print("delete", sql)
    # if "from" not in sql


def opera_update(sql):
    print("update", sql)


def opera_select(sql):
    print("select", sql)


def get_input(db_data_list):
    # notice_format = "add/delete/update/select [feild] from [table_name] where [field] >/=/</like [condition]"
    first_method = {
        'add': opera_add,
        'delete': opera_delete,
        'update': opera_update,
        'select': opera_select
    }

    while True:
        sql = input("Please input sql >>>")
        sql_head = sql.split()[0]
        if not sql:
            continue
        elif sql_head not in ["add", "delete", "update", "select"] or len(sql.split()) < 3:
            print("Error1 : syntax is error, please input : %s" % notice_format)
        else:
            for method in first_method:
                if sql_head in method:
                    # print(sql_head, method)
                    first_method[method](sql)


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
