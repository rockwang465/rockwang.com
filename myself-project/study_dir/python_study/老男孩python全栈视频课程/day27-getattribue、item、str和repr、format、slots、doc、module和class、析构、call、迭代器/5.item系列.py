class Foo:
    def __getitem__(self, item):
        print("__getitem__")
        return self.__dict__[item]

    def __setitem__(self, key, value):
        print("__setitem__")
        self.__dict__[key] = value  # 要这样赋值才对，self.key = value 会死循环的。

    def __delitem__(self, key):
        print("__delitem__")


f1 = Foo()
f1['name'] = 'rock'  # 设置值
print(f1.__dict__)
print("查看-->", f1['name'])  # 查看值
del f1['name']  # 删除值

# 总结:
# 总结:
# .  系列的操作，和 __xxxattr__ 相关
# [] 系列的操作，和 __xxxitem__ 相关