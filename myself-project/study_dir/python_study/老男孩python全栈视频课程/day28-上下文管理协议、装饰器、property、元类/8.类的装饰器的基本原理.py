# 例1: 函数装饰器
# def test(fun):
#     print("test")
#     return fun
#
# @test  # 相当于: FOO=test(Foo)
# def Foo():
#     print("Foo")
# 结果 : test

# 例2: 函数装饰器用在类上
# def test(fun):
#     print("test")
#     return fun

# @test  # 相当于: FOO=test(Foo)
# class Foo:
#     print("Foo")

# 结果: Foo
#       test


# 例3: 类装饰器基本用法
def deco(obj):
    obj.x = 1
    obj.y = 2
    obj.z = 3
    return obj  # 需要返回obj的值才有__dict__


@deco  # Foo=deco(Foo)
class Foo:
    pass

print(Foo.__dict__)