# encoding: utf-8

from flask import Flask, session  # 导入session
import os

app = Flask(__name__)
# 2. 添加一个SECRET_KEY,作用是给session加密的。也可以将配置放入config.py中。
# 2.1 值必须为24位的字符串, 所以用os.urandom(24)随机生成一个24位的字符串。用于打乱原有session的字符串和加密。
app.config['SECRET_KEY'] = os.urandom(24)


# 添加数据到session中
# 操作 session的时候，跟操作字典是一样的
# SECRET_KEY


@app.route('/')
def add():
    session['username'] = 'zhiliao'  # 1.添加一条session数据
    return 'hello world'


@app.route('/get/')
def get():  # 3. 获取session
    # session['username'] # 不建议用，没有此值则报错
    return session.get('username')  # 建议使用，没有此值返回None


@app.route('/delete/')
def delete():  # 4. 删除session
    print(session.get('username'))
    session.pop('username')  # 删除指定的一个key
    print(session.get('username'))
    return 'success'


@app.route('/clear/')
def clear():
    print(session.get('username'))
    session.clear()  # 删除session中的所有数据，pop只是删除单个key
    print(session.get('username'))
    return 'success'


if __name__ == '__main__':
    app.run(debug=True)

# 注意: 每次服务(pycharm)重启后，SECRET_KEY都会变化
# 当重启后访问服务时，拿浏览器之前保存的cookie去新的session比对时，则比对无法通过，就会页面报错。
# 为了解决重启后cookie的变化，可以将SECRET_KEY固定住。
