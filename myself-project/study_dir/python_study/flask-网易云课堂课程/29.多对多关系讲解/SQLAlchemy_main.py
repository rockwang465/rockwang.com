# -*- encoding: utf-8 -*-

from flask import Flask
from flask_sqlalchemy import SQLAlchemy
import config  # 导入config文件

app = Flask(__name__)
app.config.from_object(config)  # 使用config文件的配置
db = SQLAlchemy(app)  # 初始化SQLAlchemy

#  create table article (
#      id int primary key autoincrement,
#      title varchar(100) not null,
#  )
#
#  create table tag (
#      id int primary key autoincrement,
#      name varchar(50) not null,
#  )
#
#  create table article_tag (
#      article_id int,
#      tag_id int,
#      primary key('article_id', 'tag_id'),
#      foreign key `article_id` references `article.id`,
#      foreign key `tag_id` references `tag.id`,
#  )

# 创建article_tag表
article_tag = db.Table('article_tag',
                       db.Column('article_id', db.Integer, db.ForeignKey('article.id'), primary_key=True),
                       db.Column('tag_id', db.Integer, db.ForeignKey('tag.id'), primary_key=True)
                       )


class Article(db.Model):
    __tablename__ = 'article'
    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    title = db.Column(db.String(100), nullable=False)

    # 关联 Tag模型和article_tag表,并做反向引用。通过参数 `secondary=中间表` 进行关联
    tags = db.relationship('Tag', secondary=article_tag, backref=db.backref('articles'))


class Tag(db.Model):
    __tablename__ = 'tag'
    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    title = db.Column(db.String(100), nullable=False)


db.create_all()  # 开始运行数据库的操作


@app.route('/')
def index():
    # 插入一些基础数据
    article1 = Article(title='aaa1')
    article2 = Article(title='aaa2')
    tag1 = Tag(title='bbb1')
    tag2 = Tag(title='bbb2')

    # article_tag表中插入上面4条对应的数据的id
    # 即 多对多关系
    article1.tags.append(tag1)
    article1.tags.append(tag2)
    article2.tags.append(tag1)
    article2.tags.append(tag2)

    db.session.add(article1)
    db.session.add(article2)
    db.session.add(tag1)
    db.session.add(tag2)

    db.session.commit()
    article1 = Article.query.filter(Article.title == 'aaa').first()
    tags = article1.tags
    for tag in tags:
        print(tag.name)

    return 'index'


if __name__ == "__main__":
    app.run(debug=True)
