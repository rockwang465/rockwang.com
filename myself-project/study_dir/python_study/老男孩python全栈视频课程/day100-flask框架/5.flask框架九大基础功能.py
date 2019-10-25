from flask import Flask, render_template, request, make_response

# 1.静态文件和模板路径配置
app = Flask(__name__)  # __name__可以随便写，只是一个名字


# 5.自定义模板函数，给下面hello函数中使用，渲染到index.html中
def rock():
    return "<h1>橙子</h1>"


# 6.设置请求方式 methods=['GET', 'POST']
@app.route("/", methods=['GET', 'POST'])  # "/" 为url地址的根，即http://127.0.0.1:5000；也可以设为其他，如: "/index/"
def hello():
    # 3.返回字符串
    # return "Hello World!"

    # 8.(响应额外的数据)request方式获取数据(url中传参 http://127.0.0.1:5000/?k=123  传参k=123)
    print(request.args)  # args为get请求,pycharm中返回 ImmutableMultiDict([('k1', '123')])，此为获取的参数

    # 4.返回模板
    return render_template('index.html', k1='rock', k2=[1, 2, 3], k3={'name': 'alex', 'age': 73}, k4=rock)
    # 'index.html'表示模板html文件名， k1,k2,k3,k4为不同类型的值。 k4为上面定义的函数rock


# 7. 固定的url字段访问，如 http://127.0.0.1:5000/test/about，必须是(about, help, imprint, class, rock,"foo,bar") 其中之一的url字段, 可以自定义
@app.route('/test/<any(about, help, imprint, class, rock, "foo,bar"):page_name>')  # 可以自定义，但访问必须用其中的字段
def test(page_name):
    return page_name


if __name__ == "__main__":
    # 2.IP和端口的配置
    app.run()
