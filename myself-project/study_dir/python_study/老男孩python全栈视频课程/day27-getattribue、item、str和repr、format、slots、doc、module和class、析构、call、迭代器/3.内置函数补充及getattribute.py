# class Foo:
#     pass
#
# class Bar(Foo):
#     pass
#
# f1 = Foo()
# print(isinstance(f1, Foo))  #判断f1为Foo的实例化
# print(issubclass(Bar, Foo))  #判断Bar为Foo的子类


class Foo():
    def __init__(self, x):
        self.x = x

    def __getattr__(self, item):  # 如果没有__getattr__，则执行没有的属性时，则会执行raise中的错误
        print("这是__getattr__")

    def __getattribute__(self, item):
        print("这是__getattribute__")
        raise AttributeError("抛出异常，报错啦")


f1 = Foo(10)
f1.yy  # 调用的yy是不存在的

# 总结:
# __getattribute__ 不管何时都会执行
# 当 __getattribute__ 中有 raise 时，则只执行一次，然后执行 __getattr__
# 当没有__getattr__时，则会执行raise中的报错提示
# 当没有 __getattribute__ 时， 且调用方法不存在时，则会执行 __getattr__
