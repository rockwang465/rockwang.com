# flask项目搭建
## 1.数据库
### A.配置config.py
+ 建表: 
```
mysql> CREATE DATABASE `heros_info`  DEFAULT CHARACTER SET utf8;
mysql> show create database heros_info;
+------------+---------------------------------------------------------------------+
| Database   | Create Database                                                     |
+------------+---------------------------------------------------------------------+
| heros_info | CREATE DATABASE `heros_info` /*!40100 DEFAULT CHARACTER SET utf8 */ |
+------------+---------------------------------------------------------------------+
```
+ 配置config.py

### B.exts.py中创建初始化db
```
# encoding: utf-8

from flask_sqlalchemy import SQLAlchemy

db = SQLAlchemy()  # main.py中db.init_app(app) 进行app的初始化
```

### C.models.py中建表配置
```
# encoding: utf-8

from exts import db  # 这样就可以用exts的db了


class Heros_detail_info(db.Model):
    __tablename__ = 'heros_detail_info'
    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    name = db.Column(db.String(20), nullable=False)
    country = db.Column(db.String(20), nullable=False)
    comment = db.Column(db.String(100), nullable=False)
    img = db.Column(db.Text, nullable=False)
    # 如果text字段长度不够，可以用mediumtext。 text是64kb，mediumtext是16mb。
```

### D.manage.py命令管理
```
# encoding: utf-8

from flask_script import Manager
from flask_migrate import Migrate, MigrateCommand  # 导入mgirate的命令
from flask_main import app
from exts import db
from models import Heros_detail_info

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
```

### E.flask_main.py主入口文件
```
# encoding: utf-8

from flask import Flask
from exts import db
import config

app = Flask(__name__)
app.config.from_object(config)
db.init_app(app)  # db初始化app(在一个文件中db直接初始化；但这里导入外面的db，所以这里传递app给db，给exts.py中的db初始化app用的)
```

### F.创建数据库
+ 初始化数据库，迁移文件，更新到表中，三步操作
```
# python manage.py db init
# python manage.py db migrate
# python manage.py db upgrade
```
+ 查看表
```
mysql> desc heros_detail_info;
+---------+--------------+------+-----+---------+----------------+
| Field   | Type         | Null | Key | Default | Extra          |
+---------+--------------+------+-----+---------+----------------+
| id      | int(11)      | NO   | PRI | NULL    | auto_increment |
| name    | varchar(20)  | NO   |     | NULL    |                |
| country | varchar(20)  | NO   |     | NULL    |                |
| comment | varchar(100) | NO   |     | NULL    |                |
| img     | text         | NO   |     | NULL    |                |
+---------+--------------+------+-----+---------+----------------+
```

## 2.nginx图片存储服务
### A.沿用之前传智前端27期时的docker起的nginx
+ 服务器
```
docker ps -a |grep nginx
docker start xxxxxx
```
+ 检查docker
```
systemctl status docker
```
+ 之前起nginx的命令
```
# docker pull nginx  #镜像拉取
# docker load -i nginx.tar  #如果下不了镜像，就导入镜像
# docker run -dit --name web-api-nginx -p 6001:80 -v /root/images:/usr/share/nginx/html/images -v /root/default.conf:/etc/nginx/conf.d/default.conf -v /root/nginx.conf:/etc/nginx/nginx.conf nginx
```
+ 具体文件配置看《传智前端27期实战记录》 274行

### B.安装nginx服务 -- 二进制这里就不用了
+ rpm安装
```
rpm -Uvh http://nginx.org/packages/centos/7/noarch/RPMS/nginx-release-centos-7-0.el7.ngx.noarch.rpm
yum install -y nginx
systemctl start nginx.service
systemctl enable nginx.service
```
+ 修改端口
```
vi /etc/nginx/conf.d/default.conf  //改为8081或其他端口
```

## 3.flask项目端口运行问题
### A.项目端口运行问题
+ 如果vue请求访问出现跨域问题，通过10.151.x.xx:5000 真实本地ip+端口访问异常。
+ 可以用浏览器先访问下，如果无法请求，说明电脑有异常。
### B.解决方法
+ 修改flask运行的端口，如下面固定端口和host
+ 安装及导入`flask_cors`模块，并使用 `CORS(app, supports_credentials=True)`
```
from flask_cors import CORS
...
if __name__ == "__main__":
    CORS(app, supports_credentials=True)
    app.run(host='0.0.0.0', port='5001', debug=True)
```
+ 正常上面都是可以的，实在不行可以考虑解决windows的防火墙端口规则问题。