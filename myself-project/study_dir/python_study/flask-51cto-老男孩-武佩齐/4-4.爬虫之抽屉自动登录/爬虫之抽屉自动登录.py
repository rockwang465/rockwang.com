# encoding: utf-8

import requests
from bs4 import BeautifulSoup

# 1. 查看首页
url1 = "https://dig.chouti.com"
user_agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36"
r1 = requests.get(
    url=url1,
    headers={
        "user-agent": user_agent
    }
)
# print(r1.text)
# print(r1.cookies)  # 列表形式
print(r1.cookies.get_dict())  # 一般用这个，改为字典形式
print("------------------------")
# print(r1.cookies.get_dict().get('deviceId'))  # 一般用这个，改为字典形式
r1_deviceId = r1.cookies.get_dict().get('deviceId')  # 一般用这个，改为字典形式
# print("------------------------")

# 2. 提交用户名密码进行登录
url2 = "https://dig.chouti.com/login"  # 点击登录的时候，F12，network，Headers中会显示网址
r2 = requests.post(
    url=url2,
    headers={
        "Accept": "application/json, text/javascript, */*; q=0.01",
        "Accept-Encoding": "gzip, deflate, br",
        "Accept-Language": "zh-CN,zh;q=0.9",
        "Cache-Control": "no-cache",
        "Connection": "keep-alive",
        "Content-Length": "580",
        "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
        # "Cookie": "deviceId=web.eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqaWQiOiI5ZTExOGE1NS0yMjVhLTQ3M2EtOTkzZS01NTliOTdlOTZlMzUiLCJleHBpcmUiOiIxNTc1NDU5MzMyMjcxIn0.OvQ6pgOy1tR_K6DxViUviTpWjmXPfEXyLm65m_8TBvw; Hm_lvt_03b2668f8e8699e91d479d62bc7630f1=1572867365,1572953432; _9755xjdesxxd_=32; YD00000980905869%3AWM_NI=BaWkoo%2BUr0DyO%2B3TsyQHn3LlKbNeAP4QRI512DOJd0xln0vAzcIWqdwslY35Tno%2FXe8JYfMacyWMEWV%2FQL0R8HXYyDE1jWsSqdUsFaG5HoIilrYw8ArOby4Jo8SIoBfjQTk%3D; YD00000980905869%3AWM_NIKE=9ca17ae2e6ffcda170e2e6ee86f95ca286ffb6c674bc9a8aa3c14b978a9a84bc7f83f1bcb0e9448caa9b87db2af0fea7c3b92a82bf998cee65e98bbed1c7738df5fdb5f2739398a2d4cf6889be8286c45cf88f9eaab83bf29c81abe73cbcb4a2d1f54bbabd87b8f470b5ada4bab23f98b2a1d0db6ffc969787d06ba1a799aec85d96f19fadb35afc9898bafc5087efa6b8f259b3bc9b94c56b85eb8890cc63b0ae9798bb4bb291ae9ab742ade89daaee4aa6b69fb5ea37e2a3; YD00000980905869%3AWM_TID=DrnF1H%2BwQHNERRUFQRcs%2FFv4ajGTrAbS; gdxidpyhxdE=K4nX%2BlTh%5Cb%5CJOzl7%2BiaNhpvKzcDvGaXb%2BWEjJcz2XNzeHim2VT9h3MsWhXx2t6Z2Qj%2BWO5HqoJzm56o%2FTmsSeSLRYHeBeu%2Fho8xbDWgaALQeBG7PoA5hp7gTwZ8ks72vm0Sca7e%2F%2FDVMtZtQ%2BRwfLOCuDOUoA6jfkC5jTl8657UoDfcy%3A1573012340300; Hm_lpvt_03b2668f8e8699e91d479d62bc7630f1=1573012087",
        "Cookie": "deviceId=%s" % r1_deviceId,
        "user-agent": user_agent,
        "Host": "dig.chouti.com",
        "Origin": "https://dig.chouti.com",
        "Pragma": "no-cache",
        "Referer": "https://dig.chouti.com/",
        "Sec-Fetch-Mode": "cors",
        "Sec-Fetch-Site": "same-origin",
        "X-Requested-With": "XMLHttpRequest",
    },
    data={  # 这个就是浏览器中的Form Data部分
        # "phone": "+8613121758648",  # 武佩齐老师的账号
        "phone": "+8617521004208",
        # "password": "woshiniba",
        "password": "5229465cheng",
        # 'oneMonth': 1,
        "loginType": 2,
        # "NECaptchaValidate": "9FpUonOSLoreo8.v6fzayLeLCt8w.utq_vjxzS-cqT2MIoca6W95FKVRShGSogMsc0FtQ6pkza2GbsjO9BuM7WvHC1Ywy4KZvWbCF5Qq7U_GYuth1eYx_5T.g0j0KkPVxQhSJyEqbcQDjcDhqC7226VeiftwieOohxXS_opWqwHvA.Jd9HAUBVoBMM8Izq-ZV9a8UO1o5rFT_Y-VguuIJCytHsKsoxBmOxZJgfStPP.WX2XtMsJdUo5z5lJkkrNBejfzJsYVMCcrgt1PYKSSvB0xWz2G4-u50TyyxZLFfDsorZXXtjTEVFC8rAa_1J4nev71C05b62SBHiDqjfIG1fdBWa4FFgQx9qOI0Q2IPvvHlimG-5iZtekUGqJxWmrVR6i8nXY0eauUIiAeUR4nFRcO8oFjEEePgB-ujtabyyWG7aTKsmanwDfGu1WQ2bEnBq.g1ATjc_sT9N1NEIf1Z_CjiSczidLnzpsS_DMFFyu_iftQPoNg_es-QPc3"
    },
    # 2.1 有的需要带入cookie，给后面使用
    # cookies=r1.cookies.get_dict()  # 不同的网站，套路不同
)
# 注意，我这里执行后是报错的，还没找到原因。老师执行是可以的。
print(r2.text)
# print(r2.cookies.get_dict())  # 获取cookie信息

# # 3. 点赞一篇文章
url3 = "https://dig.chouti.com/link/vote?linkId=20435396"
r3 = requests.post(
    url=url3,
    headers={
        'user-agent': user_agent
    },
    # cookies=r2.cookies.get_dict() # 发现用r2的cookie不行，
    # 则在上面2.1中加入cookies，用于这里继续引用之前的cookie
    # 不同家的网站，套路不同
    cookies=r1.cookies.get_dict()  # 用之前的cookie,而非r2的cookie
)

