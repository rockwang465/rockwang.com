# class Foo:
#     def __init__(self, x):
#         self.x = x
#
#     def __getattr__(self, item):
#         print("你找的属性 [%s] 不存在" % item)
#
# f1 = Foo(10)
# f1.y


class to_list(list):
    def append(self, p_object):
        if type(p_object) is str:
            list.append(self, p_object)  # 在self中添加p_object参数
            super().append(p_object)  # 用super() 更好
            print(self)  # 这里可以看到新加入的p_object值（即'rock'）
        else:
            print("只支持传入字符串")

    def show_middle(self):
        mid_index = int(len(self) / 2) + 1
        print(self[mid_index])  # 返回中间的字符


str1 = "hello_world_!"
liebiao = to_list(str1)  # 传入字符串到to_list中，self就是str1的列表形式
liebiao.show_middle()  # 现在中间的字符
liebiao.append("rock")  # 这里是调用了append函数，如果没有append函数，则调用系统的append方法。
liebiao.append(111)  # 传值到append函数中，p_object就是传入的值。

# print(to_list(str1), type(to_list(str1)))
# print(list(str1), type(list(str1)))
