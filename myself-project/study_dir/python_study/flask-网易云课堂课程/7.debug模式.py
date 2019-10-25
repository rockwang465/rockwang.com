from flask import Flask

app = Flask(__name__)


@app.route("/")
def index():
    a = 1
    b = 0
    c = a / b  # 由于1/0报错，为了显示报错，就在下面启动了debug模式
    return "主页"


if __name__ == "__main__":
    app.run(debug=True)
