from flask import Flask

app = Flask(__name__)


@app.route("/article/<id>")
def article(id):
    return u'您传入的id为: %s' % id


if __name__ == "__main__":
    app.run()
