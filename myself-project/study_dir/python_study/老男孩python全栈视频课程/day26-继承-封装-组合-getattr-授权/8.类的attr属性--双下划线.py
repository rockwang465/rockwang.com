class Foo:
    x = 1

    def __init__(self, y):
        self.y = y

    def __getattr__(self, item):
        print('执行__getattr__')

    def __delattr__(self, item):
        print('执行__delattr__', item)

    def __setattr__(self, key, value):
        print('执行__setattr__', key, value)  # key, value 为设置时的key和赋值的value。
        if type(value) is str:
            # self.key = value  # 但这里定义，__setattr__会死循环的，因为会再次调用 __getattr__
            self.__dict__[key] = value  # 所以应该操作底层字典的方式，则不会死循环
        else:
            print("必须时字符串类型才可以")

f1 = Foo(10)  # 实例化的时候，相当于设置值，则会自动执行__setattr__
# print(getattr(f1, 'y'))
# f1.abcdefg  #当调用不存在的属性时，则执行 __getattr__

# del f1.y  #删除属性的时候，会执行 __delattr__

f1.x = 3  # 设置的时候会调用__setattr__

# 注意: 只有 __getattr__比较常用，其他都很少使用的
