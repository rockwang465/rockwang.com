import re

# s = 'abc21d3e'
s = 'adsdf'

# print(re.search('[0-9]',s))
# print(re.search('[0-9]',s).group())

result = re.search('[0-9]',s)
if result:
    print(result.group())
else:
    print(result)
