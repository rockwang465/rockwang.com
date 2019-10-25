# -*- encoding: utf-8 -*-

from flask import Flask
from flask import render_template  # 模版渲染方法

app = Flask(__name__)


@app.route('/')
def index():
    # 1.渲染传参地方
    # return render_template('index.html', username=u'Rock', gender=u'男', age=u'18')

    # 2. 优化传参方式
    context = {  # 2.1 通过字典方式传参
        'username': u'Rock',
        'gender': u'男',
        'age': u'18'
    }
    # 2.2 使用 **context 传入所有参数
    return render_template('index.html', **context)


if __name__ == "__main__":
    app.run(debug=True)
