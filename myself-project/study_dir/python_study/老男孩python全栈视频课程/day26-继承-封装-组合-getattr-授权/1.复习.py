dic = {'k1': 123, 123: 123, '999': 123}
# print(dic.items())
# 结果 : dict_items([('k1', 123), (123, 123), ('999', 123)])

# for i in dic:
#     print(i)
# 结果 : k1  123  999

# for index, value in enumerate(dic.items(), 1):
#     print(index, value)
# 索引从1开始
# 结果: 1 ('k1', 123)    2 (123, 123)     3 ('999', 123)

# for i, v in enumerate(dic.items()):
#     new_list = list(v)
#     new_list.insert(0, i)
#     print(" ".join("%s" % v2 for v2 in new_list))  # 先要把值转为str，才能用''.join()，否则报错类型错误
# 结果: 0 k1 123    1 123 123    2 999 123

for i, v in enumerate(dic, 1):
    # print(i, v, dic.get(v))
    print(i, v, dic[v])
# (这个比较好)结果: 0 k1 123    1 123 123    2 999 123
