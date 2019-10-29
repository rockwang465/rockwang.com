# -*- encoding: utf-8 -*-

from flask import Flask
from flask_sqlalchemy import SQLAlchemy
import config  # 导入config文件

app = Flask(__name__)
app.config.from_object(config)  # 使用config文件的配置

# 初始化SQLAlchemy
db = SQLAlchemy(app)


# 创建 article表的 SQL 语句:
# create table article (
#   id int primary key autoincrement,
#   title varchar(100) not null,
#   content text not null,
# )

class Article(db.Model):  # 继承db的Model模型
    __tablename__ = 'article'  # 表名定义
    # db.Column 表示创建一个字段
    # db.Integer 表示为数字, primary_key 为主键, autoincrement为自增
    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    # db.String(100) 表示varchar字符串长度为100
    # nullable 表示是否可以为空，这里False表示不允许为空
    title = db.Column(db.String(100), nullable=False)
    # db.Text 表示文本
    content = db.Column(db.Text, nullable=False)


db.create_all()  # 开始运行数据库的操作


@app.route('/')
def index():
    return 'index'


if __name__ == "__main__":
    app.run(debug=True)
