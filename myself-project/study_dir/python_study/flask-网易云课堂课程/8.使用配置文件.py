from flask import Flask
import config  # 导入config.py文件

app = Flask(__name__)
app.config.from_object(config)  # 使用config


@app.route("/")
def index():
    return "主页"


if __name__ == "__main__":
    app.run()
