class Foo:
    '我是Foo的描述信息'
    pass

class Bar(Foo):
    pass

b1 = Bar()
print(b1.__doc__)
# 结果为None


# + 总结:
#   - 默认类里面自动添加`__doc__ = None`，所以当你找`__doc__`时，会显示None。
#   - 由于自身有None，所以不会继承父类的`__doc__`
#   - `__doc__`属性没啥用处。
