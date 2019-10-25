# -*- encoding: utf-8 -*-

from flask import Flask
from flask import render_template  # 模版渲染方法

app = Flask(__name__)


@app.route('/')
def index():
    comment = [
        {
            'name': 'Rock',
            'connent': '这本书不错，推荐+1'
        },
        {
            'name': 'Lss',
            'connent': '还有其他的推荐吗？'
        }
    ]

    # img_url 地址是右击任意图片，点击复制图片地址。
    # return render_template('index.html', img_url='https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=2031658692,3163565505&fm=26&gp=0.jpg')
    return render_template('index.html', comment=comment)


if __name__ == "__main__":
    app.run(debug=True)
