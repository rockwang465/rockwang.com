import re

s = '&ab234'

# print(re.search('..', s))  #匹配任意两个字符
# print(re.search('^&', s))  #匹配以xx开头的字符
# print(re.search('4$', s))  #匹配以xx结尾的字符


# print(re.match('b.b$', 'bob'))  #匹配以xx结尾的字符
# print(re.search('b$', 'bob'))  #匹配以xx结尾的字符

print(re.search('a*','Alex'))