# -*- encoding: utf-8 -*-

from flask import Flask
from flask import redirect  # 重定向方法
from flask import url_for

app = Flask(__name__)


@app.route('/')
def index():
    # 配置: 当访问/ 根url时，让他跳转到login页面
    login_url = url_for('login')  # 1.获取login的url
    return redirect(login_url)  # 2.重定向到login的url，注意: 要加return
    return u"这是主页"


@app.route('/login/')
def login():
    return u"登录页面"


@app.route('/question/<is_login>')
def question(is_login):
    if is_login == '1':  # 当传参为1，表示为已经登录，允许继续访问
        return u'欢迎进入问答页面'
    else:  # 不为1，表示没有登录，则返回登录页面
        return redirect(url_for('login'))


if __name__ == "__main__":
    app.run(debug=True)
