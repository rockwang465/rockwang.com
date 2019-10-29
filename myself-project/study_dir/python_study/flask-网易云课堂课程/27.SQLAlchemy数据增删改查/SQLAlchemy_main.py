# -*- encoding: utf-8 -*-

from flask import Flask
from flask_sqlalchemy import SQLAlchemy
import config  # 导入config文件

app = Flask(__name__)
app.config.from_object(config)  # 使用config文件的配置

# 初始化SQLAlchemy
db = SQLAlchemy(app)


class Article(db.Model):  # 继承db的Model模型
    __tablename__ = 'article'  # 表名定义
    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    title = db.Column(db.String(100), nullable=False)
    content = db.Column(db.Text, nullable=False)


db.create_all()  # 开始运行数据库的操作


@app.route('/')
def index():
    # 1 增加:
    # 1.1 需要插入的数据内容
    article1 = Article(title='aaa', content='bbb')
    # 1.2 引用sqlalchemy的session中add插入数据功能
    db.session.add(article1)
    # 1.3 事务
    db.session.commit()  # 提交
    return 'hello Rock !'  # 当每访问一次页面的时候，就会插入一条数据

    # 2 查找:
    # 2.1 Article前面已经引用了db.Model，所以这里可以直接用Article
    #     Article.query.filter 表示查询数据，并过滤后面括号中的条件
    #     Article.title 表示 Article中定义的article表的title字段
    #     .all() 返回一个对象数组
    result = Article.query.filter(Article.title == 'r2').all()
    #     .first() 取第一条数据用first()比[0]更好，因为如果没有数据会返回None的。
    # result = Article.query.filter(Article.title == 'r2').first()
    res1 = result[1]  # 取结果的第一个数据(以列表展示)[如果有多个值要注意这里取值问题]
    print(res1.title)  # 展示结果中title字段的值
    # print('title: %s' % res1.title)
    print(res1.content)  # 展示结果中content字段的值
    # print('content: %s' % res1.content)
    return 'hello Rock !!'

    # 3 修改:
    # 3.1 先把你要改的数据查出来
    res2 = Article.query.filter(Article.title == 'aaa').first()
    # 3.2 把这条数据，你需要修改的地方进行修改
    res2.title = 'new value2'  # 修改键值为 new value
    # 3.3 做事务提交
    db.session.commit()  # 提交修改
    return 'hello Rock !!!'

    # 4 删除:
    # 4.1 先把你要删的数据查出来
    res3 = Article.query.filter(Article.title == 'aaa').first()
    # 4.2 把这条数据删掉
    db.session.delete(res3)
    # 4.3 做事务提交
    db.session.commit()  # 当没有数据可删时，则会报错
    return 'hello Rock !!！!'


if __name__ == "__main__":
    app.run(debug=True)
