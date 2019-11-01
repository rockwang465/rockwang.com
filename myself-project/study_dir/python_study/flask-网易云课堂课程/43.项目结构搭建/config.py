# encoding: utf-8
import os

DIALECT = 'mysql'
# DRIVER = 'mysqldb'
DRIVER = 'pymysql'
USERNAME = 'root'
PASSWORD = 'Rock1314^'
HOST = '10.11.172.251'
PORT = '3306'
DATABASE = 'zlkt_demo'  # 注意: 数据库要确定存在，不存在要先创建

SQLALCHEMY_DATABASE_URI = "{}+{}://{}:{}@{}:{}/{}?charset=utf8".format(DIALECT, DRIVER, USERNAME, PASSWORD, HOST, PORT,
                                                                       DATABASE)
SQLALCHEMY_TRACK_MODIFICATIONS = True  # 取消警告

DEBUG = True
SECRET_KEY = os.urandom(24)
