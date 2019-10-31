# encoding: utf-8

from flask import Flask
from models import Article
from exts import db
import config

app = Flask(__name__)
app.config.from_object(config)

db.init_app(app)  # 给db传入一个app进行初始化

# app推入到栈顶中(即把app推入app上下文中)
with app.app_context():  # 如果是导入文件，则需要用此推入上下文方式，否则报错
    db.create_all()


@app.route('/')
def index():
    return 'hello my world'


if __name__ == '__main__':
    app.run(debug=True)
