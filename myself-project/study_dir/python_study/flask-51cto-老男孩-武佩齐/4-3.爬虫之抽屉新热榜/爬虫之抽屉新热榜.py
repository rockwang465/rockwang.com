# encoding: utf-8

import requests
from bs4 import BeautifulSoup

url1 = "https://dig.chouti.com"
# 1. 通过定义user_agent浏览器信息，来伪装请求为浏览器的请求。通过浏览器访问网站，F12，network，response中的请求头内容中找到user_agent
user_agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36"
# res = requests.get(url)
res = requests.get(
    url=url1,
    headers={
        "user-agent": user_agent
    }
)
# res.encoding = "gbk"  # 这里不用修改字符编码
# print(res.text)

soup = BeautifulSoup(res.text, 'html.parser')

# 2. link_div是一个标签对象，因为用的find，只获取第一个div对应的class名，可以获取对象下的子标签内容
# link_div = soup.find(name='div', _class='link-con')
link_div = soup.find(name='div', attrs={'class': 'link-con'})  # 字典形式
# print(link_div.text)

# 3. link_list是一个列表，因为是find_all，会获取所有匹配到的class名，只能for循环后，获取列表中的对象的子标签内容
link_list = link_div.find_all(name='div', attrs={'class': 'link-item'})
# print(link_list)

# 循环遍历
for item in link_list:
    # 4. 拿到里面的a标签中的标题文字部分， 可以放多个class名在一个value中
    a = item.find(name='a', attrs={'class': 'link-title link-statistics'})
    print(a.text.strip())   # 打印a标签中的文字部分，并且去除空行
