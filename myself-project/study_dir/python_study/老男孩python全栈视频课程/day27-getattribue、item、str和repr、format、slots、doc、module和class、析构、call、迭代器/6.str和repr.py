# class Foo:
#     def __init__(self, name, age):
#         self.name = name
#         self.age = age
#
#     def __str__(self):
#         return "名字: %s , 年龄: %s " % (self.name, self.age)
#
#
# f1 = Foo('rock', 18)
# print(f1)

# 结果:
# 当没有 __str__定义时 : <__main__.Foo object at 0x00B9B8B0>
# 当有 __str__定义时 : 名字: rock , 年龄: 18

# + 总结:
#   - 默认函数中是有个 `__str__`返回函数名的。
#   - 当类中定义了 `__str__`后，则返回定义的内容。


class Foo:
    def __init__(self, name, age):
        self.name = name
        self.age = age

    def __str__(self):
        return "这是str"

    def __repr__(self):
        return "名字: %s , 年龄: %s " % (self.name, self.age)


f1 = Foo('rock', 18)
print(f1)
