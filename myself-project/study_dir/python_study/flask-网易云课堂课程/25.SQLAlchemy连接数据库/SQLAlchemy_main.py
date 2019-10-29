# -*- encoding: utf-8 -*-

from flask import Flask
from flask_sqlalchemy import SQLAlchemy
import config  # 导入config文件

app = Flask(__name__)
app.config.from_object(config)  # 使用config文件的配置

# 初始化SQLAlchemy
db = SQLAlchemy(app)

db.create_all()  # 做测试，看有没有问题（如果这里执行后没有报错，说明一切都成功）


@app.route('/')
def index():
    return 'index'


if __name__ == "__main__":
    app.run(debug=True)
