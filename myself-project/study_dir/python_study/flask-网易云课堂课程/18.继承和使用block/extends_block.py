# -*- encoding: utf-8 -*-

from flask import Flask
from flask import render_template  # 模版渲染方法

app = Flask(__name__)


@app.route('/')
def index():
    return render_template('index.html')


@app.route('/login/')
def login():
    return render_template('login.html')


if __name__ == "__main__":
    app.run(debug=True)
