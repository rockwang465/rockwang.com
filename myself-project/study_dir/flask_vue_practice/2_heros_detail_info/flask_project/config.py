# encoding: utf-8

# dialect+driver://username:password@host:port/database
DIALECT = 'mysql'
# DRIVER = 'mysqldb'
DRIVER = 'pymysql'
USERNAME = 'root'
PASSWORD = 'Rock1314^'
HOST = '192.168.137.128'
PORT = '3306'
DATABASE = 'heros_info'  # 注意: 数据库要确定存在，不存在要先创建

SQLALCHEMY_DATABASE_URI = "{}+{}://{}:{}@{}:{}/{}?charset=utf8".format(DIALECT, DRIVER, USERNAME, PASSWORD, HOST, PORT,
                                                                       DATABASE)
# SQLALCHEMY_COMMIT_ON_TEARDOWN = True
SQLALCHEMY_TRACK_MODIFICATIONS = True  # 取消警告
