#!/usr/bin/env python
# -*- coding: utf-8 -*-

import redis

host_ip = "10.151.3.94"
password = "redis123"
redis_port = "31531"  # 提示只读，其实每3次，有一次连接到的是master
sentinel_port = "30379"  # 提示AUTH报错

red = redis.Redis(host=host_ip, port=redis_port, db=1, password=password)
res = red.set("name", "zhangsan")
print(res)  # True

res2 = red.get("name")
print(res2)  # zhangsan
