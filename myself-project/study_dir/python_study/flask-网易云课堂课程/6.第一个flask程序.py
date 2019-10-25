# encoding: utf-8

from flask import Flask  # 从flask框架中导入Flask这个类

# 初始化一个Flask对象，并命名叫app
# 需要传递一个参数 __name__，作用是:
#   1. 方便flask框架去寻找资源
#   2. 方便flask插件(比如Flask-Sqlalchemy出现错误的时候)好去寻找问题所在的位置
app = Flask(__name__)


# @app.route是一个装饰器
# 这个装饰器的作用，是做一个url域视图函数的映射
# 127.0.0.1:5000/  -->  去请求 hello_world 这个函数，然后将结果返回给浏览器
@app.route('/')
def hello_world():  # 视图函数
    return '我是第一个flask程序!'


# 如果当前文件作为入口程序运行，那么久执行app.run()
if __name__ == '__main__':
    app.run()  # 以while True:形式一直监听请求
