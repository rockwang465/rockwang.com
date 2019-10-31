# encoding: utf-8

from flask import Flask
from exts import db
import config

app = Flask(__name__)
app.config.from_object(config)
db.init_app(app)  # db初始化app(这个上一讲说过，在一个文件中db直接初始化；但这里导入外面的db，所以这里传递app给db，给exts.py中的db初始化app用的)
