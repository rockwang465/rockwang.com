# encoding: utf-8

from functools import wraps
from flask import session, redirect, url_for


# 登录限制的装饰器
def login_required(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        if session.get('user_id'):
            return func(*args, **kwargs)
        else:
            # Rock:注意，这里必须加return，否则上层函数无法拿到返回值，则无法被调用
            return redirect(url_for('login'))

    return wrapper
