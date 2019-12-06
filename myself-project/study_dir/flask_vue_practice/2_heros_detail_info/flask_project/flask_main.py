# encoding: utf-8

from flask import Flask
from exts import db
import config
from models import Heros_detail_info  # 下面需要用到数据的操作，所以这里导入数据库的模型-rock
import json
from flask_cors import CORS

# import base64

app = Flask(__name__)
app.config.from_object(config)
db.init_app(app)  # db初始化app(在一个文件中db直接初始化；但这里导入外面的db，所以这里传递app给db，给exts.py中的db初始化app用的)


# 根目录信息
@app.route("/")
def index():
    return {"info": ["This is root directory"]}


# 根目录信息
@app.route("/v1/apps/api")
def api():
    api_info = {"api": [
        {"/": "root directory",
         "/v1/apps/api": "api",
         "/v1/apps/hero_detail_info/field": "hero_detail_info table filed information",
         "/v1/apps/batch_insert_data": "batch insert data",
         "/v1/apps/get_all_heros_data": "get all hero_detail_info table data",
         "/v1/apps/get_hero_info/<id>": "use id to get a hero info",
         }
    ]}
    return api_info


# 获取表字段,还未定义
@app.route("/v1/apps/hero_detail_info/field")
def field():
    return {"info": ["get table filed"]}


# 获取hero_detail_info表中所有数据
@app.route("/v1/apps/get_all_heros_data")
def get_all_heros_data():
    # result = Article.query.filter(Article.title == 'r2').all()
    # res = Heros_detail_info.query.count()
    res = Heros_detail_info.query.all()
    print(res)
    # 将获取到的表数据，通过json方式返回
    hero_detail_info_json = {"hero_detail_info": []}
    for i in res:  # res是个列表，所以用循环拿到每条数据
        tmp_dict = {"id": i.id, "name": i.name, "country": i.country, "comment": i.comment, "img": i.img}
        # print(tmp_dict)
        hero_detail_info_json["hero_detail_info"].append(tmp_dict)
    # print(hero_detail_info_json)
    return hero_detail_info_json
    # return u'%s' % hero_detail_info_json


# # 通过传参id，进行查询对应id的值
@app.route("/v1/apps/get_hero_info/<id>")
def get_hero_info(id):
    res = Heros_detail_info.query.filter(Heros_detail_info.id == id).all()[0]  # all()结果是列表，所以用[0]拿到第一条数据
    res_dict = {"id": res.id, "name": res.name, "country": res.country, "comment": res.comment, "img": res.img}
    # print(res_dict)
    return res_dict
    # return u'%s' % res_dict


# 批量数据插入 - 用于前期准备的数据，后期此部分可以注释掉
@app.route("/v1/apps/batch_insert_data")
def batch_insert_data():
    filename = "./init_mysql_data.json"
    with open(filename, encoding="utf-8") as f:
        # data1 = f.read().decode(encoding="utf-8").encode(encoding="utf-8")  # python2需要转码
        data = json.loads(f.read())
        # print(data.get("hero_detail_info")[0].get("img"))
    # 开始插入上面获取的data字典数据
    for sql in data.get("hero_detail_info"):
        # print(sql)
        print(sql.get("id"), sql.get("name"), sql.get("country"), sql.get("comment"), sql.get("img"))

        # 原本准备base64后放入数据库中，后因为sqlalchemy的text字段长度不够，只能用nginx做图片存储了。
        # img_path = sql.get("img")
        # with open(img_path, 'rb') as f2:
        #     base64_data = base64.b64encode(f2.read())
        #     s = base64_data.decode()
        #     img_data = 'data:image/jpeg;base64,%s' % s
        hero_detail_info1 = Heros_detail_info(id=sql.get("id"), name=sql.get("name"), country=sql.get("country"),
                                              comment=sql.get("comment"), img=sql.get("img"))
        db.session.add(hero_detail_info1)  # 添加数据
        db.session.commit()  # 提交事务
    return "完成批量数据上传"


if __name__ == "__main__":
    CORS(app, supports_credentials=True)
    # app.run(debug=True)
    app.run(host='0.0.0.0', port='5001', debug=True)
