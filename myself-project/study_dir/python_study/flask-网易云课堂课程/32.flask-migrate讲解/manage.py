# encoding: utf-8

# from flask import Flask
# import config
# from models import Article

from flask_script import Manager
from flask_migrate import Migrate, MigrateCommand  # 导入mgirate的命令
from migrate_demo import app
from exts import db
from models import Article

# app = Flask(__name__)
# app.config.from_object(config)

# db.init_app(app)  # 给db传入一个app进行初始化
# app推入到栈顶中(即把app推入app上下文中)
# with app.app_context():  # 如果是导入文件，则需要用此推入上下文方式，否则报错
#     db.create_all()


# 1. manager初始化app
manager = Manager(app)

# 2. 使用flask_migrate，必须绑定app和db
migrate = Migrate(app, db)

# 3. 把MigrateCommand命令添加到manager中
# 3.1 上面导入models 的 Article类，让 MigrateCommand 来执行数据库的操作
manager.add_command('db', MigrateCommand)

if __name__ == '__main__':
    manager.run()  # 这里是执行manager命令

# 4. linux命令行执行
# 4.1 【初始化迁移环境】。 如果没有报错，则表示执行成功
#     执行后，当前目录下会生成migrations目录
# python manage.py db init
# 4.2 【生成迁移文件】，migrations/versions/ 目录下会生产 .py文件
#     且会生成一个 alembic_version 的表，用于定义迁移版本号，后期可以考虑学习下。
# python manage.py db migrate
# 4.3 【将迁移文件映射到表中】
#     此时就会生成定义的article的表
# python manage.py db upgrade
# 4.4 后期需要新增字段时，例如在models.py中新加一个tags字段
#     则执行migrate生成迁移文件，和执行upgrade将迁移文件映射到表中。
# python manage.py db migrate
# python manage.py db upgrade
# mysql> desc article;  # 发现新增字段添加上去了
# +---------+--------------+------+-----+---------+----------------+
# | Field   | Type         | Null | Key | Default | Extra          |
# +---------+--------------+------+-----+---------+----------------+
# | id      | int(11)      | NO   | PRI | NULL    | auto_increment |
# | title   | varchar(100) | NO   |     | NULL    |                |
# | content | varchar(100) | NO   |     | NULL    |                |
# | tags    | varchar(100) | NO   |     | NULL    |                |
# +---------+--------------+------+-----+---------+----------------+
