# encoding: utf-8
from flask_script import Manager
from flask_script_demo import app
from db_script import DBManager

# 1 这里的app可以自己初始化；
# 或者用外面的app，导入进来即可，这里用外面导入进来的(flask_script_demo.py)
manager = Manager(app)


# 2 常用操作
@manager.command
def runserver():
    print('服务跑起来了')


# 3 导入命令操作
# add_command表示添加命令
# 'db'表示给这个命令起一个主的引用名， DBManager表示引用哪个命令
manager.add_command('db', DBManager)

if __name__ == '__main__':
    manager.run()

# 5. Linux命令行执行方法
# python manage.py runserver # 执行manager.py自带的runserver函数
# python manage.py db init  # 执行db中的init命令
# python manage.py db migrate  # 执行db中的migrate命令
