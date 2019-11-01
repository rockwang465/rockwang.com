# encoding: utf-8

from flask import Flask, render_template, request
import config

app = Flask(__name__)
app.config.from_object(config)


@app.route('/')
def index():
    # 1.5 渲染index.html
    return render_template('index.html')


@app.route('/login/', methods=['get', 'post'])
def login():
    if request.method == 'GET':  # 当用户是get请求时(查看页面内容)
        return render_template('login.html')  # 则展示login页面内容
    else:
        pass


@app.route('/regist/', methods=['get', 'post'])
def regist():
    if request.method == 'GET':  # 当用户是get请求时(查看页面内容)
        return render_template('regist.html')  # 则展示login页面内容
    else:
        pass


if __name__ == '__main__':
    app.run()
