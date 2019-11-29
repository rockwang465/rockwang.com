# encoding: utf-8

import requests
from bs4 import BeautifulSoup
import re

url1 = "https://passport.lagou.com/login/login.html"
# 1. 通过定义user_agent浏览器信息，来伪装请求为浏览器的请求。通过浏览器访问网站，F12，network，response中的请求头内容中找到user_agent
user_agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36"
r1 = requests.get(
    url=url1,
    headers={
        "user-agent": user_agent
    }
)
# print(r1.text)
print(r1.cookies.get_dict())


# 获取官方自定制的两个动态token
# window.X_Anti_Forge_Token = '';
# window.X_Anti_Forge_Code = ''
X_Anti_Forge_Token = re.findall("X_Anti_Forge_Token = '(.*?)'", r1.text, re.S)[0]
X_Anti_Forge_Code = re.findall("X_Anti_Forge_Code = '(.*?)'", r1.text, re.S)[0]
print(X_Anti_Forge_Token)
print(X_Anti_Forge_Code)
print("--------------------------------------")

# 进行登录页面请求
url2 = "https://passport.lagou.com/login/login.json"
r2 = requests.post(
    url=url2,
    headers={
        "User-Agent": user_agent,
        "X_Anti_Forge_Code": X_Anti_Forge_Code,
        "X_Anti_Forge_Token": X_Anti_Forge_Token,
        # Referer用于防盗链，要放入上一次请求地址是什么？
        "Referer": "https://passport.lagou.com/login/login.html",  # 上一次请求地址是什么？
        # "Referer": "https://www.lagou.com/",  # 上一次请求地址是什么？
        # "Sec-Fetch-Mode": "no-cors",
        # "Sec-Fetch-Site": "same-site"
    },
    data={
        "isValidate": True,  # true
        "username": "15131255089",
        # "username": "17521004208",
        "password": "ab18d270d7126ea65915c50288c22c0d",
        # "password": "3557fe927b30a6cf4721b89dd6c69fb2",
        "request_form_verifyCode": "",
        "submit": "",
        # "soncallback": "jQuery1113009683463918345114_1573025151047",
        # "challenge": "878bc8c923c49a099cadbbcc60a9661e",
        # "_": "1573025151051"
    },
    cookies=r1.cookies.get_dict()
)


print(r2.text)
# bs = BeautifulSoup(r1.text, features='lxml'
