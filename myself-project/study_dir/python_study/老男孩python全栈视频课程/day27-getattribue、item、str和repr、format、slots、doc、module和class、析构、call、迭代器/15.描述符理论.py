class Foo:
    def __get__(self, instance, owner):
        print("__get__")

    def __set__(self, instance, value):
        print("__set__")

    def __delete__(self, instance):
        print("__delete__")


class Bar:
    x = Foo()  # A.必须要先被其他类引用


# B. 要实例化
b1 = Bar()
# C. 最终才能执行里面的 get set delete操作，才有效果
b1.x  # get
b1.x = 2  # set
del b1.x  # delete
print(b1.__dict__)  # 上面没有做赋值操作，所以里面肯定为空