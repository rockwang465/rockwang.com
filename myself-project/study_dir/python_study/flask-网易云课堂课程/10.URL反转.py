from flask import Flask
from flask import url_for  # 要导入url_for

app = Flask(__name__)


@app.route("/")
def index():
    print(url_for('my_list'))  # 这里放入下面定义的视图函数名称，并用引号引起来
    print(url_for('article', id="rock666"))  # 这里放入定义的视图函数名称，且id要指定，否则报错
    return u'主页'


@app.route("/list/")
def my_list():
    return u'我的列表'


@app.route("/article/<id>")
def article(id):
    return u'文章，您输入的id为: %s' % id


if __name__ == "__main__":
    app.run()
