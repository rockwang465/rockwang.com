# encoding: utf-8

from exts import db  # 这样就可以用exts的db了


class Heros_detail_info(db.Model):
    __tablename__ = 'heros_detail_info'
    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    name = db.Column(db.String(20), nullable=False)
    country = db.Column(db.String(20), nullable=False)
    comment = db.Column(db.String(100), nullable=False)
    img = db.Column(db.Text, nullable=False)
    # img = db.Column(db.MediumText)
    # img = db.Column(db.Longblob)
    # 如果text字段长度不够，可以用mediumtext。 text是64kb，mediumtext是16mb。
