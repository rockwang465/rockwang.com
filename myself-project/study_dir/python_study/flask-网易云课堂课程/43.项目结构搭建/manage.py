# encoding: utf-8

from flask_script import Manager
from flask_migrate import MigrateCommand, Migrate
from main import app
from exts import db

manager = Manager(app)

# 使用migrate帮i的那个app和db
migrate = Migrate(app, db)

# 添加迁移的脚本到manager中
manager.add_command('db', MigrateCommand)

if __name__ == '__main__':
    manager.run()
