# -*- encoding: utf-8 -*-

from flask import Flask
from flask import render_template  # 模版渲染方法

app = Flask(__name__)


@app.route('/<is_login>/')
def index(is_login):
    if is_login == '1':
        user = {
            'name': 'Rock',
            'age': 15
        }
        return render_template('index.html', user=user)
    else:
        return render_template('index.html')


if __name__ == "__main__":
    app.run(debug=True)
