# -*- encoding: utf-8 -*-

from flask import Flask
from flask_sqlalchemy import SQLAlchemy
import config  # 导入config文件

app = Flask(__name__)
app.config.from_object(config)  # 使用config文件的配置
db = SQLAlchemy(app)  # 初始化SQLAlchemy


# 用户表创建sql
# create table user2 (
#     id int primary key autoincrement
#     username varchar(100) not null
#  )

# 文章表创建sql
#  create table article2 (
#      id int primary key autoincrement,
#      title varchar(100) not null,
#      content text not null,
#      author_id int,
#      # 这里做外键约束，让author_id 绑定 user2表的id
#      foreign key `author_id` references `user2.id`
#  )

class User(db.Model):  # 继承db的Model模型
    __tablename__ = 'user2'  # 表名定义
    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    username = db.Column(db.String(100), nullable=False)


class Article(db.Model):  # 继承db的Model模型
    __tablename__ = 'article2'  # 表名定义
    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    title = db.Column(db.String(100), nullable=False)
    content = db.Column(db.Text, nullable=False)
    # 此处通过db.ForeignKey来设置外键约束; 但db.Integer一定要加
    author_id = db.Column(db.Integer, db.ForeignKey('user2.id'))

    # D. 将表Aticle与表User关联起来,实现需求：
    # 1、可以通过一个article 标题查找到对应的作者（username）
    # 2、db.backref 可以通过作者username 查找到对应作者写的所有文章
    # 注意：
    #     1、关联的class名要用引号括起来'User'，关联User模型，则会找有关联的只有author_id这个外键。
    #     2、反向关联写成 backref = db.backref('articles'),引号里可以随便命名，以后就用这个名称来关联
    # Rock:  给Article模型添加了一个author属性，就可以访问这篇文章的作者的数据，
    #        就像属于模型内的字段一样。  【这就相当于是联表后的结果，让表内多了你要的字段！！】
    #        backref是定义反向引用，可以通过User.articles访问这个模型缩写的所有文章。
    author = db.relationship('User', backref=db.backref('articles'))
    # Article 用 author为正向引用。
    # User 用 articles 为反向引用。


db.create_all()  # 开始运行数据库的操作


@app.route('/')
def index():
    # # A. 想要添加一篇文章，因为文章必须依赖用户而存在，所以首先添加一个用户
    # user1 = User(username='zhiliao')
    # db.session.add(user1)
    # db.session.commit()
    # return 'index'

    # # B. 新增文章
    # article = Article(title='aaa', content='bbbbb3', author_id=1)
    # db.session.add(article)
    # db.session.commit()
    # return 'index'

    # # C. 我要找文章标题为aaa的这个作者
    # article = Article.query.filter(Article.title == 'aaa').first()
    # author_id = article.author_id  # 拿到上面结果中的author_id对应的值
    # user = User.query.filter(User.id == author_id).first()
    # username = user.username
    # print(username)
    # return 'index'

    # D. 比C更快找到文章标题为aaa的作者 -- ***** 【重点、难点】
    #    见上面 db.relationship('User', backref=db.backref('articles')) 用法定义
    #    2行代码查到作者(user2表中的username)
    # article = Article.query.filter(Article.title == 'aaa').first()  # a.先查找条件 title='aaa'的
    # print(article.author)  # 正向引用
    # print(article.author.username)  # b.然后直接找到user2表中作者(username): zhiliao
    # return 'index'

    # # 这里先新增的操作也可以看下，正常不用做这里的操作，这里只是告诉你还有个单独的用法
    # #    I: 创建一篇文章
    # article = Article(title='aaa', content='bbb') # 这里可以不加author_id=1
    # #    II. 指定插入的数据中 id = 1
    # article.author = User.query.filter(User.id == 1).first()  # 这里指定id=1，相当于上面author_id=1
    # db.session.add(article)
    # db.session.commit()

    # E. 查作者username=='zhiliao'的 写过哪些文章
    user = User.query.filter(User.username == 'zhiliao').first()  # 先拿到zhiliao的信息
    result = user.articles  # 反向引用
    for i in result:
        print(i.title)
    return 'index'


if __name__ == "__main__":
    app.run(debug=True)
