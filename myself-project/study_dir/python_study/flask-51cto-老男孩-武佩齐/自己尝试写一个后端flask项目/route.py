# encoding: utf-8

from flask import Flask, request
from main import *

app = Flask(__name__)


@app.route('/', methods=['get', 'post'])
def index():
    if request.method == 'GET':
        return 'hello world'
    elif request.method == 'POST':
        # http://127.0.0.1:5000/?env=10.5.6.92&infra_ansible=dev&tools=v1.2.1&version=v1.4.1
        print(request.args.to_dict())
        args_dict = request.args.to_dict()
        if args_dict:
            return t1.test_main(args_dict)
            # return args_dict
        else:
            return 'Not input args'
    else:
        return 'Error: request method is error'


if __name__ == '__main__':
    app.run(debug=True)
