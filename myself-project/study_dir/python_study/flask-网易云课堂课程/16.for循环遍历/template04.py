# -*- encoding: utf-8 -*-

from flask import Flask
from flask import render_template  # 模版渲染方法

app = Flask(__name__)


@app.route('/')
def index():
    user = {
        'name': 'Rock',
        'age': 15
    }
    books = [
        {
            'name': u'西游记',
            'author': u'吴承恩',
            'price': 120
        },
        {
            'name': u'红楼梦',
            'author': u'曹雪芹',
            'price': 200
        },
        {
            'name': u'水浒传',
            'author': u'施耐庵',
            'price': 140
        },
        {
            'name': u'三国演义',
            'author': u'罗贯中',
            'price': 170
        }
    ]
    return render_template('index.html', user=user, books=books)


if __name__ == "__main__":
    app.run(debug=True)
